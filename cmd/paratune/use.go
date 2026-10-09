package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/janearc/libtheme-css/colourway"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// pointer is the one line in a program's config that names the colourway
// it draws with: the file, relative to home, the line as a pattern, and
// the line that names another. The patterns stop short of a carriage
// return, so a file with windows line endings keeps them.
type pointer struct {
	file    func(home string) string
	pattern *regexp.Regexp
	line    func(home, name string) string
}

// pointers are ghostty's theme, neovim's colorscheme and glow's style.
// zsh, tmux and claude code have none of their own: they draw from
// ghostty's palette.
var pointers = []pointer{
	{inConfig("ghostty/config"),
		regexp.MustCompile(`(?m)^theme = [^\r\n]*`),
		func(_, name string) string { return "theme = " + name }},
	{inConfig("nvim/init.lua"),
		regexp.MustCompile(`(?m)^vim\.cmd\.colorscheme\("[^"\r\n]*"\)`),
		func(_, name string) string {
			return `vim.cmd.colorscheme("` + name + `")`
		}},
	{inHome("Library/Preferences/glow/glow.yml"),
		regexp.MustCompile(`(?m)^style: [^\r\n]*`),
		func(home, name string) string {
			styles, _ := colourway.Place("glamour", home)
			file := filepath.Join(styles, name+".json")
			return `style: "` + file + `"`
		}},
}

// inConfig is a file in the user's config folder, which XDG_CONFIG_HOME
// moves, where install puts the themes.
func inConfig(file string) func(home string) string {
	return func(home string) string {
		return filepath.Join(colourway.Config(home), file)
	}
}

// inHome is a file at a fixed place in the home, as glow's config is on a
// mac.
func inHome(file string) func(home string) string {
	return func(home string) string { return filepath.Join(home, file) }
}

// plainName is what a colourway's name may hold, so that it can go into a
// config line and a neovim command as it is.
var plainName = regexp.MustCompile(`^[a-z0-9-]+$`)

// use is the switcher: it installs a colourway where the programs read
// it, points each program's config at it, and tells every neovim already
// running to draw with it. Every config is read and checked before any is
// written, so a refusal changes nothing.
//
// A write error partway leaves the files before it switched, each one
// whole, and running it again finishes the job. Ghostty reads its config
// again on command shift comma.
func use(dir, name, home string) error {
	if !plainName.MatchString(name) {
		return fmt.Errorf("%q: a colourway's name is lowercase "+
			"letters, digits and dashes", name)
	}
	if !filepath.IsAbs(home) {
		return fmt.Errorf("no home directory to switch in")
	}
	changes := []change{}
	for _, each := range pointers {
		planned, err := plan(each.file(home), each.pattern,
			each.line(home, name))
		if err != nil {
			return err
		}
		changes = append(changes, planned)
	}
	set, err := load(dir, name)
	if err != nil {
		return err
	}
	if note := set.install(home); !strings.HasPrefix(note, "installed") ||
		strings.Contains(note, "not rendered") {
		return fmt.Errorf("not switched: %s", note)
	}
	for index, each := range changes {
		if err := each.write(); err != nil {
			return err
		}
		fmt.Println("pointed", pointers[index].file(home), "at", name)
	}
	answered, listening := switchNeovims(name)
	fmt.Printf("told %d of %d running neovim to switch\n",
		answered, listening)
	fmt.Println("ghostty: command shift comma reads the config again")
	return nil
}

// change is a config file with its one line replaced, not yet written:
// the file, symlinks followed, as it is and as it will be.
type change struct {
	path        string
	was, will   []byte
	permissions os.FileMode
}

// plan reads a config and replaces the one line that matches a pattern,
// writing nothing. It refuses a file where the line is missing or appears
// more than once.
func plan(path string, pattern *regexp.Regexp, line string) (change, error) {
	target, err := filepath.EvalSymlinks(path)
	if err != nil {
		return change{}, err
	}
	info, err := os.Stat(target)
	if err != nil {
		return change{}, err
	}
	raw, err := os.ReadFile(target)
	if err != nil {
		return change{}, err
	}
	if found := len(pattern.FindAllIndex(raw, -1)); found != 1 {
		return change{}, fmt.Errorf(
			"%s: %d lines like %s, not 1; nothing changed",
			path, found, pattern)
	}
	will := pattern.ReplaceAllLiteral(raw, []byte(line))
	return change{path: target, was: raw, will: will,
		permissions: info.Mode().Perm()}, nil
}

// write puts a change in place. The first time paratune changes a file it
// keeps the file as it was beside it, with .bak-paratune on the end, and
// never writes over that copy, so it is always the file from before.
func (each change) write() error {
	backup := each.path + ".bak-paratune"
	if _, err := os.Stat(backup); errors.Is(err, os.ErrNotExist) {
		err := replaceFile(backup, each.was, each.permissions)
		if err != nil {
			return err
		}
	}
	return replaceFile(each.path, each.will, each.permissions)
}

// replaceFile writes a whole file through a temporary one beside it,
// flushed to disk and then renamed over the old, so the file is either as
// it was or as it will be, never half of each. The temporary file is
// removed if anything fails.
func replaceFile(path string, data []byte, permissions os.FileMode) error {
	pattern := "." + filepath.Base(path) + ".paratune-*"
	temporary, err := os.CreateTemp(filepath.Dir(path), pattern)
	if err != nil {
		return err
	}
	written := func() error {
		if _, err := temporary.Write(data); err != nil {
			return err
		}
		if err := temporary.Chmod(permissions); err != nil {
			return err
		}
		return temporary.Sync()
	}()
	if err := errors.Join(written, temporary.Close()); err != nil {
		os.Remove(temporary.Name())
		return err
	}
	if err := os.Rename(temporary.Name(), path); err != nil {
		os.Remove(temporary.Name())
		return err
	}
	return nil
}

// switchNeovims asks every neovim listening on its default socket to take
// a colourscheme, by evaluating the command in it, not by typing into it,
// and says how many answered of how many were listening.
func switchNeovims(name string) (answered, listening int) {
	who, err := user.Current()
	if err != nil {
		return 0, 0
	}
	pattern := filepath.Join(os.TempDir(), "nvim."+who.Username,
		"*", "nvim.*.0")
	sockets, _ := filepath.Glob(pattern)
	expr := `execute("colorscheme ` + name + `")`
	for _, socket := range sockets {
		ctx, cancel := context.WithTimeout(context.Background(),
			2*time.Second)
		err := exec.CommandContext(ctx, "nvim", "--server", socket,
			"--remote-expr", expr).Run()
		cancel()
		if err == nil {
			answered++
		}
	}
	return answered, len(sockets)
}
