package main

import (
	"bufio"
	"bytes"
	"embed"
	"errors"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/janearc/libtheme-css/primitives/bands"
	"github.com/janearc/libtheme-css/primitives/swatch"
)

// spectrum is values by wavelength in nanometres, as a data file gives
// them, read between its points by a straight line and not at all outside.
type spectrum struct {
	nm, value []float64
}

// at is the spectrum at a wavelength, and whether the data reaches it.
func (curve spectrum) at(nm float64) (float64, bool) {
	if len(curve.nm) == 0 || nm < curve.nm[0] ||
		nm > curve.nm[len(curve.nm)-1] {
		return 0, false
	}
	next := sort.SearchFloat64s(curve.nm, nm)
	if next == 0 || curve.nm[next] == nm {
		return curve.value[next], true
	}
	share := (nm - curve.nm[next-1]) / (curve.nm[next] - curve.nm[next-1])
	return curve.value[next-1] +
		share*(curve.value[next]-curve.value[next-1]), true
}

// reaches is whether the spectrum's data runs from one wavelength to
// another. A colour made from less of visible than 400 to 700 nm is the
// colour of what was measured, not of the thing, so no chip is drawn for it.
func (curve spectrum) reaches(lo, hi float64) bool {
	return len(curve.nm) > 0 && curve.nm[0] <= lo &&
		curve.nm[len(curve.nm)-1] >= hi
}

// times is two spectra multiplied where both have data.
func (curve spectrum) times(other spectrum) spectrum {
	product := spectrum{}
	for index, nm := range curve.nm {
		if value, has := other.at(nm); has {
			product.nm = append(product.nm, nm)
			both := curve.value[index] * value
			product.value = append(product.value, both)
		}
	}
	return product
}

// scaledBy is the spectrum with every value multiplied by k; a negative k
// turns light taken away into light.
func (curve spectrum) scaledBy(k float64) spectrum {
	turned := spectrum{
		nm: curve.nm, value: make([]float64, len(curve.value)),
	}
	for index, value := range curve.value {
		turned.value[index] = max(value*k, 0)
	}
	return turned
}

// daylight is D65 as a spectrum, to light a reflectance by.
func daylight() spectrum {
	table, curve := swatch.D65(), spectrum{}
	for nm := range table {
		curve.nm = append(curve.nm, float64(nm))
	}
	sort.Float64s(curve.nm)
	for _, nm := range curve.nm {
		curve.value = append(curve.value, table[int(nm)])
	}
	return curve
}

// readSpectrum is a data file's first two columns, wavelength and value,
// skipping comments and the header line.
func readSpectrum(path string) (spectrum, error) {
	raw, err := readData(path)
	if err != nil {
		return spectrum{}, err
	}
	curve, lines := spectrum{}, bufio.NewScanner(bytes.NewReader(raw))
	for lines.Scan() {
		fields := strings.Split(lines.Text(), ",")
		nm, errNm := strconv.ParseFloat(fields[0], 64)
		if strings.HasPrefix(lines.Text(), "#") || len(fields) < 2 ||
			errNm != nil {
			continue
		}
		value, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			return spectrum{}, err
		}
		curve.nm = append(curve.nm, nm)
		curve.value = append(curve.value, value)
	}
	return curve, lines.Err()
}

// strip is a light spread across visible as an eye would see it, one
// colour for each of a number of columns even in the logarithm of
// wavelength, exposed alike; empty where the data does not reach.
func (curve spectrum) strip(columns int) []string {
	lo, hi := swatch.Observed()
	seen, has := make([]swatch.Swatch, columns), make([]bool, columns)
	for index := range seen {
		nm := float64(lo) * math.Pow(float64(hi)/float64(lo),
			float64(index)/float64(columns-1))
		value, reaches := curve.at(nm)
		wave := swatch.Monochrome(int(math.Round(nm)))
		seen[index], has[index] = scaled(wave, value), reaches
	}
	exposure, colours := exposed(seen, 0.4), make([]string, columns)
	for index := range seen {
		if has[index] {
			colours[index] = hex(scaled(seen[index], exposure))
		}
	}
	return colours
}

// colour is the whole light as one colour, at one lightness.
func (curve spectrum) colour() string {
	lines := map[int]float64{}
	last := curve.nm[len(curve.nm)-1]
	for nm := float64(int(curve.nm[0]) + 1); nm <= last; nm++ {
		if value, has := curve.at(nm); has {
			lines[int(nm)] = value
		}
	}
	return hex(atLightness(bands.Split(lines).Swatch(), 0.72))
}

// data is the data folder as it was when the program was built, so an
// installed spectra has its spectra with it.
//
//go:embed data
var data embed.FS

// dataPath is a data file in data/, in the folder spectra runs in; run
// from cmd/spectra, that is the repository's own data folder.
func dataPath(name string) string {
	return filepath.Join("data", name)
}

// readData reads a data file from disk when it is there, so a file just
// made by webb is the one read, and from the copy built in when it is not.
// A file that is there and cannot be read is an error, not the old copy.
func readData(path string) ([]byte, error) {
	raw, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return data.ReadFile(filepath.ToSlash(path))
	}
	return raw, err
}

// inBand is the spectrum's mean between two wavelengths, as the power in
// that band, and whether it has data there.
func (curve spectrum) inBand(lo, hi float64) (float64, bool) {
	sum, count := 0.0, 0
	for index, nm := range curve.nm {
		if nm >= lo && nm < hi {
			sum, count = sum+curve.value[index], count+1
		}
	}
	if count == 0 {
		return 0, false
	}
	return sum / float64(count), true
}

// perLogarithm is a spectrum given per unit of frequency, as janskys are,
// as power in equal steps of the logarithm of wavelength: divided by the
// wavelength.
func (curve spectrum) perLogarithm() spectrum {
	turned := spectrum{
		nm: curve.nm, value: make([]float64, len(curve.value)),
	}
	for index, value := range curve.value {
		turned.value[index] = max(value, 0) / curve.nm[index]
	}
	return turned
}
