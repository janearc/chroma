package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/janearc/libtheme-css/colourway"
)

// install puts a colourways folder in place, colourways/ unless another
// is named: its sources in the user's own colourways folder, and each
// render where its program reads it. A file of theirs that differs is
// kept and named, unless --replace is asked for.
func install(arguments []string) error {
	flags := flag.NewFlagSet("install", flag.ContinueOnError)
	replace := flags.Bool("replace", false,
		"take ours in place of the user's own files that differ")
	if err := flags.Parse(arguments); err != nil {
		return err
	}
	if flags.NArg() > 1 {
		return fmt.Errorf("one folder, and its flags before it: " +
			"colourway install [--replace] [FOLDER]")
	}
	from := "colourways"
	if flags.NArg() > 0 {
		from = flags.Arg(0)
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	done, err := colourway.Install(from, home, *replace)
	if err == nil || len(done.Written)+len(done.Same)+len(done.Kept) > 0 {
		done.Say(os.Stdout, home)
	}
	return err
}
