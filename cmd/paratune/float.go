package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/janearc/libtheme-css/spaces/ok"
)

// quietAfter is how long without a key or a click before the float shows
// all of itself. While keys are being pressed, the sample is what is
// being looked at, so the float is its field's row alone.
const quietAfter = time.Minute

// quietMsg is the time to look again at whether it has gone quiet.
type quietMsg struct{}

// box is where the float was last drawn, for clicks: its first column,
// its width, how many rows, and the row ? help is on.
type box struct {
	left, width, rows, help, field int
}

// float is the box over the top right of the page: the colourway, the
// page, how few of its colours the colourway sets when that is few, the
// field steered in its own colour with its ratio and where an unset one
// comes from, a note when there is one, and ? help.
//
// While keys are being pressed it is the field's row and the note alone.
func (screen *model) float(tint string, ground colour) []string {
	set, here := screen.set, screen.stops[screen.steering]
	value := set.shown(here.name)
	steered := field(here.label, value, ground, tint) +
		fmt.Sprintf("  %.1f:1", contrast(value, set.against(here)))
	note := []string{}
	if screen.note != "" {
		note = append(note, tint+screen.note)
	}
	if screen.busy {
		return append([]string{steered}, note...)
	}
	name := set.source.Name
	if set.dirty {
		name += "*"
	}
	shown := pageNames[screen.page]
	if theirs := displays[screen.display].name; theirs != "" {
		shown += ", " + theirs
	}
	rows := []string{tint + name, tint + shown}
	if share := screen.unsetNote(); share != "" {
		rows = append(rows, tint+share)
	}
	rows = append(rows, steered)
	if _, own := set.roles.get(here.name); !own {
		rows = append(rows, tint+set.origin(here.name))
	}
	if screen.grouping != "" {
		rows = append(rows,
			tint+"grouping "+screen.grouping+": g joins")
	}
	if group := set.groupOf(here.name); group != nil {
		others := slices.DeleteFunc(slices.Clone(group),
			func(member string) bool {
				return member == here.name
			})
		rows = append(rows, tint+"grouped: "+strings.Join(others, ", "))
	}
	if copied := screen.held.copied(); copied != "" {
		rows = append(rows, tint+copied)
	}
	return append(append(rows, note...), tint+"? help")
}

// overlay draws the float over the top rows of the page, ending one cell
// from the row's right edge, on a ground of its own, and says where it
// went. Each row keeps its width: the page left of the float as it was,
// the float, then the page's ground for the margin.
func overlay(rows, float []string, rowWidth int, ground, page colour) box {
	width := 0
	for _, text := range float {
		width = max(width, ansi.StringWidth(text)+2)
	}
	width = min(width, rowWidth/2)
	placed := box{left: rowWidth - width - 1, width: width,
		rows: min(len(float), len(rows)), help: -1, field: -1}
	for index := range placed.rows {
		text := ansi.Truncate(float[index], width-2, "…")
		padding := strings.Repeat(" ", width-2-ansi.StringWidth(text))
		left := ansi.Truncate(rows[index], placed.left, "")
		short := strings.Repeat(" ", placed.left-ansi.StringWidth(left))
		rows[index] = left + groundCode(page) + short +
			groundCode(ground) + " " + text + groundCode(ground) +
			padding + " " + groundCode(page) + " "
		if strings.HasSuffix(float[index], "? help") {
			placed.help = index
		}
		if strings.Contains(float[index], ":1") && placed.field < 0 {
			placed.field = index
		}
	}
	return placed
}

// stirred is a key or a click: the float shrinks to its field's row, and
// a look at whether it has gone quiet is scheduled, if none is waiting.
func (screen *model) stirred(message tea.Msg) tea.Cmd {
	switch message.(type) {
	case tea.KeyPressMsg, tea.MouseClickMsg, tea.MouseWheelMsg:
	default:
		return nil
	}
	screen.lastStirred, screen.busy = time.Now(), true
	if screen.quietWaiting {
		return nil
	}
	screen.quietWaiting = true
	return tea.Tick(quietAfter,
		func(time.Time) tea.Msg { return quietMsg{} })
}

// quiet is a look at whether nothing has been pressed for quietAfter: if
// so the float shows all of itself, and if not, another look when it
// would be.
func (screen *model) quiet() tea.Cmd {
	screen.quietWaiting = false
	since := time.Since(screen.lastStirred)
	if since >= quietAfter {
		screen.busy = false
		return nil
	}
	screen.quietWaiting = true
	return tea.Tick(quietAfter-since,
		func(time.Time) tea.Msg { return quietMsg{} })
}

// floatGround is the float's own ground: the page's, mixed a tenth of the
// way in oklab away from the marks' ink, toward black under a light ink
// and white under a dark one, so the float reads as a thing on the page
// and its marks read on it at least as well as on the page.
func floatGround(ground, ink colour) colour {
	away := colour{255, 255, 255}
	if ink.luminance() > ground.luminance() {
		away = colour{}
	}
	return colourOf(ok.Mix.Mix(swatchOf(ground), swatchOf(away), 0.1))
}
