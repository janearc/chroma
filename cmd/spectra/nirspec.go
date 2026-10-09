package main

import (
	"fmt"
	"os"
)

// nirspecFaces are the two observations of program 1249 that pointed
// nirspec's cube at neptune, on 22 june 2023, about eight hours apart, so
// half a turn of the planet apart: two faces.
var nirspecFaces = []string{"o005", "o006"}

// methaneBand is where in the cube's first grating the clouds are chosen:
// methane's band at 2.30 to 2.35 microns, dark enough that light from deep
// down does not get back out, so what is bright there is high cloud.
var methaneBand = [2]float64{2.30, 2.35}

// fromNirspec makes neptune's spectra, from 1.66 to 5.27 microns, out of
// the four cubes in a folder: the inner disc, its high clouds and its clear
// half, each the mean of the two faces; and writes them to data.
func fromNirspec(folder string) error {
	faces := map[string][]spectrum{}
	used := []string{}
	for _, observation := range nirspecFaces {
		found, files, err := nirspecFace(folder, observation)
		if err != nil {
			return err
		}
		for name, curve := range found {
			faces[name] = append(faces[name], curve)
		}
		used = append(used, files...)
	}
	source := "source: JWST NIRSpec integral field unit, program 1249, " +
		"gratings G235H (F170LP) and G395H (F290LP), " +
		"2023-06-22, the two faces averaged"
	units := "units: surface brightness in MJy/sr, mean over the region; " +
		"G235H below 3000 nm, G395H from 3000 nm"
	for _, name := range []string{"disc", "clouds", "clear"} {
		comments := append([]string{
			"target: Neptune, " + nirspecRegion[name],
			source,
			"archive: MAST, calibration level 3 cubes:"},
			used...)
		comments = append(comments, units)
		path := dataPath("neptune-nirspec-" + name + ".csv")
		mean := meanOf(faces[name]...)
		err := writeSpectrum(path, comments, "mjy_per_sr", mean)
		if err != nil {
			return err
		}
		fmt.Fprintln(os.Stderr, "wrote", path)
	}
	return nil
}

// nirspecRegion is what each region is, for its file's comments.
var nirspecRegion = map[string]string{
	"disc": "the inner 85% of the disc's radius",
	"clouds": "the disc's brightest 15% in methane's band at " +
		"2.30 to 2.35 microns: high cloud",
	"clear": "the disc's darker half in methane's band at " +
		"2.30 to 2.35 microns",
}

// nirspecFace is one face's regions as spectra, and the files they came
// from.
func nirspecFace(
	folder, observation string,
) (map[string]spectrum, []string, error) {
	prefix := "jw01249-" + observation + "_t001_nirspec_"
	short, err := readCube(fileIn(folder, prefix+"g235h"))
	if err != nil {
		return nil, nil, err
	}
	long, err := readCube(fileIn(folder, prefix+"g395h"))
	if err != nil {
		return nil, nil, err
	}
	disc := short.disc()
	methane := short.image(methaneBand[0], methaneBand[1])
	clouds, clear := split(disc, methane, 0.15, 0.5)
	regions := map[string][]bool{
		"disc": disc, "clouds": clouds, "clear": clear,
	}
	found := map[string]spectrum{}
	for name, region := range regions {
		low := short.spectrumOf(region, 0, 3.0)
		high := long.spectrumOf(moved(region, short, long), 3.0, 99)
		found[name] = joined(low, high, 3000)
	}
	fmt.Fprintf(os.Stderr, "%s: centre %.1f,%.1f and %.1f,%.1f; "+
		"spots %d disc, %d clouds, %d clear\n",
		observation, short.x, short.y, long.x, long.y,
		count(disc), count(clouds), count(clear))
	files := []string{provenance(short.science), provenance(long.science)}
	return found, files, nil
}

// provenance is a product's file name and the pipeline version that made
// it, as its header gives them.
func provenance(found science) string {
	return "  " + found.text("FILENAME") +
		", pipeline " + found.text("CAL_VER")
}

// count is how many of a region's spots are in it.
func count(region []bool) int {
	total := 0
	for _, in := range region {
		if in {
			total++
		}
	}
	return total
}
