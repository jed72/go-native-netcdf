package hdf5

import (
	"bytes"
	"encoding/binary"
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/batchatco/go-native-netcdf/netcdf/api"
)

// genHugeFile builds testdata/mkhuge and uses it to generate an HDF5 file
// containing a group with one large (5000-byte) dense-stored attribute,
// which HDF5 promotes to a "huge" fractal-heap ID. Skips the test if the
// HDF5 C library is unavailable.
func genHugeFile(t *testing.T) string {
	t.Helper()
	bin := "testdata/mkhuge_bin"
	buildCmd := exec.Command("go", "build", "-o", bin, "./testdata/mkhuge")
	buildCmd.Env = append(os.Environ(), "CGO_ENABLED=1")
	if out, err := buildCmd.CombinedOutput(); err != nil {
		t.Skip("cannot build mkhuge (HDF5 C library may not be available):", string(out))
		return ""
	}
	t.Cleanup(func() { os.Remove(bin) })

	filename := "testdata/testhuge.h5"
	cmd := exec.Command("./"+bin, filename)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatal("mkhuge failed:", string(out))
	}
	t.Cleanup(func() { os.Remove(filename) })
	return filename
}

// TestHugeHeapID covers SCN-002: a >4 KB attribute lives on the fractal
// heap's huge track and must round-trip byte-equal through netcdf.Open().
func TestHugeHeapID(t *testing.T) {
	filename := genHugeFile(t)
	if filename == "" {
		return
	}

	nc, err := Open(filename)
	if err != nil {
		t.Fatalf("Open failed: %v", err)
	}
	defer nc.Close()

	g, err := nc.GetGroup("/dense")
	if err != nil {
		t.Fatalf("GetGroup(/dense) failed: %v", err)
	}
	defer g.Close()

	attrs := g.Attributes()
	if attrs == nil {
		t.Fatal("no attributes on /dense")
	}
	keys := attrs.Keys()
	if len(keys) != 2 {
		t.Fatalf("want 2 attributes, got %d: %v", len(keys), keys)
	}

	// The huge one.
	v, has := attrs.Get("payload")
	if !has {
		t.Fatal(`attribute "payload" missing`)
	}
	s, ok := v.(string)
	if !ok {
		t.Fatalf(`attribute "payload" not a string: %T`, v)
	}
	if len(s) != 5000 {
		t.Errorf("payload length: want 5000, got %d", len(s))
	}
	const filler = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	want := strings.Repeat(filler, 5000/len(filler)+1)[:5000]
	if s != want {
		t.Errorf("payload bytes differ — first mismatch around position %d", firstMismatch(s, want))
	}

	// The managed one — must still work alongside the huge one.
	v, has = attrs.Get("small")
	if !has {
		t.Fatal(`attribute "small" missing`)
	}
	if v != "hello" {
		t.Errorf(`attribute "small": want "hello", got %q`, v)
	}
}

// TestHugeHeapID_PyPSAEur covers SCN-001: the real PyPSA-Eur file with a
// ~45 KB meta attribute opens, surfaces 138 datasets, and round-trips
// meta. Gated on the file existing locally (PYPSA_NETCDF_TEST_FILE env
// var or the canonical path).
func TestHugeHeapID_PyPSAEur(t *testing.T) {
	path := os.Getenv("PYPSA_NETCDF_TEST_FILE")
	if path == "" {
		path = "/Users/jed/dev/pypsa-eur/results/my200/networks/base_s_200_elec_730h_solved.nc"
	}
	if _, err := os.Stat(path); err != nil {
		t.Skipf("PyPSA-Eur fixture not available at %s: %v", path, err)
	}

	nc, err := Open(path)
	if err != nil {
		t.Fatalf("Open failed: %v", err)
	}
	defer nc.Close()

	vars := nc.ListVariables()
	if len(vars) != 138 {
		t.Errorf("ListVariables: want 138, got %d", len(vars))
	}

	attrs := nc.Attributes()
	if attrs == nil {
		t.Fatal("no root attributes")
	}
	meta, has := attrs.Get("meta")
	if !has {
		t.Fatal(`attribute "meta" missing`)
	}
	s, ok := meta.(string)
	if !ok {
		t.Fatalf(`attribute "meta" not a string: %T`, meta)
	}
	if len(s) < 40000 {
		t.Errorf("meta length: want > 40000 bytes, got %d", len(s))
	}
}

// We compare nominally identical strings; if they differ this helper
// finds the first byte offset.
func firstMismatch(a, b string) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	if len(a) != len(b) {
		return n
	}
	return -1
}

// genManagedFile builds testdata/mkmanaged and uses it to generate an HDF5
// file whose "/dense" group holds one small managed-track attribute. Skips
// the test if the HDF5 C library is unavailable.
func genManagedFile(t *testing.T) string {
	t.Helper()
	bin := "testdata/mkmanaged_bin"
	buildCmd := exec.Command("go", "build", "-o", bin, "./testdata/mkmanaged")
	buildCmd.Env = append(os.Environ(), "CGO_ENABLED=1")
	if out, err := buildCmd.CombinedOutput(); err != nil {
		t.Skip("cannot build mkmanaged (HDF5 C library may not be available):", string(out))
		return ""
	}
	t.Cleanup(func() { os.Remove(bin) })

	filename := "testdata/testmanaged.h5"
	cmd := exec.Command("./"+bin, filename)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatal("mkmanaged failed:", string(out))
	}
	t.Cleanup(func() { os.Remove(filename) })
	return filename
}

// TestUnsupportedHeapID covers SCN-003: when an attribute index record carries
// a heap ID on the tiny (idType 2) or filtered/reserved (idType 3) track —
// tracks this reader does not support — Open() must surface a clean error that
// wraps ErrUnsupportedHeapID and names the track, never a panic.
//
// Neither track is reachable through a real attribute via libhdf5, so we take
// a genuine managed-track fixture and surgically flip the single type-8
// record's heap-ID type nibble, repairing the v2 B-tree leaf checksum so the
// reader reaches the heap-ID branch rather than tripping the checksum first.
func TestUnsupportedHeapID(t *testing.T) {
	src := genManagedFile(t)
	if src == "" {
		return
	}
	data, err := os.ReadFile(src)
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	// Locate the one attribute-name v2 B-tree leaf. The fixture is built to
	// contain exactly one, holding exactly one type-8 record.
	if n := bytes.Count(data, []byte("BTLF")); n != 1 {
		t.Fatalf("fixture invariant broken: want exactly 1 BTLF leaf, got %d", n)
	}
	leaf := bytes.Index(data, []byte("BTLF"))

	// BTLF node: "BTLF"(4) + version(1) + type(1) + records + checksum(4).
	// One type-8 record is 17 bytes (8-byte heap ID + 1 flags + 4 creation
	// order + 4 name hash), so the checksummed region is 6+17 = 23 bytes.
	const recHdr = 6
	const recSize = 17
	nbytes := recHdr + recSize
	recByte := leaf + recHdr // first byte of the heap ID = version-and-type byte

	// Self-check the layout against the on-disk checksum before trusting it:
	// if this fails, the fixture's shape changed and the offsets below are
	// no longer valid.
	wantSum := binary.LittleEndian.Uint32(data[leaf+nbytes : leaf+nbytes+4])
	gotSum := computeChecksumStream(bytes.NewReader(data[leaf:leaf+nbytes]), nbytes)
	if gotSum != wantSum {
		t.Fatalf("BTLF layout assumption wrong: recomputed checksum 0x%08x != on-disk 0x%08x", gotSum, wantSum)
	}
	if idType := (data[recByte] >> 4) & 0b11; idType != 0 {
		t.Fatalf("fixture invariant broken: want managed idType 0 record, got %d", idType)
	}

	cases := []struct {
		name   string
		idType byte
		track  string // substring the error must name
	}{
		{"tiny", 2, "tiny"},
		{"filtered_reserved", 3, "filtered"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			patched := make([]byte, len(data))
			copy(patched, data)
			// Flip just the heap-ID type nibble (bits 4-5).
			patched[recByte] = (data[recByte] &^ (0b11 << 4)) | (tc.idType << 4)
			// Repair the leaf checksum so the reader reaches the heap-ID
			// branch rather than failing the checksum assertion first.
			fixed := computeChecksumStream(bytes.NewReader(patched[leaf:leaf+nbytes]), nbytes)
			binary.LittleEndian.PutUint32(patched[leaf+nbytes:leaf+nbytes+4], fixed)

			dst := t.TempDir() + "/patched.h5"
			if err := os.WriteFile(dst, patched, 0o644); err != nil {
				t.Fatalf("write patched fixture: %v", err)
			}

			nc, err := Open(dst)
			if err == nil {
				nc.Close()
				t.Fatalf("Open succeeded; want ErrUnsupportedHeapID for idType %d", tc.idType)
			}
			if !errors.Is(err, ErrUnsupportedHeapID) {
				t.Fatalf("error does not wrap ErrUnsupportedHeapID: %v", err)
			}
			if !strings.Contains(err.Error(), tc.track) {
				t.Errorf("error message does not name the %q track: %v", tc.track, err)
			}
		})
	}
}

// Compile-time assertion that we have an api.Group from Open(). Keeps
// the import used even if tests are stripped.
var _ api.Group = (*HDF5)(nil)
