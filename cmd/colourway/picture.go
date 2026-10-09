package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"unicode/utf8"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/gomonobold"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	"github.com/janearc/libtheme-css/colourway"
	"github.com/janearc/libtheme-css/dialects/glamour"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// pictureSize is the type's size in pixels, and pictureColumns the
// fewest columns a picture is drawn at, so every picture is one width.
const (
	pictureSize    = 28
	pictureColumns = 64
)

// span is a run of the sample in one colour: its text, the role it is
// drawn in, the role behind it when it has a ground of its own, and
// whether it is bold.
type span struct {
	text, ink, ground string
	bold              bool
}

// in is text in a role's colour.
func in(ink, text string) span {
	return span{text: text, ink: ink}
}

// strong is text in a role's colour, bold.
func strong(ink, text string) span {
	return span{text: text, ink: ink, bold: true}
}

// on is text in a role's colour on a ground of its own.
func on(ground, ink, text string) span {
	return span{text: text, ink: ink, ground: ground}
}

// line is a line of the sample, and the ground the whole line is drawn
// on when it is part of a block, like a code block in glow.
type line struct {
	ground string
	spans  []span
}

// plain is a line on the colourway's own ground.
func plain(spans ...span) line {
	return line{spans: spans}
}

// block is a line on a ground of its own, across the text, inside the
// margins.
func block(ground string, spans ...span) line {
	return line{ground: ground, spans: spans}
}

// sixteen is the terminal's eight colours, normal or bright, each named
// in its own colour.
func sixteen(label, prefix string) line {
	spans := []span{in("ink", label)}
	for _, name := range []string{"black", "red", "green", "yellow",
		"blue", "magenta", "cyan", "white"} {
		spans = append(spans, in(prefix+name, name), in("ink", " "))
	}
	return plain(spans...)
}

// sample is what every picture shows: a shell, where the sixteen
// colours are, then a page of markdown as glow draws it, with a code
// block, so each colourway is seen where it lands.
func sample() []line {
	prompt := strong("ink", "~/chroma % ")
	code := func(spans ...span) line {
		return block("md-block-ground", spans...)
	}
	return []line{
		plain(prompt, in("ink", "colours")),
		sixteen("normal ", ""),
		sixteen("bright ", "bright-"),
		plain(prompt, in("ink", "git log --oneline -3")),
		plain(in("yellow", "3f1c0a2 ("), strong("cyan", "HEAD -> "),
			strong("green", "main"), in("yellow", ")"),
			in("ink", " colourways: a picture of each")),
		plain(in("yellow", "9b27e41"),
			in("ink", " libreadme: contrast is measured")),
		plain(in("yellow", "51d0c8e"),
			in("ink", " libtheme: colour, decomposed")),
		plain(prompt, in("ink", "ls")),
		plain(in("blue", "cmd"), in("ink", "  "),
			in("blue", "colourways"), in("ink", "  "),
			in("blue", "libreadme"), in("ink", "  "),
			in("blue", "libtheme"),
			in("ink", "  go.mod  LICENSE  README.md")),
		plain(prompt, in("ink", "go test ./...")),
		plain(in("ink", "ok    chroma/libtheme/spaces/ok    0.214s")),
		plain(in("red", "--- FAIL: TestInkOnGround (0.01s)")),
		plain(),
		plain(on("md-heading-ground", "md-heading", " chroma ")),
		plain(),
		plain(in("md-body", "Colourways for reading at length. "),
			in("md-link", "libtheme"),
			in("md-body", " takes colour")),
		plain(in("md-body", "apart; "),
			on("md-code-ground", "md-code", " libreadme "),
			in("md-body", " measures what is readable.")),
		plain(),
		plain(in("md-quote",
			"| Every colour is a role, and every role a source.")),
		plain(),
		code(in("md-comment",
			"  // ground is the paper the ink sits on.")),
		code(in("md-keyword", "  func "), in("md-function", "ground"),
			in("md-operator", "("), in("md-name", "ink "),
			in("md-type", "Swatch"), in("md-operator", ") "),
			in("md-type", "Swatch"), in("md-operator", " {")),
		code(in("md-keyword", "      return "), in("md-name", "ink"),
			in("md-operator", "."), in("md-function", "Toward"),
			in("md-operator", "("), in("md-name", "white"),
			in("md-operator", ", "), in("md-number", "0.85"),
			in("md-operator", ")")),
		code(in("md-operator", "  }")),
	}
}

// palette is a colourway's roles as the picture reads them: its own,
// then the general vocabulary, then glamour's fallbacks.
type palette struct {
	resolved  colourway.Resolved
	fallbacks map[string][]string
}

// paletteOf is a source's roles resolved for the picture.
func paletteOf(source colourway.Source) palette {
	fallbacks := map[string][]string{}
	for _, role := range glamour.Roles {
		fallbacks[role.Name] = role.Fallbacks
	}
	return palette{colourway.Resolve(source.Roles), fallbacks}
}

// of is a role's colour: the role itself, else its fallbacks in order,
// and whether there is one. It walks the glamour dialect's own table, which
// has no cycles; the dialect's resolve folds the ink in as a last resort,
// and a ground must not have one, so the walk is kept here.
func (colours palette) of(name string) (color.RGBA, bool) {
	if value, has := colours.resolved.Of(name); has {
		lit, _ := srgb.FromSwatch(value)
		r, g, b := lit.Bytes()
		return color.RGBA{r, g, b, 255}, true
	}
	for _, next := range colours.fallbacks[name] {
		if found, has := colours.of(next); has {
			return found, true
		}
	}
	return color.RGBA{}, false
}

// text is a role's colour for letters, which end at the ink.
func (colours palette) text(name string) color.RGBA {
	if found, has := colours.of(name); has {
		return found
	}
	found, _ := colours.of("ink")
	return found
}

// faces are the regular and bold type, Go Mono, which ships with Go's
// own image library, so a picture is drawn the same on every machine.
func faces() ([2]font.Face, error) {
	var both [2]font.Face
	for index, ttf := range [][]byte{gomono.TTF, gomonobold.TTF} {
		parsed, err := opentype.Parse(ttf)
		if err != nil {
			return both, err
		}
		options := &opentype.FaceOptions{
			Size:    pictureSize,
			DPI:     72,
			Hinting: font.HintingFull,
		}
		both[index], err = opentype.NewFace(parsed, options)
		if err != nil {
			if index > 0 {
				both[0].Close()
			}
			return both, err
		}
	}
	return both, nil
}

// paint draws the sample in a colourway, a cell per character, on the
// colourway's ground, with a margin of two cells and one line.
func paint(lines []line, colours palette, typeface [2]font.Face) *image.RGBA {
	metrics := typeface[0].Metrics()
	width := font.MeasureString(typeface[0], "M").Ceil()
	height := metrics.Height.Ceil()
	columns := pictureColumns
	for _, each := range lines {
		count := 0
		for _, run := range each.spans {
			count += utf8.RuneCountInString(run.text)
		}
		columns = max(columns, count)
	}
	canvas := image.NewRGBA(image.Rect(0, 0,
		(columns+4)*width, (len(lines)+2)*height))
	draw.Draw(canvas, canvas.Bounds(),
		image.NewUniform(colours.text("ground")),
		image.Point{}, draw.Src)
	for row, each := range lines {
		top := (row + 1) * height
		if ground, has := groundOf(colours, each.ground); has {
			fill(canvas, image.Rect(2*width, top,
				(columns+2)*width, top+height), ground)
		}
		column := 2
		for _, run := range each.spans {
			count := utf8.RuneCountInString(run.text)
			if ground, has := groundOf(colours, run.ground); has {
				left := column * width
				right := left + count*width
				area := image.Rect(left, top, right, top+height)
				fill(canvas, area, ground)
			}
			face := typeface[0]
			if run.bold {
				face = typeface[1]
			}
			baseline := top + metrics.Ascent.Ceil()
			pen := font.Drawer{Dst: canvas,
				Src:  image.NewUniform(colours.text(run.ink)),
				Face: face,
				Dot:  fixed.P(column*width, baseline)}
			pen.DrawString(run.text)
			column += count
		}
	}
	return canvas
}

// groundOf is the colour behind a line or a run, when it names a ground
// and the colourway has one; none is drawn otherwise.
func groundOf(colours palette, name string) (color.RGBA, bool) {
	if name == "" {
		return color.RGBA{}, false
	}
	return colours.of(name)
}

// fill paints a rectangle in one colour.
func fill(canvas *image.RGBA, area image.Rectangle, colour color.RGBA) {
	draw.Draw(canvas, area, image.NewUniform(colour),
		image.Point{}, draw.Src)
}

// picture draws a colourway's source as the sample and writes it as a
// png.
func picture(sourcePath, outPath string) error {
	source, err := colourway.Read(sourcePath)
	if err != nil {
		return err
	}
	colours := paletteOf(source)
	for _, needed := range []string{"ground", "ink"} {
		if _, has := colours.of(needed); !has {
			return fmt.Errorf(
				"%s has no %s, so there is nothing to draw",
				sourcePath, needed)
		}
	}
	typeface, err := faces()
	if err != nil {
		return err
	}
	defer typeface[0].Close()
	defer typeface[1].Close()
	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	drawing := paint(sample(), colours, typeface)
	if err := png.Encode(out, drawing); err != nil {
		out.Close()
		os.Remove(outPath)
		return err
	}
	return out.Close()
}
