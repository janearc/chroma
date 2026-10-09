package main

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/janearc/libtheme-css/dialects/claude"
)

// slotRole is the colourway's role for one of claude code's palette
// numbers, as the claude dialect names it.
func slotRole(number int) string {
	for _, slot := range claude.Slots {
		if slot.Number == number {
			return slot.Role
		}
	}
	return ""
}

// claudeExchange is one exchange as claude code draws it, in the palette
// codes it uses, so a slot moved is seen where it really appears: her
// words in their band, a reply, an edit with its diff, the prompt between
// its rules, the mode line, the spinner and the bar.
//
// each line has the slots it is drawn in, and the band it runs in to the
// edge, if any; a {rule} is drawn across what is left of the width.
var claudeExchange = []struct {
	text  string
	slots []int
	fill  string
}{
	{"\x1b[38;5;252m\x1b[48;5;0m❯ \x1b[38;5;188m" +
		"can the controls stop short of the gradient?",
		[]int{252, 188, 0}, "\x1b[48;5;0m"},
	{"", nil, ""},
	{"\x1b[38;5;188m●\x1b[39m They run two columns into the gap at " +
		"your width. I'll",
		[]int{188}, ""},
	{"  single-space the two key lines in \x1b[38;5;153mcontrols.go" +
		"\x1b[39m so they stop short.",
		[]int{153}, ""},
	{"", nil, ""},
	{"\x1b[38;5;188m●\x1b[39m \x1b[1mUpdate\x1b[22m" +
		"(paratune/controls.go)", []int{188}, ""},
	{"\x1b[38;5;252m  ⎿  Updated paratune/controls.go " +
		"with 1 addition and 1 removal",
		[]int{252}, ""},
	{"\x1b[38;5;242m       110    // the keys, in the controls' ink",
		[]int{}, ""},
	{"\x1b[38;5;167m\x1b[48;5;52m       111 -" +
		"\x1b[38;5;231m  \"  arrows: lighter, darker.  tab: next.\"",
		[]int{52, 231}, "\x1b[48;5;52m"},
	{"\x1b[38;5;77m\x1b[48;5;22m       111 +" +
		"\x1b[38;5;231m  \"  arrows: lighter, darker. tab: next.\"",
		[]int{22, 231}, "\x1b[48;5;22m"},
	{"\x1b[38;5;242m       112    // c v foreground, d f ground",
		[]int{}, ""},
	{"", nil, ""},
	{"\x1b[38;5;252m✻ Cogitated for 54s · done 4:34 PM", []int{252}, ""},
	{"", nil, ""},
	{"\x1b[38;5;175m{rule} \x1b[38;5;16m\x1b[48;5;175m chroma " +
		"\x1b[49m\x1b[38;5;175m ─",
		[]int{175}, ""},
	{"\x1b[38;5;252m❯ \x1b[48;5;219m \x1b[49m", []int{252, 219}, ""},
	{"\x1b[38;5;175m{rule}", []int{175}, ""},
	{"\x1b[38;5;180m  ⏵⏵ auto mode on\x1b[38;5;252m (shift+tab to cycle)",
		[]int{180, 252}, ""},
	{"\x1b[38;5;219m● Thinking… \x1b[38;5;252m(54s · thinking)",
		[]int{219, 252}, ""},
	{"\x1b[38;5;251m\x1b[48;5;17m the bar \x1b[0m " +
		"\x1b[38;5;252mOpus 5.5 xhigh   c 70%",
		[]int{252}, ""},
}

// claudeLines are the exchange as the claude page's lines, width wide,
// each with the stops it shows.
func claudeLines(width int) []line {
	lines := make([]line, len(claudeExchange))
	for position, entry := range claudeExchange {
		text := entry.text
		rest := strings.Replace(text, "{rule}", "", 1)
		if rest != text {
			filled := max(width-ansi.StringWidth(rest), 0)
			text = strings.Replace(text, "{rule}",
				strings.Repeat("─", filled), 1)
		}
		lines[position] = line{text: text, fill: entry.fill}
		for _, number := range entry.slots {
			lines[position].shows = append(lines[position].shows,
				slotRole(number))
		}
	}
	return lines
}
