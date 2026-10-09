// Command paratune tunes a colourway live, over samples of what is read
// in it, so rivers and smear are seen as the colours move, on the screen
// the reading is done on, not on a card.
//
//	go run . [-colourways DIR] [-name NAME] [-text FILE]
//	go run . [-colourways DIR] -use NAME
//	go run . -install FOLDER [-replace]
//
// A colourway is one css sheet, its source, in DIR/sources, rendered by
// libtheme's dialects into each program's file under DIR.
//
// paratune opens one, shows it on four pages (the shell, claude code,
// markdown as glamour draws it, and vim), and recolours them as the keys
// move: inside tmux through the pane's own style and palette, elsewhere
// by the terminal's OSC 4, 10 and 11.
//
// enter writes the colourway's roles to ~/paratune-kept.txt and stays. u
// undoes a step. s saves the source and renders it again; i puts the
// rendered files where the programs read them; escape leaves and puts
// the pane back.
//
// The frame is bubbletea's, the menu huh's with lipgloss's styles,
// contrast libreadme's and colour libtheme's.
package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/janearc/chroma/internal/age"
	libcolour "github.com/janearc/libreadme/colour"
	"github.com/janearc/libtheme-css/colourway"
	"github.com/janearc/libtheme-css/css"
	"github.com/janearc/libtheme-css/dialects/claude"
	"github.com/janearc/libtheme-css/dialects/ghostty"
)

// colour is one rgb colour, each part nought to 255.
type colour struct{ r, g, b float64 }

// hex is the colour as #rrggbb.
func (value colour) hex() string {
	return fmt.Sprintf("#%02x%02x%02x", int(math.Round(value.r)),
		int(math.Round(value.g)), int(math.Round(value.b)))
}

// levels is the colour as the three numbers a terminal's colour codes
// take, rounded as hex rounds them, so what is drawn is what is kept.
func (value colour) levels() string {
	return fmt.Sprintf("%d;%d;%d", int(math.Round(value.r)),
		int(math.Round(value.g)), int(math.Round(value.b)))
}

// clamp keeps each part inside nought to 255.
func (value colour) clamp() colour {
	inside := func(part float64) float64 {
		return math.Max(0, math.Min(255, part))
	}
	return colour{inside(value.r), inside(value.g), inside(value.b)}
}

// reader is the colour as libreadme takes it.
func (value colour) reader() libcolour.RGB {
	return libcolour.RGB{R: value.r, G: value.g, B: value.b}
}

// luminance is wcag's relative luminance, as libreadme works it out.
func (value colour) luminance() float64 {
	return libcolour.Luminance(value.reader())
}

// contrast is wcag's ratio between two colours, as libreadme works it
// out.
func contrast(first, second colour) float64 {
	return libcolour.Contrast(first.reader(), second.reader())
}

// xterm is a palette colour as xterm's 256 colours have it, for a number
// the colourway does not set.
func xterm(index int) colour {
	base := []colour{{0, 0, 0}, {205, 0, 0}, {0, 205, 0}, {205, 205, 0},
		{0, 0, 238}, {205, 0, 205}, {0, 205, 205}, {229, 229, 229},
		{127, 127, 127}, {255, 0, 0}, {0, 255, 0}, {255, 255, 0},
		{92, 92, 255}, {255, 0, 255}, {0, 255, 255}, {255, 255, 255}}
	switch {
	case index < 16:
		return base[index]
	case index < 232:
		cube := index - 16
		level := func(step int) float64 {
			if step == 0 {
				return 0
			}
			return float64(55 + step*40)
		}
		return colour{
			level(cube / 36), level(cube / 6 % 6), level(cube % 6)}
	default:
		grey := float64(8 + (index-232)*10)
		return colour{grey, grey, grey}
	}
}

// tuning is the colourway being tuned: its source, its roles, the
// directory of colourways it lives in, whether it has moved since it was
// last saved, and, if saving would lose something the file holds, where.
type tuning struct {
	source colourway.Source
	roles  *roles
	dir    string
	dirty  bool
	lossy  string

	groups [][]string // roles that move together, each group a list

	history []*roles // every colour before each step, for u to walk back
	saved   *roles   // every colour as the source on disk has it

	// the roles as libtheme resolves them, and the roles and the change
	// it was worked from
	resolved   colourway.Resolved
	resolvedOf *roles
	resolvedAt int
}

// resolution is the colourway's roles resolved by libtheme, which is what
// the editor and markdown are rendered from: worked out again only when
// a role has changed since.
func (set *tuning) resolution() colourway.Resolved {
	if set.resolvedOf != set.roles || set.resolvedAt != set.roles.changes ||
		set.resolved.Roles == nil {
		set.resolved = colourway.Resolve(set.roles.sheet())
		set.resolvedOf, set.resolvedAt = set.roles, set.roles.changes
	}
	return set.resolved
}

// general is whether a role is in libtheme's general vocabulary, which
// Resolve derives when a colourway does not set it.
func general(name string) bool {
	return slices.Contains(colourway.Grounds, name) ||
		slices.Contains(colourway.Text, name) ||
		slices.Contains(colourway.States, name)
}

// derived is whether Resolve may give a role its colour: a general role,
// or one of claude code's, which follow the general roles by the claude
// dialect's table.
func derived(name string) bool {
	return general(name) || strings.HasPrefix(name, "claude-")
}

// has is whether a role has a colour from the colourway: its own, or one
// libtheme derives for it.
func (set *tuning) has(name string) bool {
	if _, found := set.roles.get(name); found {
		return true
	}
	if !general(name) {
		return false
	}
	_, found := set.resolution().Of(name)
	return found
}

// load opens a colourway by name from a directory of colourways.
func load(dir, name string) (*tuning, error) {
	path := filepath.Join(dir, "sources", name+".css")
	source, err := colourway.Read(path)
	if err != nil {
		return nil, err
	}
	if _, found := source.Roles.Get("ground"); !found {
		return nil, fmt.Errorf("%s has no ground", name)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	roles := rolesOf(source.Roles)
	return &tuning{source: source, roles: roles, saved: roles.copied(),
		dir: dir, lossy: unheld(string(raw)),
		groups: readGroups(dir, name)}, nil
}

// unheld is what saving this source with nothing moved would lose: the
// first line a save would not write back the same, or nothing. A source
// paratune can hold is one it writes back byte for byte, so this is the
// whole of the check, not a list of things it knows to look for.
func unheld(text string) string {
	about := colourway.Opening(text)
	held := rolesOf(css.Read(text).Roles).sheet().String()
	again := colourway.Commented(about) + held
	if again == text {
		return ""
	}
	have, want := strings.Split(text, "\n"), strings.Split(again, "\n")
	for number, line := range have {
		if number >= len(want) || line != want[number] {
			return fmt.Sprintf("line %d: %s", number+1,
				strings.TrimSpace(line))
		}
	}
	return fmt.Sprintf("line %d: the end of the file", len(have))
}

// colourwayNames are the colourways in a directory, by name.
func colourwayNames(dir string) []string {
	paths, _ := filepath.Glob(filepath.Join(dir, "sources", "*.css"))
	names := []string{}
	for _, path := range paths {
		names = append(names,
			strings.TrimSuffix(filepath.Base(path), ".css"))
	}
	sort.Strings(names)
	return names
}

// paintedNumbers are the palette numbers paratune sets in the pane: the
// sixteen and claude code's.
func paintedNumbers() []int {
	numbers := []int{}
	for number := range ghostty.Ansi {
		numbers = append(numbers, number)
	}
	for _, slot := range claude.Slots {
		if slot.Number >= 16 {
			numbers = append(numbers, slot.Number)
		}
	}
	return numbers
}

// paletteOf is the colour the colourway gives each painted number.
func (set *tuning) paletteOf(number int) colour {
	if number < 16 {
		return set.shown(ghostty.Ansi[number])
	}
	for _, slot := range claude.Slots {
		if slot.Number == number {
			return set.shown(slot.Role)
		}
	}
	return xterm(number)
}

// paint puts the pane in the colourway's colours: its ink and ground as
// the pane's style, the sixteen and claude code's numbers as its own
// palette, so every page shows them where the programs would.
//
// It sets the pane's options and never selects the pane: select-pane
// would take her focus, and her typing with it. Inside tmux with no pane
// named, it paints nothing, rather than her pane.
func paint(set *tuning) {
	ink, ground := set.shown("ink"), set.shown("ground")
	pane := os.Getenv("TMUX_PANE")
	if os.Getenv("TMUX") != "" && pane == "" {
		return
	}
	if os.Getenv("TMUX") == "" {
		fmt.Printf("\x1b]10;%s\x07\x1b]11;%s\x07",
			ink.hex(), ground.hex())
		for _, number := range paintedNumbers() {
			fmt.Printf("\x1b]4;%d;%s\x07",
				number, set.paletteOf(number).hex())
		}
		return
	}
	exec.Command("tmux", paintArgs(pane, set)...).Run()
}

// paintArgs is the whole paint as one call to tmux, its commands joined
// by tmux's own separator: one process, not one for every colour.
func paintArgs(pane string, set *tuning) []string {
	style := "fg=" + set.shown("ink").hex() +
		",bg=" + set.shown("ground").hex()
	args := []string{"set", "-p", "-t", pane, "window-style", style,
		";", "set", "-p", "-t", pane, "window-active-style", style}
	for _, number := range paintedNumbers() {
		args = append(args, ";", "set", "-p", "-t", pane,
			fmt.Sprintf("pane-colours[%d]", number),
			set.paletteOf(number).hex())
	}
	return args
}

// unpaint puts the pane back as the terminal has it.
func unpaint() {
	pane := os.Getenv("TMUX_PANE")
	if os.Getenv("TMUX") != "" && pane == "" {
		return
	}
	if os.Getenv("TMUX") == "" {
		fmt.Print("\x1b]110\x07\x1b]111\x07\x1b]104\x07")
		return
	}
	exec.Command("tmux", "set", "-pu", "-t", pane, "window-style").Run()
	exec.Command("tmux", "set", "-pu", "-t", pane,
		"window-active-style").Run()
	for _, number := range paintedNumbers() {
		option := fmt.Sprintf("pane-colours[%d]", number)
		exec.Command("tmux", "set", "-p", "-t", pane,
			"-u", option).Run()
	}
}

// build and built are stamped by game build: the commit, and the
// commit's time. --age prints them.
var build, built = "dev", ""

// main switches a colourway when -use is given, and otherwise opens paratune.
func main() {
	if len(os.Args) > 1 &&
		(os.Args[1] == "--age" || os.Args[1] == "version") {
		fmt.Println(age.Of("paratune", build, built, time.Now()))
		return
	}
	home, _ := os.UserHomeDir()
	dir := flag.String("colourways", colourway.Folder(home),
		"the directory of colourways: sources/, "+
			"and what they render into")
	name := flag.String("name", "twilight-deep", "the colourway to open")
	textFile := flag.String("text", "",
		"the file shown in less on the shell page; "+
			"paratune has one of its own")
	switchTo := flag.String("use", "",
		"switch ghostty, neovim and glow to a colourway, and leave")
	installFrom := flag.String("install", "",
		"put a colourways folder in place, in your own "+
			"colourways folder (not -colourways) and where "+
			"every program reads it, and leave")
	replace := flag.Bool("replace", false,
		"with -install, take its files in place of yours that differ")
	flag.Parse()

	if *replace && *installFrom == "" {
		fmt.Fprintln(os.Stderr,
			"paratune: -replace goes with -install FOLDER")
		os.Exit(2)
	}

	if *installFrom != "" {
		done, err := colourway.Install(*installFrom, home, *replace)
		touched := len(done.Written) + len(done.Same) + len(done.Kept)
		if err == nil || touched > 0 {
			done.Say(os.Stdout, home)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "paratune:", err)
			os.Exit(1)
		}
		return
	}

	if *switchTo != "" {
		if err := use(*dir, *switchTo, home); err != nil {
			fmt.Fprintln(os.Stderr, "paratune:", err)
			os.Exit(1)
		}
		return
	}

	set, err := load(*dir, *name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "paratune:", err)
		os.Exit(1)
	}
	raw := []byte(sampleText)
	if *textFile != "" {
		raw, err = os.ReadFile(*textFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "paratune:", err)
			os.Exit(1)
		}
	}
	screen := &model{home: home, text: string(raw), set: set,
		stops: stopsOf(set.roles), painter: paint}
	screen.steering = screen.first(shellPage)
	paint(set)
	defer unpaint()
	program := tea.NewProgram(screen,
		tea.WithColorProfile(colorprofile.TrueColor))
	if _, err := program.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "paratune:", err)
	}
}
