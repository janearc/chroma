package main

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/janearc/libtheme-css/spaces/ok"
)

// a drawn pane, so clicks land on what is on screen.
func drawnPane(t *testing.T) *model {
	t.Helper()
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.View()
	return screen
}

// a click steers the colour under the pointer: git's hash in yellow, the
// branch's HEAD in cyan, its name in green, and the words after in the
// ink.
func TestClickALine(t *testing.T) {
	screen := drawnPane(t)
	row, plain := -1, ""
	for position, drawn := range screen.drawn {
		text := escapes.ReplaceAllString(drawn.text, "")
		if strings.HasPrefix(text, "4e1a9c2") {
			row, plain = position, text
		}
	}
	if row < 0 {
		t.Fatal("no line of git's log")
	}
	for word, want := range map[string]string{"4e1a9c2": "yellow", "HEAD": "cyan",
		"main": "green", "undo": "ink"} {
		column := len([]rune(plain[:strings.Index(plain, word)]))
		screen.Update(tea.MouseClickMsg{X: 2 + column + 1, Y: row, Button: tea.MouseLeft})
		if got := screen.stops[screen.steering].name; got != want {
			t.Errorf("a click on %q steered %s, not %s", word, got, want)
		}
	}
}

// the wheel moves the colour steered, and with shift its warmth.
func TestWheel(t *testing.T) {
	screen := drawnPane(t)
	before := screen.set.shown("ink")
	screen.Update(tea.MouseWheelMsg{Button: tea.MouseWheelUp})
	if up := screen.set.shown("ink"); up.luminance() <= before.luminance() ||
		!screen.set.dirty {
		t.Errorf("the wheel up took the ink from %s to %s", before.hex(), up.hex())
	}
	screen.Update(tea.MouseWheelMsg{Button: tea.MouseWheelDown})
	screen.Update(tea.MouseWheelMsg{Button: tea.MouseWheelDown})
	if down := screen.set.shown("ink"); down.luminance() >= before.luminance() {
		t.Errorf("the wheel down left the ink at %s", down.hex())
	}
	screen.Update(tea.MouseWheelMsg{Button: tea.MouseWheelUp})
	lighter := screen.set.shown("ink")
	screen.Update(tea.MouseWheelMsg{Button: tea.MouseWheelUp, Mod: tea.ModShift})
	if warm := screen.set.shown("ink"); warm.r <= lighter.r || warm.b >= lighter.b {
		t.Errorf("shift and the wheel did not warm %s to more than %s", warm.hex(), lighter.hex())
	}
}

// a click on the gradient gives the colour steered that hue: red at the
// top, on either side.
func TestClickTheGradient(t *testing.T) {
	screen := drawnPane(t)
	screen.Update(tea.MouseClickMsg{X: screen.width - 1, Y: 0, Button: tea.MouseLeft})
	if got := screen.set.shown("ink").hex(); got != "#ff0000" || !screen.set.dirty {
		t.Errorf("the top of the gradient gave %s", got)
	}
	row := screen.height / 3
	oklab, lamps := gradientColour(screen.height, row, false),
		gradientColour(screen.height, row, true)
	if oklab == lamps {
		t.Fatalf("row %d is the same on both sides", row)
	}
	screen.Update(tea.MouseClickMsg{X: screen.width - 4, Y: row, Button: tea.MouseLeft})
	if got := screen.set.shown("ink"); got != oklab {
		t.Errorf("the oklab side gave %s, not %s", got.hex(), oklab.hex())
	}
	screen.Update(tea.MouseClickMsg{X: screen.width - 1, Y: row, Button: tea.MouseLeft})
	if got := screen.set.shown("ink"); got != lamps {
		t.Errorf("the srgb side gave %s, not %s", got.hex(), lamps.hex())
	}
}

// with the menu open the mouse does nothing.
func TestMouseWaitsForTheMenu(t *testing.T) {
	screen := drawnPane(t)
	screen.openMenu()
	screen.Update(tea.MouseWheelMsg{Button: tea.MouseWheelUp})
	if screen.set.dirty {
		t.Error("the wheel moved a colour behind the menu")
	}
}

// the colourway's strip holds its colours, coloured ones by hue then
// greys, and a click on it picks the colour drawn there: the first at
// the top.
func TestThemeStrip(t *testing.T) {
	screen := drawnPane(t)
	ramp := themeRamp(screen.set)
	previous := -1.0
	for _, stop := range ramp.Stops {
		polar := ok.FromSwatch(stop.Swatch).Polar()
		if polar.C < ok.Eye {
			break
		}
		if polar.H < previous {
			t.Errorf("the strip's colours are not in hue order")
		}
		previous = polar.H
	}
	top := stripColour(ramp, screen.height, 0)
	screen.Update(tea.MouseClickMsg{X: screen.width - 6, Y: 0, Button: tea.MouseLeft})
	if got := screen.set.shown("ink"); got != top {
		t.Errorf("the strip's top gave %s, not %s", got.hex(), top.hex())
	}
}

// rowOf is the screen row a label is drawn on in the open menu.
func rowOf(t *testing.T, screen *model, label string) int {
	t.Helper()
	for position, text := range strings.Split(screen.menu.View(), "\n") {
		plain := strings.TrimSpace(escapes.ReplaceAllString(text, ""))
		plain = strings.TrimSpace(strings.TrimPrefix(plain, "┃"))
		plain = strings.TrimSpace(strings.TrimRight(strings.TrimPrefix(plain, ">"), "█ "))
		if strings.TrimSpace(plain) == label {
			return position + 1
		}
	}
	t.Fatalf("no %s in the menu", label)
	return 0
}

// in the menu, a click on a colourway chooses it and a click on a page
// goes there, with that colourway.
func TestClickTheMenu(t *testing.T) {
	screen := drawnPane(t)
	screen.Update(key("[", '['))
	screen.Update(tea.MouseClickMsg{X: 6, Y: rowOf(t, screen, "plain"), Button: tea.MouseLeft})
	if screen.menu == nil || screen.choice.Name != "plain" {
		t.Fatalf("clicking plain chose %q", screen.choice.Name)
	}
	if !strings.Contains(escapes.ReplaceAllString(screen.menu.View(), ""), "> plain") {
		t.Error("the cursor is not on plain")
	}
	screen.Update(tea.MouseClickMsg{X: 6, Y: rowOf(t, screen, "markdown"), Button: tea.MouseLeft})
	if screen.menu != nil || screen.set.source.Name != "plain" || screen.page != markdownPage {
		t.Errorf("went to %s page %d, menu open %v", screen.set.source.Name,
			screen.page, screen.menu != nil)
	}
}

// each colourway in the menu has its colours beside it, its ground first,
// right-aligned, the rows the same width.
func TestMenuSwatches(t *testing.T) {
	screen := drawnPane(t)
	screen.openMenu()
	widths := map[int]bool{}
	for _, name := range []string{"plain", "test-dusk"} {
		row := strings.Split(screen.menu.View(), "\n")[rowOf(t, screen, name)-1]
		if !strings.Contains(row, "\x1b[38;2;128;138;99m██") {
			t.Errorf("%s has no swatch of its ground", name)
		}
		widths[ansi.StringWidth(strings.TrimRight(escapes.ReplaceAllString(row, ""), " "))] = true
	}
	if len(widths) != 1 {
		t.Errorf("the swatches do not line up: %v", widths)
	}
}

// a cell's colours follow the codes before it: true colour, a reset to
// the ink and ground, reverse video, and a wide glyph's two cells.
func TestCellColours(t *testing.T) {
	ink, ground, red := colour{200, 200, 200}, colour{10, 10, 10}, colour{255, 0, 0}
	text := "a\x1b[38;2;255;0;0mb\x1b[7mc\x1b[0md漢e"
	for cell, want := range map[int][2]colour{0: {ink, ground}, 1: {red, ground},
		2: {ground, red}, 3: {ink, ground}, 5: {ink, ground}, 6: {ink, ground}} {
		fg, bg, _, found := cellColours(text, cell, ink, ground)
		if !found || fg != want[0] || bg != want[1] {
			t.Errorf("cell %d is %s on %s", cell, fg.hex(), bg.hex())
		}
	}
	if _, _, _, found := cellColours(text, 9, ink, ground); found {
		t.Error("a cell past the line's end was found")
	}
}

// where a click lands on no colour a role has, it cycles the line's own
// colours, the next each time, round to the first.
func TestSteerLineCycles(t *testing.T) {
	screen := drawnPane(t)
	clicked := line{shows: []string{"ink", "ground", "yellow", "cyan", "green"}}
	seen := []string{}
	for range 4 {
		screen.steerLine(clicked)
		seen = append(seen, screen.stops[screen.steering].name)
	}
	if strings.Join(seen, " ") != "yellow cyan green yellow" {
		t.Errorf("the cycle went %v", seen)
	}
}

// when two roles are drawn in the colour clicked, a click again on the
// same place steers the other, and a third goes back.
func TestClickSharedColour(t *testing.T) {
	screen := drawnPane(t)
	screen.set.roles.put("yellow", screen.set.shown("cyan"))
	screen.View()
	row, plain := -1, ""
	for position, drawn := range screen.drawn {
		text := escapes.ReplaceAllString(drawn.text, "")
		if strings.HasPrefix(text, "4e1a9c2") {
			row, plain = position, text
		}
	}
	column := len([]rune(plain[:strings.Index(plain, "HEAD")]))
	seen := []string{}
	for range 3 {
		screen.Update(tea.MouseClickMsg{X: 2 + column + 1, Y: row, Button: tea.MouseLeft})
		seen = append(seen, screen.stops[screen.steering].name)
	}
	if seen[0] == seen[1] || seen[0] != seen[2] {
		t.Errorf("clicks on a colour two roles share steered %v", seen)
	}
}

// a search match in vim is the page's ground cut out of a highlight, and
// a click on one of its letters steers the highlight, not the ground.
func TestClickASearchMatch(t *testing.T) {
	screen := drawnPane(t)
	screen.page, screen.steering = vimPage, screen.first(vimPage)
	screen.View()
	for y, each := range screen.drawn {
		plain := escapes.ReplaceAllString(each.text, "")
		at := strings.Index(plain, ".Swatch")
		if at < 0 {
			continue
		}
		column := len([]rune(plain[:at])) + 1
		_, behind, _, _ := cellColours(each.text, column, colour{}, colour{})
		screen.Update(tea.MouseClickMsg{X: 2 + column, Y: y, Button: tea.MouseLeft})
		here := screen.stops[screen.steering]
		if got := screen.set.shown(here.name); got.hex() != behind.hex() {
			t.Errorf("a click on the match steered %s %s, not the highlight %s",
				here.name, got.hex(), behind.hex())
		}
		return
	}
	t.Fatal("no search match drawn on the vim page")
}

// a click on bare ground, past a line's end or on an empty line, steers
// the page's ground.
func TestClickBareGround(t *testing.T) {
	screen := drawnPane(t)
	ground := screen.set.groundName(screen.page)
	for _, y := range []int{0, len(screen.drawn) - 3} {
		screen.steering = screen.stopOf("yellow")
		screen.Update(tea.MouseClickMsg{X: 40, Y: y, Button: tea.MouseLeft})
		if got := screen.stops[screen.steering].name; got != ground {
			t.Errorf("a click on bare ground at row %d steered %s", y, got)
		}
	}
}
