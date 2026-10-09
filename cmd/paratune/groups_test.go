package main

import (
	"math"
	"os"
	"strings"
	"testing"
)

// g groups the colour steered with the one steered before; a move of one
// then moves the other by the same change in oklch, and g again takes it
// out of the group.
func TestGroupMovesTogether(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	steer(t, screen, markdownPage, "md-link")
	screen.press("g")
	if screen.grouping != "md-link" {
		t.Fatalf("the first g marked %q", screen.grouping)
	}
	steer(t, screen, markdownPage, "md-link-text")
	screen.press("g")
	if group := screen.set.groupOf("md-link-text"); len(group) != 2 {
		t.Fatalf("g made the group %v, and said %q", group, screen.note)
	}
	link, text := polarOf(screen.set.shown("md-link")), polarOf(screen.set.shown("md-link-text"))
	screen.Update(key("up", 0))
	movedLink, movedText := polarOf(screen.set.shown("md-link")), polarOf(screen.set.shown("md-link-text"))
	if math.Abs((movedLink.L-link.L)-(movedText.L-text.L)) > 0.01 || movedLink.L <= link.L {
		t.Errorf("a move of link text moved it %.3f in lightness and link %.3f",
			movedText.L-text.L, movedLink.L-link.L)
	}
	screen.press("g")
	if screen.set.groupOf("md-link-text") != nil {
		t.Error("g again left md-link-text in its group")
	}
}

// a move of every colour at once, like r or +, is not carried to a group,
// or the group would be moved twice.
func TestWholeMovesAreNotCarried(t *testing.T) {
	dir, home := fixture(t)
	alone := opened(t, dir, "test-dusk", home)
	grouped := opened(t, dir, "test-dusk", home)
	grouped.set.groups = [][]string{{"md-link", "md-link-text"}}
	for _, screen := range []*model{alone, grouped} {
		steer(t, screen, markdownPage, "md-link")
		screen.Update(key("+", '+'))
	}
	if alone.set.shown("md-link-text") != grouped.set.shown("md-link-text") {
		t.Error("+ moved a grouped colour twice")
	}
}

// groups are kept beside the source when it is saved and come back when
// it is opened; with none, there is no file.
func TestGroupsKept(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.set.groups = [][]string{{"md-heading", "md-small-heading"}}
	if note := screen.set.save(); !strings.HasPrefix(note, "saved") {
		t.Fatalf("save said %q", note)
	}
	again, err := load(dir, "test-dusk")
	if err != nil {
		t.Fatal(err)
	}
	if len(again.groups) != 1 || strings.Join(again.groups[0], " ") != "md-heading md-small-heading" {
		t.Errorf("the groups came back as %v", again.groups)
	}
	again.groups = nil
	again.save()
	if _, err := os.Stat(groupsPath(dir, "test-dusk")); err == nil {
		t.Error("a colourway with no groups kept a file of them")
	}
}

// g never groups the page's own ground or ink, and escape forgets a mark.
func TestGroupRefusesThePage(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	steer(t, screen, markdownPage, "md-heading")
	screen.press("g")
	steer(t, screen, markdownPage, screen.set.groundName(markdownPage))
	screen.press("g")
	if screen.set.groupOf("md-heading") != nil || !strings.Contains(screen.note, "not grouped") {
		t.Errorf("the page's ground was grouped: %v, %q", screen.set.groups, screen.note)
	}
	steer(t, screen, markdownPage, "md-heading")
	screen.press("g")
	if command := screen.press("esc"); command != nil || screen.grouping != "" {
		t.Error("escape did not forget the mark, or left paratune")
	}
}
