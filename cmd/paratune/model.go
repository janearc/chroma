package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
	"github.com/janearc/libtheme-css/colourway"
)

// model is paratune as bubbletea runs it: the colourway being tuned,
// which page and stop are showing, the registers, and huh's menu while
// it is open.
type model struct {
	home, text    string
	set           *tuning
	stops         []stop
	steering      int
	page          page
	held          registers
	note          string
	width, height int
	menu          *huh.Form
	choice        choice
	painter       func(*tuning)
	drawn         []line
	opened        map[string]*tuning
	keys          bool  // the keys panel is open in place of the page
	picking       bool  // the picker is open over the foot of the page
	walking       *walk // n's walk through the palette, while it goes on
	busy          bool  // a key or a click came less than quietAfter ago
	quietWaiting  bool  // a look at whether it has gone quiet is on its way
	lastStirred   time.Time
	placed        box // where the float was last drawn
	// the colour g marked, waiting for a second g on another
	grouping     string
	paintDue     bool // the pane's colours are behind the colourway's
	paintWaiting bool // a paint is already on its way
	// which of displays backslash has chosen: ours first
	display int
}

// paintEvery is the most often the pane is painted while a key is held:
// painting is a call to tmux, far slower than a frame, so a held key
// would otherwise queue presses faster than they can be painted.
const paintEvery = 100 * time.Millisecond

// paintMsg is the time to paint, if the colours have moved since.
type paintMsg struct{}

// repaint marks the pane's colours as behind the colourway's. The paint
// itself comes after, at most every paintEvery, so the frame keeps up
// with a held key and the pane follows a moment later.
func (screen *model) repaint() {
	screen.paintDue = true
}

// schedulePaint is the command that will paint, when a paint is due and
// none is on its way, and paratune was given a way to paint: main gives
// it tmux's, and tests none.
func (screen *model) schedulePaint() tea.Cmd {
	if !screen.paintDue || screen.paintWaiting || screen.painter == nil {
		return nil
	}
	screen.paintWaiting = true
	return tea.Tick(paintEvery,
		func(time.Time) tea.Msg { return paintMsg{} })
}

// painted is the paint arriving: the pane takes the colourway's colours as
// they are now.
func (screen *model) painted() {
	screen.paintWaiting = false
	if screen.paintDue && screen.painter != nil {
		screen.paintDue = false
		screen.painter(screen.set)
	}
}

// choice is the colourway and page chosen in the menu.
type choice struct {
	Name string
	Page page
}

// Init starts nothing; the first window size draws the first page.
func (screen *model) Init() tea.Cmd {
	return nil
}

// Update takes a message: a new size, a key for paratune, or anything
// for the menu while it is open.
func (screen *model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	set, steering := screen.set, screen.steering
	name := screen.stops[steering].name
	before, snapshot := set.shown(name), set.roles.copied()
	_, command := screen.update(message)
	if screen.set == set && screen.steering == steering &&
		carried(message) {
		set.carry(name, before)
	}
	if screen.set == set && !undoing(message) && !set.roles.same(snapshot) {
		set.history = append(set.history, snapshot)
	}
	return screen, tea.Batch(command, screen.schedulePaint(),
		screen.stirred(message))
}

// undoing is whether a message is u, which takes a step back rather than
// being one.
func undoing(message tea.Msg) bool {
	key, isKey := message.(tea.KeyPressMsg)
	return isKey && key.String() == "u"
}

// undo puts the open colourway back as it was before its last step. Each
// step kept a snapshot of every colour, so undo needs no record of what
// the step did, and it goes back as far as there are steps: a colour
// lifted into white comes back down with its hue.
func (screen *model) undo() {
	set := screen.set
	if len(set.history) == 0 {
		screen.note = "nothing to undo: this is how it opened"
		return
	}
	last := len(set.history) - 1
	set.roles, set.history = set.history[last], set.history[:last]
	set.dirty = !set.roles.same(set.saved)
	screen.repaint()
	screen.note = "undone"
	if last == 0 {
		screen.note = "undone: back to how it opened"
	}
}

// update is Update before any paint is scheduled.
func (screen *model) update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case paintMsg:
		screen.painted()
		return screen, nil
	case quietMsg:
		return screen, screen.quiet()
	case tea.WindowSizeMsg:
		screen.width, screen.height = message.Width, message.Height
		if screen.menu != nil {
			return screen, screen.steerMenu(tea.WindowSizeMsg{
				Width:  screen.width - gradientWidth - 2,
				Height: screen.height - 7})
		}
		return screen, nil
	case tea.MouseClickMsg:
		if screen.menu != nil {
			if message.Button == tea.MouseLeft {
				return screen, screen.menuClick(message.Y)
			}
			return screen, nil
		}
		return screen, screen.mouse(message)
	case tea.MouseWheelMsg:
		return screen, screen.mouse(message)
	case tea.KeyPressMsg:
		if message.String() == "ctrl+c" {
			return screen, tea.Quit
		}
		if screen.menu == nil {
			return screen, screen.press(message.String())
		}
	}
	if screen.menu != nil {
		return screen, screen.steerMenu(message)
	}
	return screen, nil
}

// lighter, darker, warmer and cooler are what the arrows do to a colour.
func lighter(value colour) colour {
	return colour{value.r*1.03 + 1, value.g*1.03 + 1, value.b*1.03 + 1}
}

// darker is a colour one step darker.
func darker(value colour) colour {
	return colour{value.r * 0.97, value.g * 0.97, value.b * 0.97}
}

// warmer is a colour one step warmer.
func warmer(value colour) colour {
	return colour{value.r * 1.02, value.g, value.b * 0.98}
}

// cooler is a colour one step cooler.
func cooler(value colour) colour {
	return colour{value.r * 0.98, value.g, value.b * 1.02}
}

// press carries out one key: the arrows move the stop being steered, tab
// walks the page's stops, n walks the palette, r makes the page readable.
//
// [ opens the menu, ? the keys, the letters copy and paste, u undoes a
// step, enter keeps, s saves, i installs, backslash steps between our
// display and charm's, escape leaves.
//
// While the keys are open only ? and escape do anything: they close them.
func (screen *model) press(key string) tea.Cmd {
	screen.note = ""
	if screen.keys {
		if key == "?" || key == "esc" {
			screen.keys = false
		}
		return nil
	}
	if screen.grouping != "" && key == "esc" {
		screen.grouping = ""
		screen.note = "not grouped"
		return nil
	}
	if screen.picking && (key == "o" || key == "esc") {
		screen.picking = false
		return nil
	}
	switch key {
	case "?":
		screen.keys = true
	case "o":
		screen.picking = true
	case "g":
		screen.note = screen.group()
	case "esc":
		return tea.Quit
	case "up":
		screen.nudge(lighter)
	case "down":
		screen.nudge(darker)
	case "right":
		screen.nudge(warmer)
	case "left":
		screen.nudge(cooler)
	case "tab":
		screen.step(1)
	case "shift+tab":
		screen.step(-1)
	case "[":
		return screen.openMenu()
	case "-":
		screen.night(1)
	case "=":
		screen.night(-1)
	case "r":
		screen.readable()
	case "n":
		screen.paletteStep(1)
	case "N", "shift+n":
		screen.paletteStep(-1)
	case ">":
		screen.textLightness(0.02)
	case "<":
		screen.textLightness(-0.02)
	case "+":
		screen.whole(1, 1.1)
	case "_":
		screen.whole(1, 0.9)
	case "c", "v", "d", "f", "y", "p":
		screen.note = screen.held.use(key,
			screen.stops[screen.steering], screen.set)
		screen.repaint()
	case "enter":
		screen.note = keep(screen.home, screen.set)
	case "u":
		screen.undo()
		return nil
	case "s":
		screen.note = screen.set.save()
	case "i":
		screen.note = screen.set.install(screen.home)
	case "\\":
		return screen.otherDisplay()
	}
	return nil
}

// nudge moves the stop being steered, makes it the source's own, and
// repaints the pane.
func (screen *model) nudge(change func(colour) colour) {
	name := screen.stops[screen.steering].name
	screen.set.roles.put(name, change(screen.set.shown(name)).clamp())
	screen.set.dirty = true
	screen.repaint()
}

// first is a page's first stop to steer: its ink, the colour most of it
// is drawn in, or its first stop if the ink is not among them.
func (screen *model) first(which page) int {
	ink, fallback := screen.set.inkName(which), -1
	for index, here := range screen.stops {
		if here.page != which {
			continue
		}
		if here.name == ink {
			return index
		}
		if fallback < 0 {
			fallback = index
		}
	}
	return max(fallback, 0)
}

// step walks to the next stop on the page showing, or the one before,
// passing over any the page does not draw: a colour nothing on the page
// is drawn in can be moved, but not seen moving.
func (screen *model) step(direction int) {
	drawn := screen.drawnRoles()
	count := len(screen.stops)
	for range screen.stops {
		screen.steering = (screen.steering + direction + count) % count
		here := screen.stops[screen.steering]
		if here.page == screen.page &&
			(len(drawn) == 0 || drawn[here.name]) {
			return
		}
	}
}

// drawnRoles is every role a line of the page last drawn shows, and the
// page's own ground and ink, which are under and in every line, so tab
// steers only what can be seen; empty before the page is first drawn.
func (screen *model) drawnRoles() map[string]bool {
	if len(screen.drawn) == 0 {
		return map[string]bool{}
	}
	drawn := map[string]bool{screen.set.groundName(screen.page): true,
		screen.set.inkName(screen.page): true}
	for _, each := range screen.drawn {
		for _, name := range each.shows {
			drawn[name] = true
		}
	}
	return drawn
}

// open turns to another colourway. The one left keeps its unsaved moves,
// still marked unsaved, and has them again when it is turned back to.
func (screen *model) open(name string) {
	if name == screen.set.source.Name {
		return
	}
	screen.walking = nil
	if screen.opened == nil {
		screen.opened = map[string]*tuning{}
	}
	screen.opened[screen.set.source.Name] = screen.set
	set, found := screen.opened[name]
	if !found {
		loaded, err := load(screen.set.dir, name)
		if err != nil {
			screen.note = "could not open: " + err.Error()
			return
		}
		set = loaded
	}
	screen.set, screen.stops = set, stopsOf(set.roles)
	screen.repaint()
}

// save writes the source and renders it again: each program's file and
// its inverted twin. It refuses a source holding something it would
// lose, and says which programs rendered and which refused, and why.
func (set *tuning) save() string {
	if set.lossy != "" {
		return "not saved: a save would change " + set.lossy
	}
	set.source.Roles = set.roles.sheet()
	path := filepath.Join(set.dir, "sources", set.source.Name+".css")
	if err := set.source.Write(path); err != nil {
		return "could not save: " + err.Error()
	}
	lost := writeGroups(set.dir, set.source.Name, set.groups)
	if lost != nil {
		return "saved; groups not kept: " + lost.Error()
	}
	set.dirty, set.saved = false, set.roles.copied()
	_, refused := rendered(set.source, set.dir)
	if refused != "" {
		return "saved; not rendered: " + refused
	}
	return "saved and rendered"
}

// rendered renders a source and says what it wrote, by path, and what
// it refused, by program and reason.
func rendered(source colourway.Source, dir string) ([]string, string) {
	paths, refused := []string{}, []string{}
	for _, each := range colourway.Render(source, dir) {
		if each.Err != nil {
			if each.Name == source.Name {
				refused = append(refused,
					each.Program+" ("+each.Err.Error()+")")
			}
			continue
		}
		paths = append(paths, each.Path)
	}
	return paths, strings.Join(refused, ", ")
}

// install renders the saved source afresh and copies exactly what that
// render wrote, the colourway and its twin, to where the programs read
// them, so nothing stale goes out. It says how many, and what refused.
func (set *tuning) install(home string) string {
	if set.dirty {
		return "unsaved: s saves first"
	}
	saved, err := colourway.Read(filepath.Join(set.dir, "sources",
		set.source.Name+".css"))
	if err != nil {
		return "could not install: " + err.Error()
	}
	paths, refused := rendered(saved, set.dir)
	places := make([]string, len(paths))
	for index, path := range paths {
		relative, err := filepath.Rel(set.dir, path)
		if err != nil {
			return "could not install: " + err.Error()
		}
		to, known := colourway.Place(filepath.Dir(relative), home)
		if !known {
			return "nothing installed: no place known for " +
				relative
		}
		places[index] = filepath.Join(to, filepath.Base(path))
	}
	done := colourway.Installed{}
	for index, path := range paths {
		if err := done.Copy(path, places[index], true); err != nil {
			return "could not install: " + err.Error()
		}
	}
	note := fmt.Sprintf("installed %d files where each program reads them",
		len(paths))
	if count := len(done.Replaced); count > 0 {
		note += fmt.Sprintf("; %d took the place of a different file",
			count)
	}
	if refused != "" {
		note += "; not rendered: " + refused
	}
	return note
}

// View draws the frame: the page, or the menu or the keys in its place,
// every row on the page's ground, the float over its top right while the
// page shows, and the gradient down the right beside them all.
func (screen *model) View() tea.View {
	view := tea.NewView("")
	view.AltScreen = true
	view.MouseMode = tea.MouseModeCellMotion
	if screen.width < 30 || screen.height < 10 {
		view.Content = "paratune needs a bigger pane"
		return view
	}
	textWidth, rowWidth := screen.width-gradientWidth, screen.width-6
	set := screen.set
	ground := set.shown(set.groundName(screen.page))
	ink := controlInk(set.shown(set.inkName(screen.page)), ground)
	tint := inkCode(ink)
	rows := screen.body(screen.height, textWidth, rowWidth, tint, ground)
	screen.placed = box{help: -1, field: -1}
	if screen.menu == nil && !screen.keys {
		own := floatGround(ground, ink)
		float := screen.float(tint, own)
		screen.placed = overlay(rows, float, rowWidth, own, ground)
		if screen.picking && len(rows) > pickerRows+4 {
			for index, text := range screen.picker(rowWidth, tint) {
				rows[len(rows)-pickerRows+index] =
					groundCode(ground) + text
			}
		}
	}
	gradient := gradientRows(len(rows), themeRamp(set))
	for row := range rows {
		rows[row] += gradient[row]
	}
	view.Content = strings.Join(rows, "\n")
	return view
}

// body is the page's sample, or the menu in its place, cut or filled to
// count rows.
func (screen *model) body(count, textWidth, rowWidth int, tint string,
	ground colour) []string {
	steered := screen.stops[screen.steering].name
	var lines []line
	switch {
	case screen.menu != nil:
		steered = ""
		lines = []line{{}}
		for _, text := range strings.Split(screen.menu.View(), "\n") {
			lines = append(lines, line{text: text})
		}
	case screen.keys:
		steered = ""
		lines = []line{{}}
		for _, text := range keysText {
			lines = append(lines, line{text: tint + "  " + text})
		}
	default:
		lines = screen.sample(count, textWidth-2)
	}
	if len(lines) > count {
		lines = lines[:count]
	}
	for len(lines) < count {
		lines = append(lines, line{})
	}
	if screen.menu == nil && !screen.keys {
		screen.drawn = lines
	}
	ink := "\x1b[38;2;" + screen.set.shown("ink").levels() + "m"
	for position := range lines {
		lines[position].text = ink +
			screen.set.explicit(lines[position].text)
		lines[position].fill = screen.set.explicit(lines[position].fill)
	}
	everywhere := []string{screen.set.groundName(screen.page),
		screen.set.inkName(screen.page)}
	marked := []string{steered}
	if group := screen.set.groupOf(steered); group != nil {
		marked = group
	}
	return sampleRows(lines, marked, everywhere, tint, ground, textWidth,
		rowWidth)
}

// unsetNote says so, in the float, when the colourway sets few of the
// page's colours, so a page drawn mostly in derived colours is not taken
// for the colourway's own: a colourway made for the terminal alone sets
// none of vim's.
func (screen *model) unsetNote() string {
	set, own, all := screen.set, 0, 0
	for _, here := range screen.stops {
		if here.page != screen.page {
			continue
		}
		all++
		if _, found := set.roles.get(here.name); found {
			own++
		}
	}
	if own*2 >= all {
		return ""
	}
	return fmt.Sprintf("sets %d of %d here; the rest derived", own, all)
}

// sample is the page showing, as lines width wide.
func (screen *model) sample(count, width int) []line {
	switch screen.page {
	case claudePage:
		return claudeLines(width)
	case markdownPage:
		return markdownLines(screen.set, width)
	case vimPage:
		return vimLines(screen.set, count, width)
	}
	return shellLines(screen.set, screen.text, count, width)
}
