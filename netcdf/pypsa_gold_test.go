package netcdf

import (
	"math"
	"testing"
)

// TestPyPSAGoldStandard reads a real-world PyPSA-USA network file
// (NetCDF4/HDF5) and checks it against known metadata. See the
// accompanying testdata/pypsa-usa-caiso-2025-demo.nc.info.md.
func TestPyPSAGoldStandard(t *testing.T) {
	nc, err := Open("testdata/pypsa-usa-caiso-2025-demo.nc")
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()

	attrs := nc.Attributes()
	if got := len(attrs.Keys()); got != 9 {
		t.Errorf("global attributes: got %d, want 9", got)
	}
	if v, ok := attrs.Get("network_pypsa_version"); !ok || v != "0.30.2" {
		t.Errorf("network_pypsa_version: got %v, want 0.30.2", v)
	}
	if v, ok := attrs.Get("network_objective"); !ok {
		t.Error("network_objective attribute missing")
	} else if obj, isFloat := v.(float64); !isFloat || math.Abs(obj-72727606.62626582) > 1e-6 {
		t.Errorf("network_objective: got %v, want 72727606.62626582", v)
	}

	vars := nc.ListVariables()
	if len(vars) != 150 {
		t.Errorf("variables: got %d, want 150", len(vars))
	}
	if got := len(nc.ListSubgroups()); got != 0 {
		t.Errorf("subgroups: got %d, want 0", got)
	}

	// Every variable must be fully readable.
	for _, name := range vars {
		vg, err := nc.GetVarGetter(name)
		if err != nil {
			t.Errorf("GetVarGetter(%q): %v", name, err)
			continue
		}
		if _, err := vg.GetSlice(0, vg.Len()); err != nil {
			t.Errorf("read %q: %v", name, err)
		}
	}

	// Spot-check dimension sizes against the known metadata.
	wantLens := map[string]int64{
		"snapshots":       56,
		"generators_i":    2867,
		"buses_i":         70,
		"links_i":         280,
		"storage_units_i": 32,
		"loads_i":         58,
		"carriers_i":      19,
	}
	for name, want := range wantLens {
		vg, err := nc.GetVarGetter(name)
		if err != nil {
			t.Errorf("GetVarGetter(%q): %v", name, err)
			continue
		}
		if vg.Len() != want {
			t.Errorf("%s length: got %d, want %d", name, vg.Len(), want)
		}
	}

	// Spot-check decoded values: string and float64 variables.
	buses, err := nc.GetVariable("buses_i")
	if err != nil {
		t.Fatal(err)
	}
	names, ok := buses.Values.([]string)
	if !ok {
		t.Fatalf("buses_i: got %T, want []string", buses.Values)
	}
	if names[0] != "Alameda County CA (06001)" {
		t.Errorf("buses_i[0]: got %q, want %q", names[0], "Alameda County CA (06001)")
	}

	busesX, err := nc.GetVariable("buses_x")
	if err != nil {
		t.Fatal(err)
	}
	xs, ok := busesX.Values.([]float64)
	if !ok {
		t.Fatalf("buses_x: got %T, want []float64", busesX.Values)
	}
	if math.Abs(xs[0]-(-122.05424998196247)) > 1e-9 {
		t.Errorf("buses_x[0]: got %v, want -122.05424998196247", xs[0])
	}

	// 2D time-series variable: 56 snapshots x 2758 generators.
	genP, err := nc.GetVariable("generators_t_p")
	if err != nil {
		t.Fatal(err)
	}
	rows, ok := genP.Values.([][]float64)
	if !ok {
		t.Fatalf("generators_t_p: got %T, want [][]float64", genP.Values)
	}
	if len(rows) != 56 || len(rows[0]) != 2758 {
		t.Errorf("generators_t_p shape: got %dx%d, want 56x2758", len(rows), len(rows[0]))
	}
}
