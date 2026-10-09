" fermi-long -- a long gamma-ray burst seen through libtheme's fermi
" profile, 10 kev to 300 gev carried onto visible. the ground is the deep
" indigo at its violet end; the ink is the whole burst as one colour, lit to
" 11 to 1. each of the terminal's colours is the wavelength the screen's own
" colour points at, lit to 6 to 1, or 8.5 for the brights. magenta is in no
" spectrum, so it is the screen's red mixed with violet. made by spectra,
" github.com/janearc/chroma/cmd/spectra, 2026-10-05.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/fermi-long.css;
" change the source, not this file.
"
"   :colorscheme fermi-long-inverted
"
" in a terminal, vim draws these colours when termguicolors is set,
" in your vimrc:
"
"   set termguicolors

hi clear
if exists('syntax_on')
  syntax reset
endif
set background=light
let g:colors_name = 'fermi-long-inverted'

hi Normal guifg=#4a85ad guibg=#e4f687 gui=NONE cterm=NONE
hi NormalFloat guifg=#4a85ad guibg=#d9ee8d gui=NONE cterm=NONE
hi FloatBorder guifg=#a08d00 guibg=#d9ee8d gui=NONE cterm=NONE
hi CursorLine guibg=#daef8c gui=NONE cterm=NONE
hi CursorLineNr guifg=#00d406 gui=bold cterm=bold
hi LineNr guifg=#a08d00 gui=NONE cterm=NONE
hi SignColumn guibg=#e4f687 gui=NONE cterm=NONE
hi Visual guibg=#d5ff3f gui=NONE cterm=NONE
hi VertSplit guifg=#a08d00 gui=NONE cterm=NONE
hi WinSeparator guifg=#a08d00 gui=NONE cterm=NONE
hi StatusLine guifg=#4a85ad guibg=#d9ee8d gui=NONE cterm=NONE
hi StatusLineNC guifg=#a08d00 guibg=#d9ee8d gui=NONE cterm=NONE
hi Pmenu guifg=#4a85ad guibg=#d9ee8d gui=NONE cterm=NONE
hi PmenuSel guifg=#e4f687 guibg=#00d406 gui=NONE cterm=NONE
hi Search guifg=#e4f687 guibg=#6864ff gui=NONE cterm=NONE
hi IncSearch guifg=#e4f687 guibg=#00d406 gui=NONE cterm=NONE
hi MatchParen guifg=#c96900 gui=bold cterm=bold
hi Directory guifg=#c96900 gui=NONE cterm=NONE
hi Folded guifg=#a08d00 guibg=#d9ee8d gui=NONE cterm=NONE
hi NonText guifg=#a08d00 gui=NONE cterm=NONE
hi Whitespace guifg=#a08d00 gui=NONE cterm=NONE
hi Conceal guifg=#a08d00 gui=NONE cterm=NONE
hi Title guifg=#00d406 gui=bold cterm=bold
hi Comment guifg=#a08d00 gui=italic cterm=italic
hi String guifg=#6864ff gui=NONE cterm=NONE
hi Character guifg=#6864ff gui=NONE cterm=NONE
hi Number guifg=#6864ff gui=NONE cterm=NONE
hi Boolean guifg=#6864ff gui=NONE cterm=NONE
hi Identifier guifg=#4a85ad gui=NONE cterm=NONE
hi Function guifg=#c96900 gui=NONE cterm=NONE
hi Statement guifg=#00d406 gui=NONE cterm=NONE
hi Keyword guifg=#00d406 gui=NONE cterm=NONE
hi Operator guifg=#a08d00 gui=NONE cterm=NONE
hi PreProc guifg=#e352ff gui=NONE cterm=NONE
hi Type guifg=#e352ff gui=NONE cterm=NONE
hi Constant guifg=#6864ff gui=NONE cterm=NONE
hi Special guifg=#c96900 gui=NONE cterm=NONE
hi Todo guifg=#e4f687 guibg=#6864ff gui=bold cterm=bold
hi Error guifg=#00a0b8 gui=NONE cterm=NONE
hi markdownH1 guifg=#00d406 gui=bold cterm=bold
hi markdownH2 guifg=#00d406 gui=bold cterm=bold
hi markdownH3 guifg=#00d406 gui=NONE cterm=NONE
hi markdownH4 guifg=#00d406 gui=NONE cterm=NONE
hi markdownH5 guifg=#c96900 gui=NONE cterm=NONE
hi markdownH6 guifg=#c96900 gui=NONE cterm=NONE
hi markdownCode guifg=#e352ff gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#e352ff gui=NONE cterm=NONE
hi markdownLinkText guifg=#c96900 gui=underline cterm=underline
hi markdownUrl guifg=#a08d00 gui=NONE cterm=NONE
hi markdownListMarker guifg=#00d406 gui=NONE cterm=NONE
hi markdownRule guifg=#a08d00 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#a08d00 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#a08d00 gui=NONE cterm=NONE
hi DiagnosticError guifg=#00a0b8 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#6864ff gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#c96900 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#ff595f gui=NONE cterm=NONE
hi DiffAdd guifg=#e352ff gui=NONE cterm=NONE
hi DiffDelete guifg=#00a0b8 gui=NONE cterm=NONE
hi DiffChange guifg=#6864ff gui=NONE cterm=NONE
hi DiffText guifg=#00d406 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#e2ff6f', '#00a0b8', '#e352ff', '#6864ff', '#c96900', '#00d406', '#ff595f', '#015ea7',
  \ '#a08d00', '#006a7e', '#db31ff', '#4b46ff', '#874900', '#007a0d', '#ff383f', '#003860']
