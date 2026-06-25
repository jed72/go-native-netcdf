// Command mkmanaged creates a minimal HDF5 file whose "/dense" group holds a
// single small string attribute stored densely on the fractal heap's
// *managed* track (heap-ID type 0).
//
// Forcing dense storage requires setting the attribute phase-change
// thresholds to (0, 0) on the group creation property list so the very first
// attribute goes into the fractal heap rather than the object header. The
// attribute is small enough to stay on the managed track and short enough that
// the group's v2 B-tree (the attribute-name index) holds exactly one leaf node
// ("BTLF") with exactly one type-8 record.
//
// That single, predictable record is what the SCN-003 negative-path test
// byte-patches: flipping the record's heap-ID type nibble to 2 (tiny) or 3
// (filtered/reserved) — tracks the reader must reject with a named error.
// Neither track can be produced by a real attribute through libhdf5, so a
// hand-patched managed fixture is the only way to exercise those branches.
//
// Usage: mkmanaged <output.h5>
package main

// #cgo pkg-config: hdf5
// #cgo LDFLAGS: -lhdf5
// #include <stdlib.h>
// #include <hdf5.h>
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: mkmanaged <output.h5>\n")
		os.Exit(1)
	}
	filename := C.CString(os.Args[1])
	defer C.free(unsafe.Pointer(filename))

	// Use the latest libver so v2 B-trees and fractal heaps are available.
	fapl := C.H5Pcreate(C.H5P_FILE_ACCESS)
	if fapl < 0 {
		fatal("H5Pcreate(FILE_ACCESS)")
	}
	defer C.H5Pclose(fapl)
	if C.H5Pset_libver_bounds(fapl, C.H5F_LIBVER_LATEST, C.H5F_LIBVER_LATEST) < 0 {
		fatal("H5Pset_libver_bounds")
	}

	fid := C.H5Fcreate(filename, C.H5F_ACC_TRUNC, C.H5P_DEFAULT, fapl)
	if fid < 0 {
		fatal("H5Fcreate")
	}
	defer C.H5Fclose(fid)

	// A group-creation property list that forces dense storage of attributes
	// (max_compact=0, min_dense=0), exactly as mkhuge does — but here every
	// attribute stays on the managed track.
	gcpl := C.H5Pcreate(C.H5P_GROUP_CREATE)
	if gcpl < 0 {
		fatal("H5Pcreate(GROUP_CREATE)")
	}
	defer C.H5Pclose(gcpl)
	if C.H5Pset_attr_phase_change(gcpl, 0, 0) < 0 {
		fatal("H5Pset_attr_phase_change")
	}

	cname := C.CString("dense")
	defer C.free(unsafe.Pointer(cname))
	gid := C.H5Gcreate2(fid, cname, C.H5P_DEFAULT, gcpl, C.H5P_DEFAULT)
	if gid < 0 {
		fatal("H5Gcreate2(dense)")
	}
	defer C.H5Gclose(gid)

	// A single small managed-track attribute. One attribute => one type-8
	// record in one BTLF leaf node.
	writeStringAttr(C.hid_t(gid), "small", "hi")
}

func writeStringAttr(obj C.hid_t, name, value string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cval := C.CString(value)
	defer C.free(unsafe.Pointer(cval))

	sid := C.H5Screate(C.H5S_SCALAR)
	tid := C.H5Tcopy(C.H5T_C_S1)
	C.H5Tset_size(tid, C.size_t(len(value)+1))
	C.H5Tset_strpad(tid, C.H5T_STR_NULLTERM)
	aid := C.H5Acreate2(obj, cname, tid, sid, C.H5P_DEFAULT, C.H5P_DEFAULT)
	if aid < 0 {
		fatal("H5Acreate2(" + name + ")")
	}
	if C.H5Awrite(aid, tid, unsafe.Pointer(cval)) < 0 {
		fatal("H5Awrite(" + name + ")")
	}
	C.H5Aclose(aid)
	C.H5Tclose(tid)
	C.H5Sclose(sid)
}

func fatal(msg string) {
	fmt.Fprintf(os.Stderr, "mkmanaged: %s failed\n", msg)
	os.Exit(1)
}
