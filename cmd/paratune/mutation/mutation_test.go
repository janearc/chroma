//go:build mutation

// Package mutation is paratune's mutation tests. Each mutant changes
// paratune on purpose, in a copy, the way a mistake would, and the test
// fails unless paratune's own tests then fail. A mutant that survives is
// a claim the unit tests do not check. They are slow, so they sit behind
// a tag:
//
//	GOWORK=off go test -tags mutation ./mutation
//
// A mutant whose code has moved on matches nothing, and fails as stale
// rather than passing unexamined.
package mutation

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// mutant is one mistake: the file, the code as it is, the code as the
// mistake would leave it, and the claim it breaks.
type mutant struct {
	file, from, to, breaks string
}

// mutants are the mistakes, one for each claim paratune makes.
var mutants = []mutant{
	{"main.go", "if again == text {", "if true {",
		"a source a save would change is refused"},
	{"model.go", `if set.lossy != "" {`, "if false {",
		"save writes nothing for a lossy source"},
	{"model.go", "set.dirty, set.saved = false, set.roles.copied()",
		"set.saved = set.roles.copied()",
		"a save leaves the colourway clean"},
	{"model.go", "if set.dirty {\n\t\treturn \"unsaved: s saves first\"",
		"if false {\n\t\treturn \"unsaved: s saves first\"",
		"install refuses unsaved moves"},
	{"model.go", "\t\t\tcontinue\n\t\t}\n\t\tpaths = append(paths, each.Path)",
		"\t\t}\n\t\tpaths = append(paths, each.Path)",
		"install copies only what the render wrote"},
	{"model.go", "change(screen.set.shown(name)).clamp())",
		"change(screen.set.shown(name)))",
		"a move keeps every lamp in nought to 255"},
	{"model.go", "\tscreen.set.dirty = true\n\tscreen.repaint()",
		"\tscreen.repaint()",
		"a move marks the colourway unsaved"},
	{"model.go", "\tif screen.opened == nil {",
		"\tif screen.set.dirty {\n\t\treturn\n\t}\n\tif screen.opened == nil {",
		"unsaved moves never stop a colourway being changed"},
	{"model.go", "screen.opened[screen.set.source.Name] = screen.set",
		"_ = screen.set", "a colourway left keeps its unsaved moves"},
	{"model.go", "if here.name == ink {", "if here.name == ink+\"x\" {",
		"a page opens on its ink"},
	{"model.go", "screen.body(screen.height, ", "screen.body(screen.height-1, ",
		"the frame is the pane's height exactly"},
	{"float.go", "\tif screen.busy {\n", "\tif false {\n",
		"while keys are pressed the float is its field's row"},
	{"float.go", "\t\tscreen.busy = false\n", "",
		"a quiet minute brings the float back"},
	{"float.go", "if since >= quietAfter {", "if true {",
		"a look before the minute is up changes nothing"},
	{"float.go", "box{left: rowWidth - width - 1,", "box{left: rowWidth - width,",
		"the float keeps every row the pane's width"},
	{"mouse.go", "screen.keys = y == screen.placed.help", "screen.keys = true",
		"only ? help on the float opens the keys"},
	{"main.go", "return general(name) || strings.HasPrefix(name, \"claude-\")", "return general(name)",
		"claude code's colours show as libtheme derives them"},
	{"model.go", "\tscreen.walking = nil\n", "",
		"a walk ends with its colourway"},
	{"float.go", "short := strings.Repeat(\" \", placed.left-ansi.StringWidth(left))", "short := \"\"",
		"a wide glyph at the float's edge leaves the row its width"},
	{"float.go", "\t\trows = append(rows, tint+share)\n", "",
		"the float says when the colourway sets few"},
	{"mouse.go", "case fg.hex() == ground.hex() && bg.hex() != ground.hex():", "case false:",
		"a letter cut out in the ground steers the colour behind it"},
	{"mouse.go", "next = (at + 1) % len(matching)", "next = at",
		"a click again on a shared colour steers the next role"},
	{"model.go", "drawn := map[string]bool{screen.set.groundName(screen.page): true,", "drawn := map[string]bool{\"\": true,",
		"tab reaches the page's own ground on every page"},
	{"whole.go", "\t\tif isGround(name) {\n\t\t\tcontinue\n\t\t}\n", "",
		"> and < leave the grounds"},
	{"mouse.go", "\t\tfg, space = ink, true\n", "\t\tfg, space = ink, false\n",
		"a click on bare ground steers the ground"},
	{"groups.go", "\tcase slices.Contains(everywhere, name):\n", "\tcase slices.Contains(everywhere, name+\"x\"):\n",
		"g never groups the page's own ground or ink"},
	{"model.go", "\t\tset.carry(name, before)\n", "\t\t_ = before\n",
		"a move of a grouped colour moves its group"},
	{"groups.go", "return !slices.Contains(wholeKeys, message.String())", "return !slices.Contains(wholeKeys, message.String()+\"x\")",
		"a move of every colour is not carried to a group"},
	{"model.go", "writeGroups(set.dir, set.source.Name, set.groups)",
		"writeGroups(set.dir, set.source.Name, nil)",
		"groups are kept when a colourway is saved"},
	{"markdown.go", "\tdelete(styles.Registry, glamourChroma)\n", "\tdelete(styles.Registry, glamourChroma+\"x\")\n",
		"code on the markdown page follows a change to its colour"},
	{"markdown.go", "gstyle.Of(set.resolution().Roles)", "gstyle.Of(set.roles.sheet())",
		"the markdown page is drawn from the resolved roles"},
	{"mouse.go", "screen.steerAt(screen.drawn[y], x-2)", "screen.steerLine(screen.drawn[y])",
		"a click steers the colour under the pointer"},
	{"mouse.go", "\t\t\tif reverse {\n\t\t\t\tfg, bg = bg, fg\n\t\t\t}\n", "",
		"reverse video swaps a cell's colours"},
	{"model.go", "if screen.picking && (key == \"o\" || key == \"esc\") {", "if screen.picking && key == \"o\" {",
		"escape closes the picker without leaving"},
	{"picker.go", "top := screen.height - pickerRows", "top := screen.height - pickerRows + 1",
		"a click on the picker takes the colour drawn under it"},
	{"picker.go", "if contrast(value, against) < reading.Contrast.ChipMin {", "if contrast(value, against) > reading.Contrast.ChipMin {",
		"the picker marks what will not read"},
	{"mouse.go", "screen.picking = screen.picking || y == screen.placed.field", "screen.picking = screen.picking",
		"the float's field opens the picker"},
	{"model.go", "(len(drawn) == 0 || drawn[here.name])", "(len(drawn) == 0 || true)",
		"tab passes over colours the page does not draw"},
	{"shell.go", "\t\tshows = append(shows, prefix+name)\n", "",
		"the colour test shows each of the sixteen"},
	{"float.go", "if ink.luminance() > ground.luminance() {", "if ink.luminance() < ground.luminance() {",
		"the float's ground moves away from its marks"},
	{"markdown.go", "if name == \"\" && strings.TrimSpace(token.Value) != \"\" {", "if name == \"\" {",
		"a code line's spaces do not show the block's text colour"},
	{"readable.go", "\t\tif band.Contains(contrast(value, ground)) {\n\t\t\tcontinue\n\t\t}\n", "",
		"r leaves a colour that reads as it is"},
	{"readable.go", "\t\t\tband = proseBand()\n", "",
		"r puts the page's ink in the prose band"},
	{"readable.go", "if here.page != screen.page || here.ground {", "if here.page != screen.page {",
		"r never moves a ground"},
	{"readable.go", "\t\tset.dirty = true\n", "",
		"r makes the colourway unsaved"},
	{"controls.go", "\tif band.Contains(contrast(ink, ground)) {\n\t\treturn ink\n\t}\n", "",
		"marks keep an ink that already reads"},
	{"controls.go", "if contrast(black, ground) > contrast(white, ground) {", "if false {",
		"marks fall back to whichever of black and white reads"},
	{"controls.go",
		"ground.reader(),\n\t\treading.Contrast.ChipMin,",
		"ground.reader(),\n\t\t0,",
		"the field's fill is moved until the label reads"},
	{"controls.go", "if contrast(value, ground) >= reading.Contrast.ChipMin {", "if true {",
		"a field the reader cannot read gets a fill"},
	{"controls.go", "text := inkCode(value) + label", "text := label",
		"the field is written in its own colour"},
	{"model.go", "func (screen *model) repaint() {\n\tscreen.paintDue = true\n",
		"func (screen *model) repaint() {\n\tscreen.paintDue = true\n\tif screen.painter != nil {\n\t\tscreen.painter(screen.set)\n\t}\n",
		"a press does not wait for the pane to be painted"},
	{"model.go", "\tscreen.paintWaiting = true\n\treturn tea.Tick", "\treturn tea.Tick",
		"one paint is on its way at a time"},
	{"model.go", "if screen.paintDue && screen.painter != nil {", "if screen.painter != nil {",
		"a paint with nothing due paints nothing"},
	{"main.go", "\t\targs = append(args, \";\", \"set\"", "\t\targs = append(args, \"set\"",
		"a paint is one call, its commands joined"},
	{"roles.go", "\tall.values[name] = value\n\tall.changes++\n", "\tall.values[name] = value\n",
		"a derived role is worked out again after a move"},
	{"stops.go", "found && derived(name) {", "found && false {",
		"a general role shows as libtheme derives it"},
	{"stops.go", "\t\tcurrent = source\n", "\t\tbreak\n",
		"where a role comes from follows each step"},
	{"main.go", "\t_, found := set.resolution().Of(name)\n\treturn found\n", "\treturn false\n",
		"markdown follows a derived role, as glamour does"},
	{"model.go", "\tif screen.keys {\n\t\tif key == \"?\" || key == \"esc\" {", "\tif false {\n\t\tif key == \"?\" || key == \"esc\" {",
		"keys do nothing while the keys panel is open"},
	{"model.go", "\tcase \"?\":\n\t\tscreen.keys = true\n", "",
		"? opens the keys"},
	{"whole.go", "current.at = (current.at + step + count)",
		"current.at = (current.at + 1 + count)",
		"shift-n goes back through the palette"},
	{"whole.go", "\tscreen.set.dirty = true\n\tscreen.note = fmt.Sprintf(\"palette", "\tscreen.note = fmt.Sprintf(\"palette",
		"a palette step makes the colourway unsaved"},
	{"whole.go", "\t\tscreen.walking = current\n", "",
		"a walk keeps its own copy of the palette"},
	{"stops.go", "notRole := source == \"xterm\" || strings.HasPrefix(source, \"ground and \")", "notRole := false",
		"where a role comes from stops at a palette or a mix"},
	{"float.go", "shown := pageNames[screen.page]", `shown := ""`,
		"the steering line names the page"},
	{"menu.go", "index := y - 1", "index := y - 2", "a click in the menu lands on its line"},
	{"mouse.go", "screen.steerAt(screen.drawn[y], x-2)", "screen.steerAt(screen.drawn[max(y-1, 0)], x-2)",
		"a click on a line steers that line"},
	{"controls.go", "\tset.dirty = true\n\treturn \"pasted \"",
		"\treturn \"pasted \"", "a paste marks the colourway unsaved"},
	{"controls.go", "case !here.ground:", "case true:",
		"a ground's foreground register goes to its ink"},
	{"roles.go", "\t\tall.order = append(all.order, name)\n", "",
		"what is put is what is saved"},
	{"roles.go", "lamps.R * 255", "lamps.R * 254",
		"a colour comes back from a swatch the byte it was"},
	{"pages.go", `case "", "0", "49":`, `case "", "0":`,
		"a band stays its ground after a 49"},
	{"pages.go",
		"slices.Contains(steered, name) &&\n\t\t\t\t!slices.Contains(everywhere, name)",
		"slices.Contains(steered, name)",
		"the page's own ink and ground mark no lines"},
	{"pages.go", "max(rowWidth-textWidth, 0)", "0",
		"a band is the row's width, gap and all"},
	{"stops.go", "ground: grounds[entry.Role]})", "ground: false})",
		"vim's grounds are grounds"},
	{"stops.go", "return xterm(slot.Number)", "return xterm(0)",
		"an unset claude colour shows xterm's"},
	{"stops.go", "\t\t\t\treturn fallback\n",
		"\t\t\t\treturn \"ink\"\n",
		"the markdown page shows what glamour will render"},
	{"stops.go", `return set.shown(strings.TrimPrefix(name, "nvim-"))`,
		`return set.shown("ink")`,
		"vim's own colour falls back to the shared one"},
	{"menu.go", `key.String() == "esc" || key.String() == "]"`,
		`key.String() == "esc"`, "] closes the menu"},
	{"mouse.go", "next = names[(position+1)%len(names)]", "next = names[position]",
		"clicking a line again steers its next colour"},
	{"mouse.go", "screen.nudge(lighter)", "screen.nudge(darker)",
		"the wheel up makes the colour steered lighter"},
	{"mouse.go", "x >= screen.width-2))", "x >= screen.width-4))",
		"the gradient's oklab side gives oklab's hue"},
	{"model.go", "set.roles, set.history = set.history[last], set.history[:last]",
		"set.history = set.history[:last]",
		"undo puts every colour back as it was"},
	{"model.go", " && !set.roles.same(snapshot)", "",
		"a key that moves no colour is not kept as a step"},
	{"model.go", "set.dirty = !set.roles.same(set.saved)", "set.dirty = true",
		"undone back to the save counts as saved"},
	{"model.go", "!undoing(message) && ", "",
		"an undo is not kept as a step of its own"},
	{"pages.go", `said = append(said, lamps(38, number-30))`,
		`said = append(said, part)`,
		"a page draws the sixteen from the colourway, not the pane"},
}

// root is paratune's package, the directory above this one, and module
// the repository's module, which paratune lives inside.
const (
	root   = ".."
	module = "../../.."
)

// copied is paratune's package in a directory of its own: every go file
// at its top, tests included, and nothing below it, in a module that
// takes the repository's libraries from the repository itself.
func copied(t *testing.T) string {
	t.Helper()
	into := t.TempDir()
	if err := moduleAround(into); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") {
			continue
		}
		raw, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(into, name), raw, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return into
}

// moduleAround makes a directory a module that requires the repository's
// module from its own folder, with the repository's go.sum, so the copy
// builds against the libraries beside it and nothing fetched.
func moduleAround(dir string) error {
	here, err := filepath.Abs(module)
	if err != nil {
		return err
	}
	sums, err := os.ReadFile(filepath.Join(here, "go.sum"))
	if err != nil {
		return err
	}
	text := "module mutant\n\ngo 1.26\n\n" +
		"require github.com/janearc/chroma v0.0.0\n\n" +
		"replace github.com/janearc/chroma => " + here + "\n"
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(text), 0o644); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "go.sum"), sums, 0o644)
}

// tests runs paratune's tests in a directory, and says whether they
// passed, and what they said. A mutant that does not build fails them
// too, and is told apart by what go test says.
func tests(dir string) (bool, string) {
	command := exec.Command("go", "test", "-count=1", ".")
	command.Dir = dir
	command.Env = append(os.Environ(), "GOWORK=off", "GOMAXPROCS=2",
		"GOFLAGS=-p=2 -mod=mod", "GOPROXY=off")
	out, err := command.CombinedOutput()
	return err == nil, string(out)
}

// failing are the names of the tests that failed, from go test's output.
func failing(out string) string {
	names := []string{}
	for _, line := range strings.Split(out, "\n") {
		if name, found := strings.CutPrefix(line, "--- FAIL: "); found {
			names = append(names, strings.Fields(name)[0])
		}
	}
	return strings.Join(names, ", ")
}

// TestMutants runs paratune's tests unchanged, which must pass, then
// once with each mistake in, each of which must make them fail.
func TestMutants(t *testing.T) {
	if passed, out := tests(copied(t)); !passed {
		t.Fatalf("paratune's tests fail with no mutant in:\n%s", out)
	}
	for _, each := range mutants {
		t.Run(each.breaks, func(t *testing.T) {
			dir := copied(t)
			path := filepath.Join(dir, each.file)
			raw, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if count := strings.Count(string(raw), each.from); count != 1 {
				t.Fatalf("stale: %q is in %s %d times, not once", each.from,
					each.file, count)
			}
			text := strings.Replace(string(raw), each.from, each.to, 1)
			if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
				t.Fatal(err)
			}
			passed, out := tests(dir)
			switch {
			case strings.Contains(out, "[build failed]"),
				strings.Contains(out, "[setup failed]"):
				t.Errorf("broken: the mutant does not build, so it tests "+
					"nothing:\n%s", out)
			case passed:
				t.Errorf("survived: the tests pass when it is broken that %s",
					each.breaks)
			default:
				t.Logf("caught by: %s", failing(out))
			}
		})
	}
}
