package main

import (
	"encoding/json"
	"slices"
	"strings"
	"testing"

	"github.com/janearc/libtheme-css/colourway"
	"github.com/janearc/libtheme-css/css"
	"github.com/janearc/libtheme-css/dialects/claude"
	"github.com/janearc/libtheme-css/dialects/glamour"
)

// every page has stops to tune, and none names the same role twice.
func TestEveryPageHasStops(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	for which := range pageNames {
		seen := map[string]bool{}
		for _, here := range screen.stops {
			if here.page != page(which) {
				continue
			}
			if seen[here.name] {
				t.Errorf("%s is twice on page %d", here.name, which)
			}
			seen[here.name] = true
		}
		if len(seen) == 0 {
			t.Errorf("page %d has nothing to tune", which)
		}
	}
}

// where the colourway gives vim its own colour, the vim page tunes that
// and the terminal's page the shared one.
func TestVimTunesItsOwn(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.set.roles.put("nvim-ink", colour{200, 180, 160})
	screen.stops = stopsOf(screen.set.roles)
	steer(t, screen, vimPage, "nvim-ink")
	steer(t, screen, shellPage, "ink")
	for _, here := range screen.stops {
		if here.page == vimPage && here.name == "ink" {
			t.Error("the vim page still tunes the shared ink")
		}
	}
}

// vim's grounds are grounds, its own as well as shared ones, read
// against vim's ink, with no terminal field to pair with.
func TestVimGroundsAreGrounds(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.set.roles.put("nvim-ground", colour{10, 20, 30})
	screen.stops = stopsOf(screen.set.roles)
	for _, name := range []string{"nvim-ground", "cursor-line", "surface-1", "panel"} {
		here := steer(t, screen, vimPage, name)
		if !here.ground || here.partner != "" {
			t.Errorf("%s: ground %v, partner %q", name, here.ground, here.partner)
		}
	}
	if here := steer(t, screen, vimPage, "dim"); here.ground {
		t.Error("dim is taken for a ground")
	}
}

// every field pairs both ways: a ground names what is on it, and that
// names the ground.
func TestPairsBothWays(t *testing.T) {
	for ground, ink := range pairs {
		if partnerOf(ground) != ink || partnerOf(ink) != ground {
			t.Errorf("%s and %s do not pair both ways", ground, ink)
		}
		if !grounds[ground] {
			t.Errorf("%s pairs as a ground and is not one", ground)
		}
	}
}

// a role the source does not set is shown at its program's fallback:
// vim's own at the shared one, claude code's at xterm's, the sixteen at
// xterm's, a general role as libtheme derives it, another ground at the
// page's ground.
func TestShownFallsBack(t *testing.T) {
	dir, home := fixture(t)
	set := opened(t, dir, "plain", home).set
	panel, _ := colourway.Resolve(set.roles.sheet()).Of("panel")
	for name, want := range map[string]colour{
		"nvim-ink":    set.shown("ink"),
		"claude-mode": set.shown("blue"),
		"panel":       colourOf(panel),
		"cursor-ink":  set.shown("ground"),
		"cursor":      set.shown("ground"),
		"dim":         set.shown("bright-black"),
		"heading":     set.shown("magenta"),
	} {
		if got := set.shown(name); got != want {
			t.Errorf("%s shows %s, want %s", name, got.hex(), want.hex())
		}
	}
	full := opened(t, dir, "test-dusk", home).set
	if full.shown("nvim-heading") != full.shown("heading") ||
		full.shown("heading") == full.shown("ink") {
		t.Error("vim's own heading does not fall back to the shared heading")
	}
	set.roles = rolesOf(css.New())
	if got := set.shown("bright-cyan"); got != xterm(14) {
		t.Errorf("an unset bright-cyan shows %s", got.hex())
	}
	if got := set.shown("claude-mode"); got != xterm(180) {
		t.Errorf("claude-mode with nothing to follow shows %s", got.hex())
	}
}

// the markdown page shows each unset role at the colour the glamour
// dialect writes for it from the resolved roles, as render does, for a
// colourway with an editor's roles and for one with only a terminal's:
// what is tuned is what will be rendered.
func TestMarkdownFallbacksAgreeWithGlamour(t *testing.T) {
	dir, home := fixture(t)
	for _, name := range []string{"test-dusk", "plain"} {
		set := opened(t, dir, name, home).set
		style, err := glamour.Of(colourway.Resolve(set.roles.sheet()).Roles)
		if err != nil {
			t.Fatal(err)
		}
		parsed := map[string]any{}
		if err := json.Unmarshal(style, &parsed); err != nil {
			t.Fatal(err)
		}
		for _, entry := range glamour.Roles {
			if strings.HasSuffix(entry.Name, "-ground") {
				continue
			}
			written, found := lookupKey(parsed, entry.Keys[0])
			if !found {
				t.Errorf("%s: glamour wrote no %s", name, entry.Keys[0])
				continue
			}
			if got := set.shown(entry.Name).hex(); got != written {
				t.Errorf("%s: %s shows %s, glamour writes %s", name, entry.Name,
					got, written)
			}
		}
	}
}

// lookupKey is the string at a dotted key in parsed json.
func lookupKey(style map[string]any, key string) (string, bool) {
	var node any = style
	for _, part := range strings.Split(key, ".") {
		object, found := node.(map[string]any)
		if !found {
			return "", false
		}
		node = object[part]
	}
	text, found := node.(string)
	return text, found
}

// the contrast shown is against a field's partner, else the page's own
// ground or, for a ground, the page's ink.
func TestAgainst(t *testing.T) {
	dir, home := fixture(t)
	set := opened(t, dir, "test-dusk", home).set
	set.roles.put("md-heading-ground", colour{1, 2, 3})
	for _, case_ := range []struct {
		here stop
		want string
	}{
		{stop{name: "md-heading", page: markdownPage, partner: "md-heading-ground"}, "md-heading-ground"},
		{stop{name: "ground", page: shellPage, ground: true}, "ink"},
		{stop{name: "red", page: shellPage}, "ground"},
		{stop{name: "md-quote", page: markdownPage}, "ground"},
		{stop{name: "ground", page: claudePage, ground: true}, "claude-text"},
	} {
		if got, want := set.against(case_.here), set.shown(case_.want); got != want {
			t.Errorf("%s on page %d is read against %s, want %s's %s",
				case_.here.name, case_.here.page, got.hex(), case_.want, want.hex())
		}
	}
}

// a role the colourway does not set says where it comes from, following
// each step that is not set either: vim's heading from the shared one,
// which comes from magenta; a ground mixed from the ground and the ink.
func TestOrigin(t *testing.T) {
	dir, home := fixture(t)
	plain := opened(t, dir, "plain", home).set
	full := opened(t, dir, "test-dusk", home).set
	for _, each := range []struct {
		set        *tuning
		name, want string
	}{
		{plain, "heading", "from magenta"},
		{plain, "nvim-heading", "from heading, from magenta"},
		{plain, "md-heading", "from heading, from magenta"},
		{plain, "panel", "from ground and ink"},
		{plain, "claude-mode", "from info, from blue"},
		{plain, "claude-added", "from ground and ok"},
		{full, "nvim-heading", "from heading"},
	} {
		if got := each.set.origin(each.name); got != each.want {
			t.Errorf("%s in %s: %q, want %q", each.name, each.set.source.Name, got, each.want)
		}
	}
}

// a derived role follows what it comes from as soon as that moves: the
// derivation is worked out again, not kept from before.
func TestDerivedFollows(t *testing.T) {
	dir, home := fixture(t)
	set := opened(t, dir, "plain", home).set
	set.shown("heading")
	moved := colour{200, 30, 160}
	set.roles.put("magenta", moved)
	if got := set.shown("heading"); got.hex() != moved.hex() {
		t.Errorf("heading shows %s after magenta moved to %s", got.hex(), moved.hex())
	}
}

// claude code's colours show as the terminal theme will carry them: what
// libtheme resolves, so the claude page tunes what is rendered.
func TestClaudeAgreesWithRender(t *testing.T) {
	dir, home := fixture(t)
	set := opened(t, dir, "plain", home).set
	resolved := colourway.Resolve(set.roles.sheet())
	for _, slot := range claude.Slots {
		want, found := resolved.Of(slot.Role)
		if !found {
			continue
		}
		if got := set.shown(slot.Role).hex(); got != colourOf(want).hex() {
			t.Errorf("%s shows %s, the theme will have %s", slot.Role, got, colourOf(want).hex())
		}
	}
}

// a code line shows only the roles its visible characters are drawn in:
// the block's text colour, which chroma uses for nothing but spaces in
// this sample, marks no line, and the names, which are what the code is
// mostly drawn in, mark several.
func TestCodeShowsWhatIsDrawn(t *testing.T) {
	dir, home := fixture(t)
	set := opened(t, dir, "test-dusk", home).set
	block, names := 0, 0
	for _, each := range markdownLines(set, 80) {
		if slices.Contains(each.shows, "md-block-text") {
			block++
		}
		if slices.Contains(each.shows, "md-name") {
			names++
		}
	}
	if block != 0 || names < 2 {
		t.Errorf("md-block-text marks %d lines and md-name %d", block, names)
	}
}

// the markdown page is drawn from the roles as libtheme resolves them,
// the same the float shows and r reads: a derived heading is drawn in the
// colour it is derived as.
func TestMarkdownDrawnResolved(t *testing.T) {
	dir, home := fixture(t)
	set := opened(t, dir, "plain", home).set
	heading := "38;2;" + set.shown("md-heading").levels()
	if !strings.Contains(strings.Join(markdownSample(set, 80), "\n"), heading) {
		t.Errorf("no heading drawn in %s, the colour it is derived as", set.shown("md-heading").hex())
	}
}

// code drawn on the markdown page follows a change to its colour within
// one paratune: glamour keeps its first code colours unless told not to.
func TestCodeFollowsAChange(t *testing.T) {
	dir, home := fixture(t)
	set := opened(t, dir, "test-dusk", home).set
	for _, value := range []colour{{75, 30, 75}, {16, 224, 16}} {
		set.roles.put("md-name", value)
		if drawn := strings.Join(markdownSample(set, 80), "\n"); !strings.Contains(drawn,
			"38;2;"+value.levels()+"m") {
			t.Errorf("md-name moved to %s and the code is not drawn in it", value.hex())
		}
	}
}
