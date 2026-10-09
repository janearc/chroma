package main

import (
	"os"
	"testing"
)

// TestSpectrumIsBuiltIn reads a spectrum in a folder with no data/ in it,
// as an installed spectra runs: the copy built into the program answers.
func TestSpectrumIsBuiltIn(t *testing.T) {
	t.Chdir(t.TempDir())
	curve, err := readSpectrum(dataPath("neptune-disc.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if len(curve.nm) == 0 {
		t.Fatal("neptune-disc.csv read as empty")
	}
}

// TestRegionsAreBuiltIn is the same for a sheet of sampled regions.
func TestRegionsAreBuiltIn(t *testing.T) {
	t.Chdir(t.TempDir())
	found, err := regions(dataPath("pluto-2015.css"))
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"whale", "midlands", "heart", "pole"} {
		if _, has := found[name]; !has {
			t.Errorf("pluto-2015.css has no region %s", name)
		}
	}
}

// TestDiskWins reads a data file made on disk in place of the copy built
// in, as webb's output is read before spectra is built again.
func TestDiskWins(t *testing.T) {
	t.Chdir(t.TempDir())
	if err := os.Mkdir("data", 0o755); err != nil {
		t.Fatal(err)
	}
	made := "# made on disk\nnm,value\n500,1\n600,2\n"
	if err := os.WriteFile(dataPath("neptune-disc.csv"), []byte(made), 0o644); err != nil {
		t.Fatal(err)
	}
	curve, err := readSpectrum(dataPath("neptune-disc.csv"))
	if err != nil {
		t.Fatal(err)
	}
	if len(curve.nm) != 2 {
		t.Errorf("read %d points; the file on disk has 2", len(curve.nm))
	}
}

// TestDiskErrorIsSaid returns the error when a data file is there and
// cannot be read, rather than the copy built in.
func TestDiskErrorIsSaid(t *testing.T) {
	t.Chdir(t.TempDir())
	if err := os.MkdirAll(dataPath("neptune-disc.csv"), 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := readSpectrum(dataPath("neptune-disc.csv")); err == nil {
		t.Error("a data file that cannot be read was answered by the copy built in")
	}
}
