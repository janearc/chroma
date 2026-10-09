package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/janearc/libtheme-css/colourway"
	"github.com/janearc/libtheme-css/css"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// pixel is one pixel of a picture in linear light, where it is, and how
// far from the disc's centre, as a share of the disc's radius.
type pixel struct {
	r, g, b float64
	x, y    int
	out     float64
}

// sampled is a region of a picture: what it is, how it was chosen, how many
// pixels it holds, and their average in linear light.
type sampled struct {
	region, chosen string
	count          int
	r, g, b        float64
}

// sample measures a planet's picture: the disc found as everything brighter
// than space, then regions of it chosen by rule, each averaged in linear
// light. The box, when given, is where the dark spot is looked for, in
// pixels: x0 y0 x1 y1.
func sample(path string, box [4]int) ([]sampled, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	picture, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}
	disc := discOf(picture)
	inner := keep(disc, func(each pixel) bool { return each.out < 0.7 })
	byLight := append([]pixel(nil), disc...)
	sort.Slice(byLight, func(i, k int) bool {
		return luma(byLight[i]) < luma(byLight[k])
	})
	spot := keep(disc, func(each pixel) bool {
		return each.x >= box[0] && each.x < box[2] &&
			each.y >= box[1] && each.y < box[3]
	})
	sort.Slice(spot, func(i, k int) bool {
		return luma(spot[i]) < luma(spot[k])
	})
	limb := keep(disc, func(each pixel) bool {
		return each.out >= 0.9 && each.out < 0.98
	})
	darkest := fmt.Sprintf("the darkest 5%% inside x %d-%d, y %d-%d",
		box[0], box[2], box[1], box[3])
	return []sampled{
		mean("disc", "every pixel inside 70% of the radius", inner),
		mean("limb", "every pixel from 90% to 98% of the radius", limb),
		mean("clouds", "the brightest 0.5% of the disc",
			byLight[len(byLight)*995/1000:]),
		mean("dark-spot", darkest, spot[:len(spot)*5/100]),
	}, nil
}

// discOf is every pixel brighter than space, each with its distance from
// the centre of them all as a share of the radius their count implies.
func discOf(picture image.Image) []pixel {
	disc, sumX, sumY := []pixel{}, 0.0, 0.0
	bounds := picture.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := picture.At(x, y).RGBA()
			each := pixel{
				r: srgb.ToLinear(float64(r) / 65535),
				g: srgb.ToLinear(float64(g) / 65535),
				b: srgb.ToLinear(float64(b) / 65535),
				x: x, y: y,
			}
			if luma(each) > 0.01 {
				disc = append(disc, each)
				sumX, sumY = sumX+float64(x), sumY+float64(y)
			}
		}
	}
	middleX, middleY := sumX/float64(len(disc)), sumY/float64(len(disc))
	radius := math.Sqrt(float64(len(disc)) / math.Pi)
	for index := range disc {
		disc[index].out = math.Hypot(float64(disc[index].x)-middleX,
			float64(disc[index].y)-middleY) / radius
	}
	return disc
}

// luma is a pixel's luminance.
func luma(each pixel) float64 {
	return 0.2126*each.r + 0.7152*each.g + 0.0722*each.b
}

// keep is the pixels a rule holds true for.
func keep(pixels []pixel, rule func(pixel) bool) []pixel {
	kept := []pixel{}
	for _, each := range pixels {
		if rule(each) {
			kept = append(kept, each)
		}
	}
	return kept
}

// mean is pixels averaged in linear light.
func mean(region, chosen string, pixels []pixel) sampled {
	total := sampled{region: region, chosen: chosen, count: len(pixels)}
	for _, each := range pixels {
		total.r += each.r
		total.g += each.g
		total.b += each.b
	}
	share := float64(max(len(pixels), 1))
	total.r, total.g, total.b = total.r/share, total.g/share, total.b/share
	return total
}

// sheetOf is samples as a css sheet, which is what libtheme speaks: each
// region a custom property, averaged in linear light and written to the
// byte, which is the picture's own precision; where the picture came
// from, and how each region was chosen, in the comment above.
func sheetOf(name, about string, samples []sampled) colourway.Source {
	roles := css.New()
	lines := []string{strings.TrimSpace(about), ""}
	for _, one := range samples {
		roles.Set(one.region, srgb.FromLight(one.r, one.g, one.b))
		lines = append(lines, fmt.Sprintf("%s: %s, %d pixels.",
			one.region, one.chosen, one.count))
	}
	return colourway.Source{Name: name, About: strings.Join(lines, "\n"),
		Roles: roles, Origin: "data/" + name + ".css"}
}
