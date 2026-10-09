package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/janearc/libtheme-css/spaces/ok"
)

// a code puts the terminal's own ground back when it resets, or names
// 49, with no ground set after it; a colour's own numbers are not codes.
func TestResets(t *testing.T) {
	for parameters, want := range map[string]bool{
		"": true, "0": true, "49": true, "1;0": true, "0;1": true,
		"38;5;0": false, "48;5;0": false, "48;2;0;0;0": false,
		"38;2;0;49;0": false, "0;48;5;17": false, "49;48;2;1;2;3": false,
		"48;5;17;0": true, "1": false, "22;39": false,
	} {
		if got := resets(parameters); got != want {
			t.Errorf("%q resets %v, want %v", parameters, got, want)
		}
	}
}

// onGround puts the ground back after every reset and nowhere else.
func TestOnGround(t *testing.T) {
	behind := "\x1b[48;2;1;2;3m"
	got := onGround("a\x1b[0mb\x1b[38;5;3mc\x1b[49md\x1b[m", behind)
	want := "a\x1b[0m" + behind + "b\x1b[38;5;3mc\x1b[49m" + behind + "d\x1b[m" + behind
	if got != want {
		t.Errorf("got %q", got)
	}
}

// a band is exactly the row's width, whatever it holds: short, long,
// wide characters, codes.
func TestBandWidth(t *testing.T) {
	ground := colour{128, 138, 99}
	for _, text := range []string{"", "short", strings.Repeat("long ", 40),
		"▸ ✻ ● ⎿ ─ ❯ ⏵⏵", "\x1b[38;5;252mcoloured\x1b[0m and not",
		"\x1b[1mbold " + strings.Repeat("x", 100)} {
		drawn := band(text, "  ", ground, "", 82, 84)
		if width := ansi.StringWidth(drawn); width != 84 {
			t.Errorf("%q is %d wide", text, width)
		}
	}
}

// the gradient has a row for every row of the pane, four cells each.
func TestGradientRows(t *testing.T) {
	for _, height := range []int{1, 10, 34, 61} {
		rows := gradientRows(height, sweep(ok.Mix))
		if len(rows) != height {
			t.Fatalf("%d rows for %d", len(rows), height)
		}
		for _, row := range rows {
			if width := ansi.StringWidth(row); width != 6 {
				t.Errorf("a gradient row is %d wide", width)
			}
		}
	}
}

// the frame is the pane exactly, on every page and with the menu open,
// at her pane's size and others, with the keys open too: rows to the height,
// each row to the width, so the gradient sits in the last four columns.
func TestViewFillsThePane(t *testing.T) {
	dir, home := fixture(t)
	for _, size := range [][2]int{{88, 34}, {120, 50}, {60, 20}, {92, 27}} {
		for _, keys := range []bool{false, true} {
			viewFills(t, dir, home, size, keys)
		}
	}
}

// ? opens the keys in place of the page and ? or escape closes them,
// escape without leaving; other keys do nothing while they are open.
func TestKeysPanel(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	before := screen.set.shown("ink")
	for _, step := range []struct {
		key  string
		open bool
	}{{"?", true}, {"up", true}, {"esc", false}, {"?", true}, {"?", false}} {
		if command := screen.press(step.key); command != nil {
			t.Fatalf("%s gave a command: paratune would have left", step.key)
		}
		if screen.keys != step.open {
			t.Errorf("after %s the keys are open %v", step.key, screen.keys)
		}
		if step.open && !strings.Contains(escapes.ReplaceAllString(screen.View().Content, ""),
			"shift-tab") {
			t.Errorf("after %s the keys are not shown", step.key)
		}
	}
	if screen.set.shown("ink") != before {
		t.Error("an arrow moved a colour while the keys were open")
	}
}

// viewFills checks one size, with the keys open or not.
func viewFills(t *testing.T, dir, home string, size [2]int, keys bool) {
	t.Helper()
	screen := opened(t, dir, "test-dusk", home)
	screen.width, screen.height, screen.keys = size[0], size[1], keys
	views := []string{}
	for which := range pageNames {
		screen.page = page(which)
		screen.steering = screen.first(screen.page)
		views = append(views, screen.View().Content)
	}
	screen.openMenu()
	views = append(views, screen.View().Content)
	for _, content := range views {
		rows := strings.Split(content, "\n")
		if len(rows) != size[1] {
			t.Fatalf("%dx%d keys %v: %d rows", size[0], size[1], keys, len(rows))
		}
		for number, row := range rows {
			if width := ansi.StringWidth(row); width != size[0] {
				t.Fatalf("%dx%d keys %v: row %d is %d wide", size[0], size[1], keys,
					number, width)
			}
		}
	}
}

// marks: a line is marked when it shows the role steered, and never for
// the page's own ink or ground, which are in every line.
func TestMarks(t *testing.T) {
	lines := []line{{text: "one", shows: []string{"ink", "ground", "yellow"}},
		{text: "two", shows: []string{"ink", "ground"}},
		{text: "", shows: []string{"yellow"}}}
	marked := func(steered string) []bool {
		rows := sampleRows(lines, []string{steered}, []string{"ground", "ink"}, "",
			colour{}, 40, 42)
		flags := make([]bool, len(rows))
		for position, row := range rows {
			flags[position] = strings.Contains(row, "▸")
		}
		return flags
	}
	for steered, want := range map[string][]bool{
		"yellow": {true, false, false}, "ink": {false, false, false},
		"ground": {false, false, false}, "red": {false, false, false},
	} {
		got := marked(steered)
		for position := range want {
			if got[position] != want[position] {
				t.Errorf("steering %s marks %v, want %v", steered, got, want)
				break
			}
		}
	}
}

// a pane too small to draw in says so rather than drawing wrong.
func TestTooSmall(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.width, screen.height = 20, 5
	if !strings.Contains(screen.View().Content, "bigger pane") {
		t.Error("a small pane was drawn into")
	}
}

// every colour named by number, and the terminal's own ink, is drawn as
// the colourway's own: the page does not lean on the pane's palette.
func TestExplicit(t *testing.T) {
	dir, home := fixture(t)
	set := opened(t, dir, "test-dusk", home).set
	ink := "38;2;" + set.shown("ink").levels()
	lamps := func(layer int, name string) string {
		return fmt.Sprintf("%d;2;%s", layer, set.shown(name).levels())
	}
	for text, want := range map[string]string{
		"\x1b[33mx":            "\x1b[" + lamps(38, "yellow") + "mx",
		"\x1b[1;36mx":          "\x1b[1;" + lamps(38, "cyan") + "mx",
		"\x1b[39mx":            "\x1b[" + ink + "mx",
		"\x1b[0mx":             "\x1b[0;" + ink + "mx",
		"\x1b[mx":              "\x1b[0;" + ink + "mx",
		"\x1b[38;5;252mx":      "\x1b[" + lamps(38, "claude-text") + "mx",
		"\x1b[48;5;0mx":        "\x1b[" + lamps(48, "black") + "mx",
		"\x1b[94mx":            "\x1b[" + lamps(38, "bright-blue") + "mx",
		"\x1b[38;2;1;2;3mx":    "\x1b[38;2;1;2;3mx",
		"\x1b[7mx\x1b[27m":     "\x1b[7mx\x1b[27m",
		"\x1b[22;33mx\x1b[49m": "\x1b[22;" + lamps(38, "yellow") + "mx\x1b[49m",
	} {
		if got := set.explicit(text); got != want {
			t.Errorf("%q became %q, want %q", text, got, want)
		}
	}
}

// a page whose pane has been reset by something else still draws in the
// colourway: no default colours reach the screen from a sample.
func TestPagesNeedNoPalette(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	for which := range pageNames {
		screen.page = page(which)
		screen.steering = screen.first(screen.page)
		for _, row := range strings.Split(screen.View().Content, "\n") {
			for _, code := range escapes.FindAllString(row, -1) {
				parameters := strings.TrimSuffix(strings.TrimPrefix(code, "\x1b["), "m")
				for _, part := range strings.Split(parameters, ";") {
					if part == "39" || (len(part) == 2 && (part[0] == '3' || part[0] == '4') &&
						part[1] >= '0' && part[1] <= '7') {
						if !strings.Contains(parameters, ";2;") {
							t.Errorf("page %d draws %q from the pane's palette", which, code)
						}
					}
				}
			}
		}
	}
}

// the theme strip draws one colour to a cell, the colour a click there
// takes, so a colour seen in it is the colour clicked.
func TestThemeStripIsWhatIsClicked(t *testing.T) {
	dir, home := fixture(t)
	set := opened(t, dir, "test-dusk", home).set
	ramp := themeRamp(set)
	rows := gradientRows(30, ramp)
	for row := range rows {
		cell := strings.SplitN(rows[row], "  ", 2)[0]
		if want := lampCode(48, ramp.At(float64(2*row)/float64(59))); cell != want {
			t.Fatalf("row %d of the theme strip is drawn %q, a click takes %q", row, cell, want)
		}
	}
}
