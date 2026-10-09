# colourways: putting one in

each colourway is one name across several programs. take the files for
the programs you use; nothing needs to be run. every colourway here has a
file for each program, and an inverted twin.

## all of them at once

from a clone of chroma:

    go run ./cmd/colourway install

or `game run install`. every colourway's source goes in your colourways
folder (`~/.config/colourways`, or `colorways` if that is the one you
have), which is where paratune looks, and each program's file goes where
the program reads it.

a file of yours that differs is kept and named; `--replace` does the expected.

## ghostty

copy `ghostty/NAME` into `~/.config/ghostty/themes/`, then in
`~/.config/ghostty/config`:

    theme = NAME

the ghostty theme also sets the sixteen numbered colours claude code paints
its screen with. 

hi, mitchell, i love your terminal, thank you!

## neovim

copy `nvim/colors/NAME.lua` into `~/.config/nvim/colors/`, then
`:colorscheme NAME`, or in `init.lua`:

    vim.cmd.colorscheme("NAME")

the scheme also sets the sixteen colours of vim's own terminal, so a
shell inside vim (`:terminal`) looks like the terminal around it. but don't do
that, terminals in vim is weird.

## vim

copy `vim/colors/NAME.vim` into `~/.vim/colors/`, then `:colo NAME`, or
in `~/.vimrc`:

    set termguicolors
    color NAME

the vim and neovim colorways are the same, but neovim prefers lua because it
has made choices. the scheme gives exact colours only, so in a terminal vim
needs termguicolors to show it at all; gvim and macvim show it without.

## glow, and anything else that draws markdown with glamour

    glow -s glamour/NAME.json README.md

or copy the file to `~/.config/glamour/` and name it in glow's config:

    style: "~/.config/glamour/NAME.json"

glow 3 draws code blocks in the older 256 colours, so code comes out
near the style's colours rather than exactly on them. please let me know
if you have any issues with this, i haven't really been able to test
anywhere but on my machine.

## inverted

`NAME-inverted` is NAME with every colour inverted, for a screen whose
colours the system inverts (on a mac: accessibility, display, invert
colours). you may find that the inverted scheme is useful just by itself.

## sources

the files under `sources/` are what the rest are made from.

## comparing colourways

`git` unfortunately doesn't know much about radiation, let alone color,
so we needed a way to be able to compare colourways when putting stuff
in git. so, we made `colourway`, which allows one to compare colors
between revisions in git. which, actually, is really cool.

because it's just spectra, if you were the enterprising sort who keeps gamma
ray bursts in git, you could have this lil tool here tell you the difference
between two spectra, and that would be very cool. or, i think so. which is why
it exists.

    go install github.com/janearc/chroma/cmd/colourway@latest
    colourway diff OLD NEW

each of OLD and NEW can be a source, a ghostty theme, a neovim scheme or
a glamour style. to have git use it, set the driver once in a checkout
(`.gitattributes` already names the files):

    git config diff.colourway.command "colourway diff --git"
    git config diff.colourway.textconv "colourway show"

`git diff` then shows what moved; `git log -p` shows each role's hex and
oklch, so a change reads as the colours it changed.

## pictures and the page

`game run pictures` draws every colourway from its source into
`pictures/`; change a source, run it again, and the picture follows.
`game run page` writes the readme as one page,
`bin/colourways.html`, with every colourway a click away at its top.
