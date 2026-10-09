package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

// nircamFilters are the four filters webb pictured neptune through on 12
// july 2022, program 2739, each with its pivot wavelength in nanometres
// from the JWST user documentation's NIRCam filter table.
var nircamFilters = []struct {
	name  string
	pivot float64
}{{"f140m", 1404}, {"f210m", 2093}, {"f300m", 2989}, {"f460m", 4630}}

// picture is a nircam image and where neptune's centre falls in it.
type picture struct {
	science
	width, height int
	x, y          float64
}

// readPicture is a nircam image from its file.
func readPicture(path string) (picture, error) {
	found, err := readScience(path)
	if err != nil {
		return picture{}, err
	}
	if len(found.sizes) != 2 {
		return picture{}, fmt.Errorf("%s: not an image", path)
	}
	each := picture{
		science: found, width: found.sizes[0], height: found.sizes[1],
	}
	each.x, each.y = found.target()
	return each, nil
}

// within is the pixels no further than a distance from a point, in
// arcseconds, as indices.
func (each picture) within(x, y, arcsec float64) []int {
	reach := arcsec / each.scale()
	pixels := []int{}
	top, bottom := max(int(y-reach), 0), min(int(y+reach), each.height-1)
	left, right := max(int(x-reach), 0), min(int(x+reach), each.width-1)
	for row := top; row <= bottom; row++ {
		for col := left; col <= right; col++ {
			if math.Hypot(float64(col)-x, float64(row)-y) <= reach {
				pixels = append(pixels, row*each.width+col)
			}
		}
	}
	return pixels
}

// meanOver is the mean of the picture at some pixels, leaving out those
// with no value.
func (each picture) meanOver(pixels []int) float64 {
	sum, count := 0.0, 0
	for _, pixel := range pixels {
		if value := float64(each.values[pixel]); !math.IsNaN(value) {
			sum, count = sum+value, count+1
		}
	}
	return sum / float64(count)
}

// shares is the pixels of a region whose value is in its brightest share,
// and those in its darkest.
func (each picture) shares(
	region []int, top, bottom float64,
) (bright, dark []int) {
	values := []float64{}
	for _, pixel := range region {
		if value := float64(each.values[pixel]); !math.IsNaN(value) {
			values = append(values, value)
		}
	}
	sort.Float64s(values)
	above := values[int(float64(len(values)-1)*(1-top))]
	below := values[int(float64(len(values)-1)*bottom)]
	for _, pixel := range region {
		value := float64(each.values[pixel])
		if value >= above {
			bright = append(bright, pixel)
		} else if value <= below {
			dark = append(dark, pixel)
		}
	}
	return bright, dark
}

// carried is a region chosen in one picture, found again in another by
// where each pixel sits on the sky from neptune's centre.
func carried(region []int, from, to picture) []int {
	chosen := map[int]bool{}
	for _, pixel := range region {
		across := float64(pixel%from.width) - from.x
		down := float64(pixel/from.width) - from.y
		east, north := from.arcsecOf(across, down)
		dx, dy := to.pixelsOf(east, north)
		x, y := int(math.Round(to.x+dx)), int(math.Round(to.y+dy))
		if x >= 0 && y >= 0 && x < to.width && y < to.height {
			chosen[y*to.width+x] = true
		}
	}
	pixels := []int{}
	for pixel := range chosen {
		pixels = append(pixels, pixel)
	}
	return pixels
}

// nircamRegion is what each region is, for its file's comments.
var nircamRegion = map[string]string{
	"disc": "the inner 85% of the disc's radius",
	"clouds": "the disc's brightest 15% through F140M, " +
		"found again in the other filters by place: high cloud",
	"clear": "the disc's darker half through F140M, " +
		"found again in the other filters by place",
}

// fromNircam makes neptune in four numbers out of the four pictures in a
// folder: the inner disc, its high clouds and its clear half, each the
// mean over the region; and writes them to data. Triton is left out: its
// core saturated in the short filters.
func fromNircam(folder string) error {
	regions, used := map[string]spectrum{}, []string{}
	var first picture
	var clouds, clear []int
	for index, filter := range nircamFilters {
		name := "jw02739-o004_t003_nircam_clear-" + filter.name
		each, err := readPicture(fileIn(folder, name))
		if err != nil {
			return err
		}
		disc := each.within(each.x, each.y, inner*neptuneRadius)
		if index == 0 {
			first = each
			clouds, clear = each.shares(disc, 0.15, 0.5)
		}
		found := map[string][]int{
			"disc":   disc,
			"clouds": carried(clouds, first, each),
			"clear":  carried(clear, first, each),
		}
		for name, region := range found {
			curve := regions[name]
			curve.nm = append(curve.nm, filter.pivot)
			curve.value = append(curve.value, each.meanOver(region))
			regions[name] = curve
		}
		fmt.Fprintf(os.Stderr, "%s: neptune at %.0f,%.0f\n",
			filter.name, each.x, each.y)
		used = append(used, provenance(each.science))
	}
	source := "source: JWST NIRCam, program 2739, " +
		"filters F140M F210M F300M F460M, 2022-07-12; " +
		"each value at its filter's pivot wavelength"
	for name, curve := range regions {
		comments := append([]string{
			"target: Neptune, " + nircamRegion[name], source,
			"archive: MAST, calibration level 3 images:",
		}, used...)
		comments = append(comments,
			"units: surface brightness in MJy/sr, "+
				"mean over the region")
		csv := "neptune-nircam-" + name + ".csv"
		err := writeData(csv, comments, "mjy_per_sr", curve)
		if err != nil {
			return err
		}
	}
	return nil
}

// writeData writes a spectrum to a file in data and says so.
func writeData(
	name string, comments []string, column string, curve spectrum,
) error {
	path := dataPath(name)
	if err := writeSpectrum(path, comments, column, curve); err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "wrote", path)
	return nil
}
