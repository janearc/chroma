package main

import (
	"fmt"
	"strings"
)

// shellSession is a session as one meets it in a shell: the sixteen,
// each named in its own colour, then a prompt, git's log, a listing and a
// test run, each line with the terminal's colours it shows. It is drawn
// in the terminal's own ink and the sixteen tools reach for.
var shellSession = []struct {
	text  string
	shows []string
}{
	{"\x1b[1m~/chroma\x1b[22m % colours", nil},
	colourRow("normal", 30, ""),
	colourRow("bright", 90, "bright-"),
	{"\x1b[1m~/chroma\x1b[22m % git log --oneline -4", nil},
	{"\x1b[33m4e1a9c2 (\x1b[1;36mHEAD -> \x1b[32mmain\x1b[22;33m)\x1b[39m" +
		" paratune: undo, a step at a time",
		[]string{"yellow", "cyan", "green"}},
	{"\x1b[33m7d3b0f5\x1b[39m colourways: a picture of each, by family",
		[]string{"yellow"}},
	{"\x1b[33mc82e614\x1b[39m libreadme: contrast is measured, not chosen",
		[]string{"yellow"}},
	{"\x1b[33m19af3d7\x1b[39m libtheme: colour, taken apart",
		[]string{"yellow"}},
	{"\x1b[1m~/chroma\x1b[22m % ls", nil},
	{"\x1b[34mcmd\x1b[39m  \x1b[34mcolourways\x1b[39m  go.mod  go.sum  " +
		"\x1b[34mlibreadme\x1b[39m  " +
		"\x1b[34mlibtheme\x1b[39m  README.md",
		[]string{"blue"}},
	{"\x1b[1m~/chroma\x1b[22m % go test ./libtheme/...", nil},
	{"ok      chroma/libtheme/spaces/ok   0.412s", nil},
	{"\x1b[31m--- FAIL: TestInkOnGround (0.03s)\x1b[39m", []string{"red"}},
	{"\x1b[1m~/chroma\x1b[22m % less doc/colourway.md", nil},
}

// colourNames are the sixteen's names, in the terminal's order.
var colourNames = []string{"black", "red", "green", "yellow", "blue", "magenta",
	"cyan", "white"}

// colourRow is a line of the sixteen, each named in its own colour: the
// normal eight from code 30, or the bright from 90, so every one of them
// has a place on the page to be seen and clicked.
func colourRow(label string, first int, prefix string) struct {
	text  string
	shows []string
} {
	text, shows := label+" ", []string{}
	for index, name := range colourNames {
		text += fmt.Sprintf(" \x1b[%dm%s\x1b[39m", first+index, name)
		shows = append(shows, prefix+name)
	}
	return struct {
		text  string
		shows []string
	}{text, shows}
}

// sampleText is the file less shows on the shell page when none is named:
// a few words about a colourway, long enough to fill a pane.
const sampleText = `# notes on a colourway

a colourway is one css sheet: a ground, an ink, the sixteen colours a
terminal hands its programs, and whatever else the programs ask for.
paratune shows it as each program would draw it, and lets each colour
be moved where it is seen.

## reading against the ground

contrast is measured as wcag measures it: the lighter luminance over
the darker, each with a twentieth added for flare. a reader's band sets
a floor and a ceiling. below the floor, text is too dim to read for
long; above the ceiling, it glares.

## the sixteen

eight colours and their bright twins, numbered as the terminal numbers
them. most programs reach for a handful: red for what failed, green for
what passed, blue for a directory, yellow for what changed.

## derived colours

what a colourway does not say, libtheme derives from what it does: a
heading from magenta, a link from blue, a panel from the ground mixed a
little toward the ink. paratune shows each derived colour, and where it
came from, and moving one makes it the colourway's own.`

// shellLines are the session, then the file as less shows it, line for
// line, to the foot of the page: a few words of it selected, and less's
// prompt in reverse, the ink as a ground, with the cursor after it.
func shellLines(set *tuning, text string, count, width int) []line {
	lines := []line{}
	add := func(text string, shows ...string) {
		lines = append(lines, line{text: text,
			shows: append([]string{"ink", "ground"}, shows...)})
	}
	for _, entry := range shellSession {
		add(entry.text, entry.shows...)
	}
	paged := lessLines(text, width)
	room := max(count-len(lines)-1, 0)
	shown := min(len(paged), room)
	for position, text := range paged[:shown] {
		if position == 2 && len(strings.Fields(text)) > 4 {
			add(selected(set, text, 4),
				"surface-1", "selection-ink")
			continue
		}
		add(text)
	}
	for len(lines) < count-1 {
		lines = append(lines, line{})
	}
	cursor := groundCode(set.shown("cursor")) +
		inkCode(set.shown("cursor-ink")) + " \x1b[0m"
	add(fmt.Sprintf("\x1b[7mdoc/colourway.md lines 1-%d %d%%\x1b[27m%s",
		shown, shown*100/max(len(paged), 1), cursor),
		"cursor", "cursor-ink")
	return lines[:min(len(lines), count)]
}

// selected is a line with its first few words drawn as a selection is,
// whole words, in the selection's ink on its ground.
func selected(set *tuning, text string, words int) string {
	end := 0
	for range words {
		rest := text[end:]
		start := end + len(rest) - len(strings.TrimLeft(rest, " "))
		next := strings.IndexByte(text[start:], ' ')
		if next < 0 {
			end = len(text)
			break
		}
		end = start + next
	}
	return groundCode(set.shown("surface-1")) +
		inkCode(set.shown("selection-ink")) +
		text[:end] + "\x1b[0m" + text[end:]
}

// lessLines are a file's lines as less draws them: as written, and a line
// longer than the pane carried on to the next.
func lessLines(text string, width int) []string {
	var out []string
	for _, written := range strings.Split(text, "\n") {
		characters := []rune(written)
		for len(characters) > width {
			out = append(out, string(characters[:width]))
			characters = characters[width:]
		}
		out = append(out, string(characters))
	}
	return out
}
