package hdf5

import (
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

// Compile-time assertion that we have an api.Group from Open(). Keeps
// the import used even if tests are stripped.
var _ api.Group = (*HDF5)(nil)
