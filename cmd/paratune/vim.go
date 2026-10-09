package main

import (
	"fmt"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/janearc/libtheme-css/dialects/nvim"
)

// vimTable is the colours vim's table holds, by the names its groups
// use, as the colourway shows them now, and the role each is tuned as:
// the shared one, or vim's own where the colourway gives it one.
func vimTable(set *tuning) (map[string]colour, map[string]string) {
	values, roleOf := map[string]colour{}, map[string]string{}
	for _, entry := range nvim.Roles {
		role := entry.Role
		if _, own := set.roles.get("nvim-" + role); own {
			role = "nvim-" + role
		}
		values[entry.Name], roleOf[entry.Name] = set.shown(role), role
	}
	return values, roleOf
}

// cell is one character of the vim page and the names, in the scheme's
// table, of the colours it is drawn in.
type cell struct {
	text         rune
	fg, bg       string
	bold, italic bool
}

// cells are text as cells in one pair of colours.
func cells(text, fg, bg string) []cell {
	drawn := []cell{}
	for _, character := range text {
		drawn = append(drawn, cell{text: character, fg: fg, bg: bg})
	}
	return drawn
}

// padded is a row of cells cut or filled to width with a ground.
func padded(row []cell, width int, bg string) []cell {
	if len(row) > width {
		return row[:width]
	}
	for len(row) < width {
		row = append(row, cell{text: ' ', fg: "ink", bg: bg})
	}
	return row
}

// vimScene is what the vim page shows happening: where the cursor is, the
// word last searched for, the visual selection, and the diagnostics.
var vimScene = struct {
	cursorRow, cursorColumn int
	search                  string
	visualRow               int
	visualFrom, visualTo    int
	diagnostics             map[int][2]string
}{
	cursorRow: 12, cursorColumn: 10,
	search:    "Swatch",
	visualRow: 17, visualFrom: 11, visualTo: 25,
	diagnostics: map[int][2]string{
		2:  {"info", "■ organize imports"},
		9:  {"hint", "■ could be unexported"},
		20: {"warn", "■ func unused is unused"},
		21: {"err", "■ declared and not used"},
	},
}

// codeRow is one line of the file as vim draws it: the number, the code
// in her scheme's syntax colours, the search hits, the selection and the
// cursor laid over it, and any diagnostic after it.
func codeRow(row int, tokens [][]chromaToken, width int) []cell {
	ground := "bg"
	number := cells(fmt.Sprintf("%3d ", row+1), "gutter", ground)
	if row == vimScene.cursorRow {
		ground = "cursorln"
		number = cells(fmt.Sprintf("%3d ", row+1), "heading", "bg")
		for position := range number {
			number[position].bold = true
		}
	}
	code := []cell{}
	for _, token := range tokens[row] {
		piece := cells(token.text, token.group, ground)
		for position := range piece {
			piece[position].italic = token.italic
		}
		code = append(code, piece...)
	}
	for _, position := range hits(code, []rune(vimScene.search)) {
		for offset := range len([]rune(vimScene.search)) {
			hit := &code[position+offset]
			hit.fg, hit.bg = "bg", "string"
		}
	}
	if row == vimScene.visualRow {
		to := min(vimScene.visualTo, len(code))
		for position := vimScene.visualFrom; position < to; position++ {
			code[position].bg = "visual"
		}
	}
	if row == vimScene.cursorRow && vimScene.cursorColumn < len(code) {
		under := &code[vimScene.cursorColumn]
		under.fg, under.bg = "bg", "ink"
	}
	if diagnostic, ok := vimScene.diagnostics[row]; ok {
		code = append(code,
			cells("  "+diagnostic[1], diagnostic[0], ground)...)
	}
	return padded(append(number, code...), width, ground)
}

// hits are where a word starts in a row of cells, counted in cells.
func hits(row []cell, word []rune) []int {
	found := []int{}
	for start := 0; len(word) > 0 && start+len(word) <= len(row); start++ {
		matched := true
		for offset, character := range word {
			matched = matched && row[start+offset].text == character
		}
		if matched {
			found = append(found, start)
			start += len(word) - 1
		}
	}
	return found
}

// chromaToken is a token of the sample with the vim colour it is drawn
// in; comments are in italic, as her scheme sets them.
type chromaToken struct {
	text   string
	group  string
	italic bool
}

// vimTokens are the sample's lines as tokens in her scheme's colours.
func vimTokens() [][]chromaToken {
	lines := tokenLines(sampleCode)
	drawn := make([][]chromaToken, len(lines))
	for row, tokens := range lines {
		for _, token := range tokens {
			drawn[row] = append(drawn[row], chromaToken{
				text:   token.Value,
				group:  vimGroup(token.Type),
				italic: token.Type.InCategory(chroma.Comment)})
		}
	}
	return drawn
}

// netrw is the directory listing in the split on the right.
var netrw = []string{"\" chroma/cmd/paratune", "../", "./", "claude.go",
	"controls.go", "gradient.go", "main.go", "markdown.go", "pages.go",
	"shell.go", "vim.go"}

// listingRow is one row of the listing: the banner in dim, directories in
// link, files in ink.
func listingRow(row, width int) []cell {
	if row >= len(netrw) {
		return padded(cells("~", "gutter", "bg"), width, "bg")
	}
	entry := netrw[row]
	group := "ink"
	switch {
	case strings.HasPrefix(entry, "\""):
		group = "dim"
	case strings.HasSuffix(entry, "/"):
		group = "link"
	}
	return padded(cells(entry, group, "bg"), width, "bg")
}

// statusRow is the two windows' status lines: the focused one in ink on
// the panel, with the diagnostics counted in their own colours, the other
// in the gutter's colour.
func statusRow(left, right int) []cell {
	row := cells(" gradient.go [+]  ", "ink", "panel")
	for _, count := range [][2]string{{"err", "E1 "}, {"warn", "W1 "},
		{"info", "I1 "}, {"hint", "H1"}} {
		row = append(row, cells(count[1], count[0], "panel")...)
	}
	tail := cells("go  13,11  Top ", "ink", "panel")
	row = padded(row, max(left-len(tail), 0), "panel")
	row = append(row, tail...)
	if right > 0 {
		row = append(padded(row, left, "panel"),
			cells(" ", "border", "panel")...)
		title := cells(" chroma/cmd/paratune/", "gutter", "panel")
		row = append(row, padded(title, right, "panel")...)
	}
	return row
}

// vimLines are the vim page: a go file in her scheme as it is being
// tuned, split beside a directory listing when there is room, the status
// lines, and the search on the command line.
func vimLines(set *tuning, count, width int) []line {
	values, roleOf := vimTable(set)
	tokens := vimTokens()
	right := 0
	if width >= 72 {
		right = 22
	}
	left := width
	if right > 0 {
		left = width - right - 1
	}
	rows := [][]cell{}
	for row := 0; row < max(count-2, 0); row++ {
		drawn := padded(cells("~", "gutter", "bg"), left, "bg")
		if row < len(tokens) {
			drawn = codeRow(row, tokens, left)
		}
		if right > 0 {
			drawn = append(drawn,
				cell{text: '│', fg: "border", bg: "bg"})
			drawn = append(drawn, listingRow(row, right)...)
		}
		rows = append(rows, drawn)
	}
	rows = append(rows, statusRow(left, right))
	search := cells("/"+vimScene.search, "ink", "bg")
	rows = append(rows, padded(search, width, "bg"))
	lines := make([]line, 0, len(rows))
	for _, row := range rows[:min(len(rows), count)] {
		lines = append(lines, line{text: drawCells(row, values),
			shows: vimShows(row, roleOf)})
	}
	return lines
}

// vimShows are the roles a row of cells is drawn in.
func vimShows(row []cell, roleOf map[string]string) []string {
	seen := map[string]bool{}
	shows := []string{}
	for _, here := range row {
		for _, name := range []string{here.fg, here.bg} {
			if here.text == ' ' && name == here.fg {
				continue
			}
			if !seen[name] {
				seen[name] = true
				shows = append(shows, roleOf[name])
			}
		}
	}
	return shows
}

// drawCells is a row of cells as text, each run of one style drawn in
// its colours from the table as the colourway shows them now.
func drawCells(row []cell, values map[string]colour) string {
	var out strings.Builder
	last := cell{fg: "-"}
	bolds := map[bool]string{true: "\x1b[1m", false: "\x1b[22m"}
	italics := map[bool]string{true: "\x1b[3m", false: "\x1b[23m"}
	for _, here := range row {
		if here.fg != last.fg || here.bg != last.bg ||
			here.bold != last.bold || here.italic != last.italic {
			out.WriteString(inkCode(values[here.fg]))
			out.WriteString(groundCode(values[here.bg]))
			out.WriteString(bolds[here.bold])
			out.WriteString(italics[here.italic])
		}
		out.WriteRune(here.text)
		last = here
	}
	return out.String()
}
