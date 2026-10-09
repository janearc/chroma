package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/janearc/libtheme-css/dialects/claude"
	"github.com/janearc/libtheme-css/dialects/ghostty"
	"github.com/janearc/libtheme-css/dialects/glamour"
	"github.com/janearc/libtheme-css/dialects/nvim"
)

// stop is one role the arrows can move, on the page it is tuned on.
// label is how it is named on screen.
//
// A ground is read against the ink on it, anything else against the
// ground under it; partner names that other role when it is not the
// page's own, and makes the two a field for the registers.
type stop struct {
	name    string
	label   string
	page    page
	ground  bool
	partner string
}

// pairs are the fields: each ground and what is drawn on it.
var pairs = map[string]string{
	"cursor": "cursor-ink", "surface-1": "selection-ink",
	"black": "claude-words", "claude-added": "claude-diff-code",
	"md-heading-ground":       "md-heading",
	"md-small-heading-ground": "md-small-heading",
	"md-code-ground":          "md-code",
}

// grounds are the roles that are drawn behind something.
var grounds = map[string]bool{"ground": true, "cursor": true,
	"surface-1": true, "cursor-line": true, "panel": true, "black": true,
	"claude-added": true, "claude-removed": true, "md-heading-ground": true,
	"md-small-heading-ground": true, "md-code-ground": true}

// partnerOf is the role a stop is read against, when it is not the
// page's own ground or ink.
func partnerOf(name string) string {
	if fg, found := pairs[name]; found {
		return fg
	}
	for bg, fg := range pairs {
		if fg == name {
			return bg
		}
	}
	if name == "claude-removed" {
		return "claude-diff-code"
	}
	return ""
}

// markdownStops are the glamour roles the markdown page shows, in the
// order its sample shows them; the rest of glamour's roles keep their
// fallbacks.
var markdownStops = []string{"md-body", "md-heading", "md-heading-ground",
	"md-small-heading", "md-small-heading-ground", "md-emphasis",
	"md-strong", "md-link", "md-link-text", "md-bullet", "md-quote",
	"md-code", "md-code-ground", "md-block-text", "md-keyword", "md-type",
	"md-function", "md-string", "md-number", "md-comment", "md-operator",
	"md-name"}

// stopsOf is every role, page by page, in the order tab walks them, as
// the dialects name them: the terminal's settings and sixteen, claude
// code's colours, glamour's, and neovim's table, with an nvim- role in
// place of the shared one where the colourway gives vim its own.
//
// vim's grounds are read against vim's ink and its text against vim's
// ground: its fields are not the terminal's.
func stopsOf(all *roles) []stop {
	stops := []stop{}
	add := func(name, label string, on page) {
		stops = append(stops, stop{name: name, label: label, page: on,
			ground: grounds[name], partner: partnerOf(name)})
	}
	for _, setting := range ghostty.Keys {
		add(setting.Role, setting.Role, shellPage)
	}
	for _, name := range ghostty.Ansi {
		add(name, name, shellPage)
	}
	add("ground", "ground", claudePage)
	for _, slot := range claude.Slots {
		label := fmt.Sprintf("%s (%d)", slot.Role, slot.Number)
		if slot.Role == "black" {
			label = "black, the band (0)"
		}
		add(slot.Role, label, claudePage)
	}
	add("ground", "ground", markdownPage)
	for _, name := range markdownStops {
		add(name, name, markdownPage)
	}
	for _, entry := range nvim.Roles {
		name := entry.Role
		if _, own := all.get("nvim-" + name); own {
			name = "nvim-" + name
		}
		stops = append(stops, stop{name: name, label: name,
			page: vimPage, ground: grounds[entry.Role]})
	}
	return stops
}

// shown is a role's colour as the page draws it: the source's, the one
// libtheme derives for a general role, or the colour a program falls
// back to while the source does not set it.
func (set *tuning) shown(name string) colour {
	if value, found := set.roles.get(name); found {
		return value
	}
	if value, found := set.resolution().Of(name); found && derived(name) {
		return colourOf(value)
	}
	for number, ansi := range ghostty.Ansi {
		if ansi == name {
			return xterm(number)
		}
	}
	switch {
	case strings.HasPrefix(name, "nvim-"):
		return set.shown(strings.TrimPrefix(name, "nvim-"))
	case strings.HasPrefix(name, "md-"):
		return set.markdownFallback(name)
	case strings.HasPrefix(name, "claude-"):
		return set.claudeFallback(name)
	case name == "cursor-ink" || grounds[name] && name != "ground":
		return set.shown("ground")
	}
	if name == "ink" || name == "ground" {
		return xterm(map[string]int{"ink": 252, "ground": 235}[name])
	}
	return set.shown("ink")
}

// claudeFallback is the colour claude code draws a role in when the
// terminal does not set its number: xterm's.
func (set *tuning) claudeFallback(name string) colour {
	for _, slot := range claude.Slots {
		if slot.Role == name {
			return xterm(slot.Number)
		}
	}
	return set.shown("ink")
}

// markdownFallback is a glamour role followed through its fallbacks, as
// the glamour dialect follows them over the resolved roles, to the ink; a
// ground with nothing to fall back to shows the page's ground.
func (set *tuning) markdownFallback(name string) colour {
	return set.shown(set.markdownFollows(name))
}

// markdownFollows is the role a glamour role the colourway does not set
// takes its colour from: its first fallback that has one, or another
// markdown role, which follows its own; else the ground or the ink.
func (set *tuning) markdownFollows(name string) string {
	for _, entry := range glamour.Roles {
		if entry.Name != name {
			continue
		}
		for _, fallback := range entry.Fallbacks {
			if set.has(fallback) ||
				strings.HasPrefix(fallback, "md-") {
				return fallback
			}
		}
	}
	if strings.HasSuffix(name, "-ground") {
		return "ground"
	}
	return "ink"
}

// follows is the role, or the palette, a role the colourway does not set
// takes its colour from, as shown follows it; empty when nothing gives it
// one.
func (set *tuning) follows(name string) string {
	if general(name) {
		return set.resolution().From[name]
	}
	for _, ansi := range ghostty.Ansi {
		if ansi == name {
			return "xterm"
		}
	}
	switch {
	case strings.HasPrefix(name, "nvim-"):
		return strings.TrimPrefix(name, "nvim-")
	case strings.HasPrefix(name, "md-"):
		return set.markdownFollows(name)
	case strings.HasPrefix(name, "claude-"):
		if from := set.resolution().From[name]; from != "" {
			return from
		}
		return "xterm"
	case name == "cursor-ink" || grounds[name] && name != "ground":
		return "ground"
	}
	return "ink"
}

// endsChain is whether where a role comes from stops at a source: one
// the colourway sets, or one that is not a role at all, like xterm's
// palette or a ground mixed toward something.
func (set *tuning) endsChain(source string) bool {
	_, own := set.roles.get(source)
	notRole := source == "xterm" || strings.HasPrefix(source, "ground and ")
	return own || notRole
}

// origin is where a role the colourway does not set comes from, as the
// steering line says it: what it follows, and, while that is not set
// either, what that follows, a few steps at most.
func (set *tuning) origin(name string) string {
	steps := []string{}
	for current := name; len(steps) < 3; {
		source := set.follows(current)
		if source == "" {
			break
		}
		steps = append(steps, "from "+source)
		if set.endsChain(source) || source == current {
			break
		}
		current = source
	}
	if len(steps) == 0 {
		return "unset"
	}
	return strings.Join(steps, ", ")
}

// against is what a stop is read against for the contrast shown: its
// partner in the field, or its page's ink or ground.
func (set *tuning) against(here stop) colour {
	switch {
	case here.partner != "":
		return set.shown(here.partner)
	case here.ground:
		return set.shown(set.inkName(here.page))
	}
	return set.shown(set.groundName(here.page))
}

// keep adds the colourway's set roles to ~/paratune-kept.txt as one
// line, and says what happened, for the status line.
func keep(home string, set *tuning) string {
	kept := fmt.Sprintf("%s %s", time.Now().Format("2006-01-02 15:04"),
		set.source.Name)
	for _, name := range set.roles.order {
		value, _ := set.roles.get(name)
		kept += fmt.Sprintf(" %s=%s", name, value.hex())
	}
	path := filepath.Join(home, "paratune-kept.txt")
	file, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return "could not keep: " + err.Error()
	}
	defer file.Close()
	if _, err := file.WriteString(kept + "\n"); err != nil {
		return "could not keep: " + err.Error()
	}
	return "kept"
}
