package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/x/ansi"
)

// page is one screen of paratune: a sample of something she reads, drawn
// on its own ground, and the stops that colour it.
type page int

const (
	shellPage page = iota
	claudePage
	markdownPage
	vimPage
)

// pageNames are the pages as the menu and the steering line name them.
var pageNames = []string{"ghostty", "claude code", "markdown", "vim"}

// groundName is the role a page is drawn on: the terminal's ground,
// which claude code and glamour draw on too, or vim's own if the
// colourway gives it one.
func (set *tuning) groundName(which page) string {
	if _, own := set.roles.get("nvim-ground"); own && which == vimPage {
		return "nvim-ground"
	}
	return "ground"
}

// inkName is the role a page's text is mostly in, which its controls
// start from.
func (set *tuning) inkName(which page) string {
	switch which {
	case claudePage:
		return "claude-text"
	case markdownPage:
		return "md-body"
	case vimPage:
		if _, own := set.roles.get("nvim-ink"); own {
			return "nvim-ink"
		}
	}
	return "ink"
}

// line is one line of a sample, the names of the stops whose colours it
// shows, so the one being steered can be pointed at, and the escape for
// the band the rest of the line is filled with, when it has one of its
// own; empty means the page's ground.
type line struct {
	text  string
	shows []string
	fill  string
}

// escapes are terminal colour codes, taken out before matching words:
// glamour draws a code between words.
var escapes = regexp.MustCompile("\x1b\\[[0-9;:]*m")

// resets is whether a colour code's parameters leave the terminal's own
// ground showing: a full reset, or 49, the default ground, with no ground
// set after it in the same code. the numbers that follow a 38, 48 or 58
// are a colour, not codes, and are stepped over.
func resets(parameters string) bool {
	parts := strings.Split(parameters, ";")
	reset := false
	for position := 0; position < len(parts); position++ {
		switch parts[position] {
		case "", "0", "49":
			reset = true
		case "48":
			reset = false
			fallthrough
		case "38", "58":
			following := ""
			if position+1 < len(parts) {
				following = parts[position+1]
			}
			if following == "5" {
				position += 2
			} else if following == "2" {
				position += 4
			}
		}
	}
	return reset
}

// onGround puts a ground back after every code in text that would show
// the terminal's own, so a band stays one colour across whatever glamour,
// chroma, lipgloss or huh drew into it.
func onGround(text, behind string) string {
	return escapes.ReplaceAllStringFunc(text, func(code string) string {
		bare := strings.TrimPrefix(code, "\x1b[")
		if resets(strings.TrimSuffix(bare, "m")) {
			return code + behind
		}
		return code
	})
}

// explicit is a line with every colour it names by number, and the
// terminal's own ink, said as the colourway's own colours: the sixteen by
// their codes, any palette number by 38;5 and 48;5, and the default ink
// after a reset or a 39.
//
// A page then shows the colourway whatever the pane's palette and style
// hold, which anything else may have reset.
func (set *tuning) explicit(text string) string {
	ink := "38;2;" + set.shown("ink").levels()
	lamps := func(layer, number int) string {
		return fmt.Sprintf("%d;2;%s",
			layer, set.paletteOf(number).levels())
	}
	return escapes.ReplaceAllStringFunc(text, func(code string) string {
		bare := strings.TrimPrefix(code, "\x1b[")
		parameters := strings.Split(strings.TrimSuffix(bare, "m"), ";")
		said := []string{}
		for position := 0; position < len(parameters); position++ {
			part := parameters[position]
			number, err := strconv.Atoi(part)
			wide := position+2 < len(parameters) &&
				parameters[position+1] == "5"
			true24 := position+4 < len(parameters) &&
				parameters[position+1] == "2"
			settable := number == 38 || number == 48 || number == 58
			switch {
			case part == "" || part == "0":
				said = append(said, "0", ink)
			case err != nil:
				said = append(said, part)
			case number == 39:
				said = append(said, ink)
			case number >= 30 && number <= 37:
				said = append(said, lamps(38, number-30))
			case number >= 90 && number <= 97:
				said = append(said, lamps(38, number-90+8))
			case number >= 40 && number <= 47:
				said = append(said, lamps(48, number-40))
			case number >= 100 && number <= 107:
				said = append(said, lamps(48, number-100+8))
			case (number == 38 || number == 48) && wide:
				index, _ := strconv.Atoi(parameters[position+2])
				said = append(said, lamps(number, index))
				position += 2
			case settable && true24:
				said = append(said,
					parameters[position:position+5]...)
				position += 4
			default:
				said = append(said, part)
			}
		}
		return "\x1b[" + strings.Join(said, ";") + "m"
	})
}

// groundCode is the escape that draws behind text in a colour.
func groundCode(value colour) string {
	return "\x1b[48;2;" + value.levels() + "m"
}

// band draws one row of a page on its ground: the margin and the text,
// cut at textWidth, then filled to rowWidth with the line's own band if
// it has one, else the ground. the columns between the two widths are
// the gap before the gradient.
func band(text, margin string, ground colour, fill string, textWidth,
	rowWidth int) string {
	behind := groundCode(ground)
	drawn := ansi.Truncate(onGround(margin+text, behind), textWidth, "")
	if fill == "" {
		fill = behind
	}
	rest := max(textWidth-ansi.StringWidth(drawn), 0)
	gap := max(rowWidth-textWidth, 0)
	return behind + drawn + fill + strings.Repeat(" ", rest) +
		behind + strings.Repeat(" ", gap) + "\x1b[0m"
}

// sampleRows draws a page's lines as rows, marking in the margin each
// line that shows the stop being steered. The page's own ground and ink
// are in every line, so steering either marks none: a mark in every
// margin says nothing, and the steering line names them.
func sampleRows(lines []line, steered []string, everywhere []string,
	tint string, ground colour, textWidth, rowWidth int) []string {
	rows := make([]string, len(lines))
	for position, entry := range lines {
		margin := "  "
		marked := false
		for _, name := range entry.shows {
			marked = marked || slices.Contains(steered, name) &&
				!slices.Contains(everywhere, name)
		}
		bare := escapes.ReplaceAllString(entry.text, "")
		if marked && strings.TrimSpace(bare) != "" {
			margin = tint + "▸ "
		}
		rows[position] = band(entry.text, margin, ground, entry.fill,
			textWidth, rowWidth)
	}
	return rows
}
