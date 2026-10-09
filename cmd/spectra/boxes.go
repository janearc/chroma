package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"

	"github.com/janearc/libtheme-css/spaces/srgb"
)

// box is a region of a picture to sample, by its corners in pixels.
type box struct {
	region         string
	x0, y0, x1, y1 int
}

// toLinear is every byte's level in linear light, worked out once.
var toLinear = func() [256]float64 {
	table := [256]float64{}
	for level := range table {
		table[level] = srgb.ToLinear(float64(level) / 255)
	}
	return table
}()

// boxes samples a picture: the disc inside 70% of its radius, then each
// box, every pixel brighter than space averaged in linear light. It reads
// the pixels where they lie rather than gathering them, so a picture of
// 8000 by 8000 costs one pass, not a copy of itself.
func boxes(path string, regions []box) ([]sampled, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	picture, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	at := levelsOf(picture)
	bounds := picture.Bounds()
	count, sumX, sumY := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if luma(at(x, y)) > 0.01 {
				count++
				sumX += float64(x)
				sumY += float64(y)
			}
		}
	}
	middleX, middleY := sumX/count, sumY/count
	radius := math.Sqrt(count / math.Pi)
	inner := box{"disc", int(middleX - radius), int(middleY - radius),
		int(middleX + radius), int(middleY + radius)}
	samples := []sampled{averaged(at, inner, func(x, y int) bool {
		away := math.Hypot(float64(x)-middleX, float64(y)-middleY)
		return away < 0.7*radius
	})}
	samples[0].chosen = "every pixel inside 70% of the radius"
	for _, each := range regions {
		one := averaged(at, each, func(int, int) bool { return true })
		one.chosen = fmt.Sprintf("every pixel in x %d-%d, y %d-%d",
			each.x0, each.x1, each.y0, each.y1)
		samples = append(samples, one)
	}
	return samples, nil
}

// averaged is a box's pixels that a rule keeps, space left out, averaged
// in linear light.
func averaged(
	at func(x, y int) pixel, region box, keep func(x, y int) bool,
) sampled {
	total := sampled{region: region.region}
	for y := region.y0; y < region.y1; y++ {
		for x := region.x0; x < region.x1; x++ {
			each := at(x, y)
			if luma(each) > 0.01 && keep(x, y) {
				total.r += each.r
				total.g += each.g
				total.b += each.b
				total.count++
			}
		}
	}
	share := float64(max(total.count, 1))
	total.r, total.g, total.b = total.r/share, total.g/share, total.b/share
	return total
}

// levelsOf is a way to read a picture's pixel in linear light: straight
// from its bytes when it is the usual kind, through image.At when not.
func levelsOf(picture image.Image) func(x, y int) pixel {
	if rgba, ok := picture.(*image.RGBA); ok {
		return func(x, y int) pixel {
			if !(image.Point{x, y}.In(rgba.Rect)) {
				return pixel{}
			}
			at := rgba.PixOffset(x, y)
			return pixel{
				r: toLinear[rgba.Pix[at]],
				g: toLinear[rgba.Pix[at+1]],
				b: toLinear[rgba.Pix[at+2]],
				x: x, y: y,
			}
		}
	}
	return func(x, y int) pixel {
		r, g, b, _ := picture.At(x, y).RGBA()
		return pixel{
			r: toLinear[r>>8],
			g: toLinear[g>>8],
			b: toLinear[b>>8],
			x: x, y: y,
		}
	}
}
