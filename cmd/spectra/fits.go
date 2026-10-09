package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/astrogo/fitsio"
)

// science is the SCI extension of a JWST product: its values in the order
// the file keeps them, first axis fastest, its sizes, and the headers to
// read its cards from.
type science struct {
	values  []float32
	sizes   []int
	headers []*fitsio.Header
}

// readScience is a JWST product's science extension.
func readScience(path string) (science, error) {
	file, err := os.Open(path)
	if err != nil {
		return science{}, err
	}
	defer file.Close()
	fits, err := fitsio.Open(file)
	if err != nil {
		return science{}, fmt.Errorf("%s: %w", path, err)
	}
	defer fits.Close()
	image, ok := fits.Get("SCI").(fitsio.Image)
	if !ok {
		return science{}, fmt.Errorf("%s: no SCI image", path)
	}
	found := science{sizes: image.Header().Axes(),
		headers: []*fitsio.Header{image.Header(), fits.HDU(0).Header()}}
	total := 1
	for _, size := range found.sizes {
		total *= size
	}
	found.values = make([]float32, total)
	if err := image.Read(&found.values); err != nil {
		return science{}, fmt.Errorf("%s: %w", path, err)
	}
	return found, nil
}

// card is a header card by name, from the extension first and then the
// primary header, or nil.
func (found science) card(name string) *fitsio.Card {
	for _, header := range found.headers {
		if card := header.Get(name); card != nil {
			return card
		}
	}
	return nil
}

// number is a card's value as a number, or NaN when it has none.
func (found science) number(name string) float64 {
	card := found.card(name)
	if card == nil {
		return math.NaN()
	}
	switch value := card.Value.(type) {
	case float64:
		return value
	case int:
		return float64(value)
	case int64:
		return float64(value)
	}
	return math.NaN()
}

// text is a card's value as text, trimmed, or empty.
func (found science) text(name string) string {
	card := found.card(name)
	if card == nil {
		return ""
	}
	return strings.TrimSpace(fmt.Sprint(card.Value))
}

// pixelOf is where a place on the sky falls in the picture, counted from
// zero, by the gnomonic projection the header names (RA---TAN, DEC--TAN).
func (found science) pixelOf(ra, dec float64) (x, y float64) {
	radian := math.Pi / 180
	ra0 := found.number("CRVAL1") * radian
	dec0 := found.number("CRVAL2") * radian
	ra, dec = ra*radian, dec*radian
	cosine := math.Sin(dec0)*math.Sin(dec) +
		math.Cos(dec0)*math.Cos(dec)*math.Cos(ra-ra0)
	east := math.Cos(dec) * math.Sin(ra-ra0) / cosine / radian
	lift := math.Cos(dec0)*math.Sin(dec) -
		math.Sin(dec0)*math.Cos(dec)*math.Cos(ra-ra0)
	north := lift / cosine / radian
	return found.fromSky(
		east/found.number("CDELT1"), north/found.number("CDELT2"))
}

// fromSky turns steps along the sky's axes into a pixel, through the
// inverse of the header's PC matrix, counted from zero.
func (found science) fromSky(along1, along2 float64) (x, y float64) {
	a, b := found.number("PC1_1"), found.number("PC1_2")
	c, d := found.number("PC2_1"), found.number("PC2_2")
	determinant := a*d - b*c
	dx := (d*along1 - b*along2) / determinant
	dy := (-c*along1 + a*along2) / determinant
	return found.number("CRPIX1") - 1 + dx, found.number("CRPIX2") - 1 + dy
}

// arcsecOf is how far apart two pixels are on the sky, in arcseconds, east
// and north, through the header's PC matrix and steps.
func (found science) arcsecOf(dx, dy float64) (east, north float64) {
	along1 := found.number("PC1_1")*dx + found.number("PC1_2")*dy
	along2 := found.number("PC2_1")*dx + found.number("PC2_2")*dy
	east = found.number("CDELT1") * along1
	north = found.number("CDELT2") * along2
	return east * 3600, north * 3600
}

// pixelsOf is a distance on the sky, arcseconds east and north, in pixels.
func (found science) pixelsOf(east, north float64) (dx, dy float64) {
	x, y := found.fromSky(east/3600/found.number("CDELT1"),
		north/3600/found.number("CDELT2"))
	return x - (found.number("CRPIX1") - 1),
		y - (found.number("CRPIX2") - 1)
}

// target is where the moving target, here neptune, falls in the picture.
func (found science) target() (x, y float64) {
	return found.pixelOf(found.number("MT_RA"), found.number("MT_DEC"))
}

// scale is the side of one pixel, in arcseconds.
func (found science) scale() float64 {
	return math.Sqrt(found.number("PIXAR_A2"))
}
