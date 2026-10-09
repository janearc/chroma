package main

import (
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/janearc/libtheme-css/colourway"
)

// TestPictureIsTheColourway draws a two-role colourway and finds its
// ground at the corner and its ink in the letters.
func TestPictureIsTheColourway(t *testing.T) {
	dir := t.TempDir()
	sourcePath := filepath.Join(dir, "two.css")
	sheet := ":root {\n  --ground: #e0f0e0;\n  --ink: #202830;\n}\n"
	if err := os.WriteFile(sourcePath, []byte(sheet), 0o644); err != nil {
		t.Fatal(err)
	}
	outPath := filepath.Join(dir, "two.png")
	if err := picture(sourcePath, outPath); err != nil {
		t.Fatal(err)
	}
	file, err := os.Open(outPath)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	drawn, err := png.Decode(file)
	if err != nil {
		t.Fatal(err)
	}
	ground, ink := color.RGBA{0xe0, 0xf0, 0xe0, 255}, color.RGBA{0x20, 0x28, 0x30, 255}
	if got := color.RGBAModel.Convert(drawn.At(0, 0)); got != ground {
		t.Errorf("corner is %v, the ground is %v", got, ground)
	}
	bounds, found := drawn.Bounds(), false
	for y := bounds.Min.Y; y < bounds.Max.Y && !found; y++ {
		for x := bounds.Min.X; x < bounds.Max.X && !found; x++ {
			found = color.RGBAModel.Convert(drawn.At(x, y)) == ink
		}
	}
	if !found {
		t.Errorf("no letter is drawn in the ink %v", ink)
	}
}

// TestPictureNeedsGroundAndInk refuses a colourway with no ink, and
// writes no file, rather than drawing letters nobody can see.
func TestPictureNeedsGroundAndInk(t *testing.T) {
	dir := t.TempDir()
	sourcePath := filepath.Join(dir, "groundonly.css")
	sheet := ":root {\n  --ground: #e0f0e0;\n}\n"
	if err := os.WriteFile(sourcePath, []byte(sheet), 0o644); err != nil {
		t.Fatal(err)
	}
	outPath := filepath.Join(dir, "groundonly.png")
	if err := picture(sourcePath, outPath); err == nil {
		t.Error("a colourway with no ink was drawn")
	}
	if _, err := os.Stat(outPath); err == nil {
		t.Error("a picture was written for a colourway with no ink")
	}
}

// TestPaletteFollowsGlamour reads markdown's roles the way glamour's
// dialect does: a link from link, a code block on the panel, and no
// heading ground when the colourway sets none; and the panel is in the
// picture, behind markdown's code.
func TestPaletteFollowsGlamour(t *testing.T) {
	dir := t.TempDir()
	sourcePath := filepath.Join(dir, "four.css")
	sheet := ":root {\n  --ground: #e0f0e0;\n  --ink: #202830;\n" +
		"  --link: #2050a0;\n  --panel: #c8d8f0;\n}\n"
	if err := os.WriteFile(sourcePath, []byte(sheet), 0o644); err != nil {
		t.Fatal(err)
	}
	source, err := colourway.Read(sourcePath)
	if err != nil {
		t.Fatal(err)
	}
	colours := paletteOf(source)
	link, panel := color.RGBA{0x20, 0x50, 0xa0, 255}, color.RGBA{0xc8, 0xd8, 0xf0, 255}
	if got := colours.text("md-link"); got != link {
		t.Errorf("md-link is %v, link is %v", got, link)
	}
	if got, has := colours.of("md-block-ground"); !has || got != panel {
		t.Errorf("md-block-ground is %v (%v), the panel is %v", got, has, panel)
	}
	if got, has := colours.of("md-heading-ground"); has {
		t.Errorf("md-heading-ground is %v, and the colourway sets none", got)
	}
	typeface, err := faces()
	if err != nil {
		t.Fatal(err)
	}
	defer typeface[0].Close()
	defer typeface[1].Close()
	drawn, found := paint(sample(), colours, typeface), false
	bounds := drawn.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y && !found; y++ {
		for x := bounds.Min.X; x < bounds.Max.X && !found; x++ {
			found = drawn.RGBAAt(x, y) == panel
		}
	}
	if !found {
		t.Error("markdown's code is not drawn on the panel")
	}
}
