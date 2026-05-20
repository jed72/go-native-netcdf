// Command mkhuge creates a minimal HDF5 file whose root group has one
// dense-stored string attribute large enough to be promoted to a "huge"
// fractal heap ID.
//
// Forcing dense storage requires setting the attribute phase-change
// thresholds to (0, 0) on the group creation property list so the very
// first attribute goes into the fractal heap rather than the object
// header. The attribute payload (~5000 bytes) then exceeds the default
// max-managed-object-size (4096) and HDF5 stores it on the heap's huge
// track — which is precisely what the Go reader must handle.
//
// Usage: mkhuge <output.h5>
package main

// #cgo pkg-config: hdf5
// #cgo LDFLAGS: -lhdf5
// #include <stdlib.h>
// #include <hdf5.h>
import "C"

import (
	"fmt"
	"os"
	"strings"
	"unsafe"
)

const (
	attrName     = "payload"
	attrSize     = 5000
	attrFiller   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: mkhuge <output.h5>\n")
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

	// Reopen the root group with a group-creation property list that forces
	// dense storage of attributes (max_compact=0, min_dense=0).
	rootGID := C.H5Gopen2(fid, C.CString("/"), C.H5P_DEFAULT)
	if rootGID < 0 {
		fatal("H5Gopen2(root)")
	}
	defer C.H5Gclose(rootGID)

	// We can't change attribute-phase change on an already-open group's
	// creation property list, so instead create a *new* group with the
	// forced-dense property list and put the big attribute on it. That
	// triggers exactly the same huge-heap-ID code path as PyPSA-Eur uses
	// at root.
	gcpl := C.H5Pcreate(C.H5P_GROUP_CREATE)
	if gcpl < 0 {
		fatal("H5Pcreate(GROUP_CREATE)")
	}
	defer C.H5Pclose(gcpl)
	if C.H5Pset_attr_phase_change(gcpl, 0, 0) < 0 {
		fatal("H5Pset_attr_phase_change")
	}

	gid := C.H5Gcreate2(fid, C.CString("dense"), C.H5P_DEFAULT, gcpl, C.H5P_DEFAULT)
	if gid < 0 {
		fatal("H5Gcreate2(dense)")
	}
	defer C.H5Gclose(gid)

	// Build a deterministic, easy-to-verify payload.
	payload := strings.Repeat(attrFiller, attrSize/len(attrFiller)+1)[:attrSize]
	writeStringAttr(C.hid_t(gid), attrName, payload)

	// Also write a small managed-track attribute so the file exercises both
	// paths — managed records must keep working alongside huge records.
	writeStringAttr(C.hid_t(gid), "small", "hello")
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
	fmt.Fprintf(os.Stderr, "mkhuge: %s failed\n", msg)
	os.Exit(1)
}
