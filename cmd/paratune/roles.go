package main

import (
	"maps"
	"slices"

	"github.com/janearc/libtheme-css/css"
	"github.com/janearc/libtheme-css/primitives/swatch"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// roles are the colourway being tuned: the roles its source sets, by
// name, in the source's order. A role the source does not set is shown
// as libtheme derives it, or at the colour a program falls back to, and
// becomes the source's own the moment it is moved.
type roles struct {
	values map[string]colour
	order  []string
	// how many times a role has been put, so a derivation knows it is
	// stale
	changes int
}

// rolesOf is a source's roles.
func rolesOf(sheet *css.Sheet) *roles {
	all := &roles{values: map[string]colour{}}
	for _, rule := range sheet.Rules() {
		all.put(rule.Name, colourOf(rule.Swatch))
	}
	return all
}

// get is a role's colour, if the source sets it.
func (all *roles) get(name string) (colour, bool) {
	value, found := all.values[name]
	return value, found
}

// put sets a role, making it the source's own if it was not.
func (all *roles) put(name string, value colour) {
	if _, found := all.values[name]; !found {
		all.order = append(all.order, name)
	}
	all.values[name] = value
	all.changes++
}

// copied is the roles as they are now, kept apart from what comes after:
// a snapshot of every colour, which undo can go back to.
func (all *roles) copied() *roles {
	return &roles{values: maps.Clone(all.values),
		order:   slices.Clone(all.order),
		changes: all.changes}
}

// same is whether two sets of roles hold the same colours.
func (all *roles) same(other *roles) bool {
	return maps.Equal(all.values, other.values)
}

// sheet is the roles the source sets, in its order, as a sheet.
func (all *roles) sheet() *css.Sheet {
	sheet := css.New()
	for _, name := range all.order {
		sheet.Set(name, swatchOf(all.values[name]))
	}
	return sheet
}

// colourOf is a swatch as paratune's colour, each lamp nought to 255.
func colourOf(value swatch.Swatch) colour {
	lamps, _ := srgb.FromSwatch(value)
	return colour{lamps.R * 255, lamps.G * 255, lamps.B * 255}
}

// swatchOf is paratune's colour as a swatch, at the byte the file will
// hold, so what is saved is what was shown.
func swatchOf(value colour) swatch.Swatch {
	lamps, err := srgb.FromHex(value.hex())
	if err != nil {
		return srgb.RGB{}.Swatch()
	}
	return lamps.Swatch()
}
