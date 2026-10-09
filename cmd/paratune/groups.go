package main

import (
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
)

// groupOf is the roles grouped with a role, itself among them, or nothing
// when it is in no group.
func (set *tuning) groupOf(name string) []string {
	for _, group := range set.groups {
		if slices.Contains(group, name) {
			return group
		}
	}
	return nil
}

// group is g, in two presses: the first on a colour in no group marks it,
// the second on another joins the two, and their groups become one; g on
// a grouped colour takes it out, and escape forgets a mark.
//
// The page's own ground and ink are never grouped: a move of them is a
// move of the whole page, and grouping one with a text colour is a slip,
// not a wish. It says what it did.
func (screen *model) group() string {
	set, name := screen.set, screen.stops[screen.steering].name
	everywhere := []string{set.groundName(screen.page),
		set.inkName(screen.page)}
	switch {
	case slices.Contains(everywhere, name):
		screen.grouping = ""
		return name + " is the page's own, and is not grouped"
	case screen.grouping == "" && set.groupOf(name) != nil:
		set.ungroup(name)
		return name + " out of its group"
	case screen.grouping == "" || screen.grouping == name:
		screen.grouping = name
		return "grouping " + name +
			": steer another, g to join, escape to stop"
	}
	joined := []string{screen.grouping, name}
	for _, other := range []string{screen.grouping, name} {
		for _, member := range set.groupOf(other) {
			if !slices.Contains(joined, member) {
				joined = append(joined, member)
			}
		}
		set.ungroup(other)
	}
	set.groups = append(set.groups, joined)
	set.dirty = true
	screen.grouping = ""
	return "grouped " + strings.Join(joined, ", ")
}

// ungroup takes a role out of its group, and drops a group left with one.
func (set *tuning) ungroup(name string) {
	kept := [][]string{}
	for _, group := range set.groups {
		group = slices.DeleteFunc(slices.Clone(group),
			func(member string) bool {
				return member == name
			})
		if len(group) > 1 {
			kept = append(kept, group)
		}
	}
	set.groups = kept
	set.dirty = true
}

// carry moves the rest of a role's group by the change the role was just
// given, in oklch: the same shift in lightness and chroma, the same turn
// of hue, each fitted into what the screen can make.
func (set *tuning) carry(name string, before colour) {
	after := set.shown(name)
	if after.hex() == before.hex() {
		return
	}
	was, is := polarOf(before), polarOf(after)
	for _, member := range set.groupOf(name) {
		if member == name {
			continue
		}
		polar := polarOf(set.shown(member))
		polar.L = math.Min(math.Max(polar.L+is.L-was.L, 0), 1)
		polar.C = math.Max(polar.C+is.C-was.C, 0)
		polar.H = math.Mod(polar.H+is.H-was.H+360, 360)
		set.roles.put(member, fittedPolar(polar))
	}
}

// wholeKeys move every colour of the colourway, or each on its own terms,
// so a group is not carried by them: it would be moved twice.
var wholeKeys = []string{"r", "+", "_", "-", "=", ">", "<"}

// carried is a move of the colour steered, carried to its group: true for
// a key or a click that moved one colour, false for one that moved many.
func carried(message tea.Msg) bool {
	switch message := message.(type) {
	case tea.KeyPressMsg:
		return !slices.Contains(wholeKeys, message.String())
	case tea.MouseClickMsg, tea.MouseWheelMsg:
		return true
	}
	return false
}

// groupsPath is where a colourway's groups are kept: beside its source,
// a group a line, since css has no way to say colours move together.
func groupsPath(dir, name string) string {
	return filepath.Join(dir, "sources", name+".groups")
}

// readGroups is a colourway's groups, or none when it has no file of them.
func readGroups(dir, name string) [][]string {
	raw, err := os.ReadFile(groupsPath(dir, name))
	if err != nil {
		return nil
	}
	groups := [][]string{}
	for _, line := range strings.Split(string(raw), "\n") {
		if members := strings.Fields(line); len(members) > 1 {
			groups = append(groups, members)
		}
	}
	return groups
}

// writeGroups keeps a colourway's groups beside its source, and takes the
// file away when there are none.
func writeGroups(dir, name string, groups [][]string) error {
	path := groupsPath(dir, name)
	if len(groups) == 0 {
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			return err
		}
		return nil
	}
	lines := []string{}
	for _, group := range groups {
		lines = append(lines, strings.Join(group, " "))
	}
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o644)
}
