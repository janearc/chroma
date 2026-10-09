package main

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
)

// openMenu puts charm's huh in the page's place: the colourway, then the
// page, so the menu is how to get somewhere and the page is where the
// colours are tuned, in sight. It starts where paratune is.
func (screen *model) openMenu() tea.Cmd {
	screen.choice = choice{Name: screen.set.source.Name, Page: screen.page}
	return screen.buildMenu()
}

// buildMenu is huh's form for the choice as it stands, its cursors on
// the colourway and the page chosen.
func (screen *model) buildMenu() tea.Cmd {
	set := screen.set
	names := []huh.Option[string]{}
	width := max(screen.width-gradientWidth-12, 20)
	for _, name := range colourwayNames(set.dir) {
		label := name
		if held, found := screen.opened[name]; found && held.dirty ||
			name == set.source.Name && set.dirty {
			label += " *"
		}
		swatched := screen.swatched(label, name, width)
		names = append(names, huh.NewOption(swatched, name))
	}
	pages := make([]huh.Option[page], len(pageNames))
	for position, name := range pageNames {
		pages[position] = huh.NewOption(name, page(position))
	}
	ground := set.shown(set.groundName(screen.page))
	ink := controlInk(set.shown(set.inkName(screen.page)), ground)
	screen.menu = huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().Title("colourway").
			Value(&screen.choice.Name).
			Options(names...).Height(min(len(names)+1, 8)),
		huh.NewSelect[page]().Title("page").Value(&screen.choice.Page).
			Options(pages...),
	)).WithTheme(menuTheme(ink, ground)).WithShowHelp(true).
		WithWidth(max(screen.width-gradientWidth-4, 20))
	return screen.menu.Init()
}

// swatchCount is how many colours the menu shows beside a colourway.
const swatchCount = 10

// swatched is a colourway's name in the menu with its colours beside it,
// right-aligned to width: its ground, its ink, then its colours in hue
// order as the strip has them, so what it looks like is seen, not
// guessed from the name. A colourway that will not load shows its name.
func (screen *model) swatched(label, name string, width int) string {
	set, found := screen.opened[name]
	if name == screen.set.source.Name {
		set, found = screen.set, true
	}
	if !found {
		loaded, err := load(screen.set.dir, name)
		if err != nil {
			return label
		}
		set = loaded
	}
	colours := []colour{set.shown("ground"), set.shown("ink")}
	seen := map[string]bool{colours[0].hex(): true, colours[1].hex(): true}
	for _, stop := range themeRamp(set).Stops {
		if len(colours) == swatchCount {
			break
		}
		value := fitted(stop.Swatch)
		if !seen[value.hex()] {
			seen[value.hex()] = true
			colours = append(colours, value)
		}
	}
	cells := ""
	for _, value := range colours {
		cells += inkCode(value) + "██"
	}
	gap := max(width-len([]rune(label))-2*swatchCount, 1)
	return label + strings.Repeat(" ", gap) + cells + "\x1b[39m"
}

// menuClick is a left click on the open menu, which huh does not take:
// a colourway clicked is chosen, its cursor put on it, and a page
// clicked is gone to, with the colourway chosen.
func (screen *model) menuClick(y int) tea.Cmd {
	lines := strings.Split(screen.menu.View(), "\n")
	index := y - 1
	if index < 0 || index >= len(lines) {
		return nil
	}
	label := strings.TrimSpace(escapes.ReplaceAllString(lines[index], ""))
	label = strings.TrimSpace(strings.TrimPrefix(label, "┃"))
	label = strings.TrimSpace(strings.TrimPrefix(label, ">"))
	label = strings.TrimSpace(strings.TrimRight(label, "█ "))
	label = strings.TrimSpace(strings.TrimSuffix(label, "*"))
	for _, name := range colourwayNames(screen.set.dir) {
		if label == name {
			screen.choice.Name = name
			return screen.buildMenu()
		}
	}
	for position, name := range pageNames {
		if label == name {
			screen.choice.Page = page(position)
			screen.menu = nil
			screen.goChosen()
		}
	}
	return nil
}

// goChosen opens the colourway chosen and turns to the page chosen,
// steering its ink.
func (screen *model) goChosen() {
	screen.open(screen.choice.Name)
	if screen.set.source.Name == screen.choice.Name {
		screen.page = screen.choice.Page
		screen.steering = screen.first(screen.page)
	}
}

// steerMenu hands a message to the open menu: a key, or the room it has,
// which is the page's and not the pane's. ] closes it unchanged, as [
// opened it, and so does escape.
//
// When both are chosen paratune opens the colourway, unless the one
// showing has unsaved moves, and turns to the page, steering its ink.
func (screen *model) steerMenu(message tea.Msg) tea.Cmd {
	if key, ok := message.(tea.KeyPressMsg); ok &&
		(key.String() == "esc" || key.String() == "]") {
		screen.menu = nil
		return nil
	}
	updated, command := screen.menu.Update(message)
	if form, ok := updated.(*huh.Form); ok {
		screen.menu = form
	}
	switch screen.menu.State {
	case huh.StateCompleted:
		screen.menu = nil
		screen.goChosen()
		return nil
	case huh.StateAborted:
		screen.menu = nil
		return nil
	}
	return command
}

// menuTheme is huh's base theme in the page's own colours: the controls'
// ink on the page's ground, and the option under the cursor the other
// way round, so it reads as chosen on any ground being tuned.
func menuTheme(ink, ground colour) huh.Theme {
	inkColour := lipgloss.Color(ink.hex())
	groundColour := lipgloss.Color(ground.hex())
	inked := func(style lipgloss.Style) lipgloss.Style {
		return style.Foreground(inkColour)
	}
	return huh.ThemeFunc(func(isDark bool) *huh.Styles {
		styles := huh.ThemeBase(isDark)
		fields := []*huh.FieldStyles{&styles.Focused, &styles.Blurred}
		for _, field := range fields {
			field.Base = field.Base.BorderForeground(inkColour)
			field.Title = inked(field.Title).Bold(true)
			field.Description = inked(field.Description)
			field.SelectSelector = inked(field.SelectSelector)
			field.Option = inked(field.Option)
			field.UnselectedOption = inked(field.UnselectedOption)
			field.SelectedOption = field.SelectedOption.
				Foreground(groundColour).Background(inkColour)
			field.NextIndicator = inked(field.NextIndicator)
			field.PrevIndicator = inked(field.PrevIndicator)
		}
		styles.Blurred.Title = styles.Blurred.Title.Bold(false)
		styles.Help.ShortKey = inked(styles.Help.ShortKey)
		styles.Help.ShortDesc = inked(styles.Help.ShortDesc)
		styles.Help.ShortSeparator = inked(styles.Help.ShortSeparator)
		styles.Help.FullKey = inked(styles.Help.FullKey)
		styles.Help.FullDesc = inked(styles.Help.FullDesc)
		styles.Help.FullSeparator = inked(styles.Help.FullSeparator)
		return styles
	})
}
