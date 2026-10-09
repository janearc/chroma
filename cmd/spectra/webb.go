package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// neptuneRadius is neptune's equatorial radius as seen from webb in 2022
// and 2023, in arcseconds: 24,764 km at about 29.3 au.
const neptuneRadius = 1.16

// inner is how much of the disc's radius a region keeps, so that the
// dimmed and blurred limb is left out.
const inner = 0.85

// cube is a nirspec picture at every wavelength, and where neptune's
// centre falls in it.
type cube struct {
	science
	width, height, depth int
	x, y                 float64
}

// readCube is a nirspec cube from its file.
func readCube(path string) (cube, error) {
	found, err := readScience(path)
	if err != nil {
		return cube{}, err
	}
	if len(found.sizes) != 3 {
		return cube{}, fmt.Errorf("%s: not a cube", path)
	}
	each := cube{
		science: found, width: found.sizes[0], height: found.sizes[1],
		depth: found.sizes[2],
	}
	each.x, each.y = found.target()
	return each, nil
}

// micron is the wavelength of one of the cube's planes.
func (each cube) micron(plane int) float64 {
	along := float64(plane) + 1 - each.number("CRPIX3")
	return each.number("CRVAL3") + along*each.number("CDELT3")
}

// image is the cube's mean light between two wavelengths, one value for
// each spot, NaN where it has none.
func (each cube) image(lo, hi float64) []float64 {
	spots := each.width * each.height
	sums, counts := make([]float64, spots), make([]int, spots)
	for plane := range each.depth {
		if micron := each.micron(plane); micron < lo || micron > hi {
			continue
		}
		slab := each.values[plane*spots : (plane+1)*spots]
		for spot, value := range slab {
			if !math.IsNaN(float64(value)) {
				sums[spot] += float64(value)
				counts[spot]++
			}
		}
	}
	for spot := range sums {
		sums[spot] /= float64(counts[spot])
	}
	return sums
}

// disc is the spots inside the inner part of neptune's disc.
func (each cube) disc() []bool {
	chosen := make([]bool, each.width*each.height)
	reach := inner * neptuneRadius / each.scale()
	for spot := range chosen {
		x, y := spot%each.width, spot/each.width
		away := math.Hypot(float64(x)-each.x, float64(y)-each.y)
		chosen[spot] = away <= reach
	}
	return chosen
}

// split is the spots of a region whose value in a picture is above one
// share of them, and those below another: the brightest tenth, say, and
// the darker half.
func split(
	region []bool, picture []float64, top, bottom float64,
) (bright, dark []bool) {
	values := []float64{}
	for spot, in := range region {
		if in && !math.IsNaN(picture[spot]) {
			values = append(values, picture[spot])
		}
	}
	sort.Float64s(values)
	above := values[int(float64(len(values)-1)*(1-top))]
	below := values[int(float64(len(values)-1)*bottom)]
	bright, dark = make([]bool, len(region)), make([]bool, len(region))
	for spot, in := range region {
		bright[spot] = in && picture[spot] >= above
		dark[spot] = in && picture[spot] <= below
	}
	return bright, dark
}

// moved is a region carried from one cube onto another by where neptune's
// centre falls in each, to the nearest spot.
func moved(region []bool, from, to cube) []bool {
	dx := int(math.Round(to.x - from.x))
	dy := int(math.Round(to.y - from.y))
	carried := make([]bool, to.width*to.height)
	for spot, in := range region {
		x, y := spot%from.width+dx, spot/from.width+dy
		if in && x >= 0 && y >= 0 && x < to.width && y < to.height {
			carried[y*to.width+x] = true
		}
	}
	return carried
}

// spectrumOf is a region's mean light at each of the cube's wavelengths
// from one to another, in nanometres and MJy/sr; a wavelength where none
// of the region has a value is left out.
func (each cube) spectrumOf(region []bool, lo, hi float64) spectrum {
	spots, curve := each.width*each.height, spectrum{}
	for plane := range each.depth {
		micron := each.micron(plane)
		if micron < lo || micron >= hi {
			continue
		}
		sum, count := 0.0, 0
		slab := each.values[plane*spots : (plane+1)*spots]
		for spot, value := range slab {
			if region[spot] && !math.IsNaN(float64(value)) {
				sum, count = sum+float64(value), count+1
			}
		}
		if count > 0 {
			curve.nm = append(curve.nm, micron*1000)
			curve.value = append(curve.value, sum/float64(count))
		}
	}
	return curve
}

// joined is two spectra end to end, the first up to where the second
// takes over.
func joined(first, second spectrum, at float64) spectrum {
	curve := spectrum{}
	for index, nm := range first.nm {
		if nm < at {
			curve.nm = append(curve.nm, nm)
			curve.value = append(curve.value, first.value[index])
		}
	}
	for index, nm := range second.nm {
		if nm >= at {
			curve.nm = append(curve.nm, nm)
			curve.value = append(curve.value, second.value[index])
		}
	}
	return curve
}

// meanOf is the mean of spectra taken on one grid of wavelengths, where
// every one of them has a value.
func meanOf(curves ...spectrum) spectrum {
	mean := spectrum{}
	for _, nm := range curves[0].nm {
		sum := 0.0
		for _, curve := range curves {
			value, has := curve.at(nm)
			if !has {
				sum = math.NaN()
			}
			sum += value
		}
		if !math.IsNaN(sum) {
			mean.nm = append(mean.nm, nm)
			average := sum / float64(len(curves))
			mean.value = append(mean.value, average)
		}
	}
	return mean
}

// writeSpectrum writes a spectrum as a data file: its comment lines, a
// header line, then wavelength and value.
func writeSpectrum(
	path string, comments []string, column string, curve spectrum,
) error {
	var text strings.Builder
	for _, line := range comments {
		text.WriteString("# " + line + "\n")
	}
	text.WriteString("nm," + column + "\n")
	for index, nm := range curve.nm {
		fmt.Fprintf(&text, "%.2f,%.6g\n", nm, curve.value[index])
	}
	return os.WriteFile(path, []byte(text.String()), 0o644)
}

// fileIn is a file in a folder by the start of its name.
func fileIn(folder, prefix string) string {
	found, _ := filepath.Glob(filepath.Join(folder, prefix+"*"))
	if len(found) == 0 {
		return filepath.Join(folder, prefix)
	}
	return found[0]
}
