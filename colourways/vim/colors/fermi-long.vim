" fermi-long -- a long gamma-ray burst seen through libtheme's fermi
" profile, 10 kev to 300 gev carried onto visible. the ground is the deep
" indigo at its violet end; the ink is the whole burst as one colour, lit to
" 11 to 1. each of the terminal's colours is the wavelength the screen's own
" colour points at, lit to 6 to 1, or 8.5 for the brights. magenta is in no
" spectrum, so it is the screen's red mixed with violet. made by spectra,
" github.com/janearc/chroma/cmd/spectra, 2026-10-05.
"
" rendered by colourway from sources/fermi-long.css;
" change the source, not this file.
"
"   :colorscheme fermi-long
"
" in a terminal, vim draws these colours when termguicolors is set,
" in your vimrc:
"
"   set termguicolors

hi clear
if exists('syntax_on')
  syntax reset
endif
set background=dark
let g:colors_name = 'fermi-long'

hi Normal guifg=#b57a52 guibg=#1b0978 gui=NONE cterm=NONE
hi NormalFloat guifg=#b57a52 guibg=#221b78 gui=NONE cterm=NONE
hi FloatBorder guifg=#5f72ff guibg=#221b78 gui=NONE cterm=NONE
hi CursorLine guibg=#211978 gui=NONE cterm=NONE
hi CursorLineNr guifg=#ff2bf9 gui=bold cterm=bold
hi LineNr guifg=#5f72ff gui=NONE cterm=NONE
hi SignColumn guibg=#1b0978 gui=NONE cterm=NONE
hi Visual guibg=#2a00c0 gui=NONE cterm=NONE
hi VertSplit guifg=#5f72ff gui=NONE cterm=NONE
hi WinSeparator guifg=#5f72ff gui=NONE cterm=NONE
hi StatusLine guifg=#b57a52 guibg=#221b78 gui=NONE cterm=NONE
hi StatusLineNC guifg=#5f72ff guibg=#221b78 gui=NONE cterm=NONE
hi Pmenu guifg=#b57a52 guibg=#221b78 gui=NONE cterm=NONE
hi PmenuSel guifg=#1b0978 guibg=#ff2bf9 gui=NONE cterm=NONE
hi Search guifg=#1b0978 guibg=#979b00 gui=NONE cterm=NONE
hi IncSearch guifg=#1b0978 guibg=#ff2bf9 gui=NONE cterm=NONE
hi MatchParen guifg=#3696ff gui=bold cterm=bold
hi Directory guifg=#3696ff gui=NONE cterm=NONE
hi Folded guifg=#5f72ff guibg=#221b78 gui=NONE cterm=NONE
hi NonText guifg=#5f72ff gui=NONE cterm=NONE
hi Whitespace guifg=#5f72ff gui=NONE cterm=NONE
hi Conceal guifg=#5f72ff gui=NONE cterm=NONE
hi Title guifg=#ff2bf9 gui=bold cterm=bold
hi Comment guifg=#5f72ff gui=italic cterm=italic
hi String guifg=#979b00 gui=NONE cterm=NONE
hi Character guifg=#979b00 gui=NONE cterm=NONE
hi Number guifg=#979b00 gui=NONE cterm=NONE
hi Boolean guifg=#979b00 gui=NONE cterm=NONE
hi Identifier guifg=#b57a52 gui=NONE cterm=NONE
hi Function guifg=#3696ff gui=NONE cterm=NONE
hi Statement guifg=#ff2bf9 gui=NONE cterm=NONE
hi Keyword guifg=#ff2bf9 gui=NONE cterm=NONE
hi Operator guifg=#5f72ff gui=NONE cterm=NONE
hi PreProc guifg=#1cad00 gui=NONE cterm=NONE
hi Type guifg=#1cad00 gui=NONE cterm=NONE
hi Constant guifg=#979b00 gui=NONE cterm=NONE
hi Special guifg=#3696ff gui=NONE cterm=NONE
hi Todo guifg=#1b0978 guibg=#979b00 gui=bold cterm=bold
hi Error guifg=#ff5f47 gui=NONE cterm=NONE
hi markdownH1 guifg=#ff2bf9 gui=bold cterm=bold
hi markdownH2 guifg=#ff2bf9 gui=bold cterm=bold
hi markdownH3 guifg=#ff2bf9 gui=NONE cterm=NONE
hi markdownH4 guifg=#ff2bf9 gui=NONE cterm=NONE
hi markdownH5 guifg=#3696ff gui=NONE cterm=NONE
hi markdownH6 guifg=#3696ff gui=NONE cterm=NONE
hi markdownCode guifg=#1cad00 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#1cad00 gui=NONE cterm=NONE
hi markdownLinkText guifg=#3696ff gui=underline cterm=underline
hi markdownUrl guifg=#5f72ff gui=NONE cterm=NONE
hi markdownListMarker guifg=#ff2bf9 gui=NONE cterm=NONE
hi markdownRule guifg=#5f72ff gui=NONE cterm=NONE
hi markdownBlockquote guifg=#5f72ff gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#5f72ff gui=NONE cterm=NONE
hi DiagnosticError guifg=#ff5f47 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#979b00 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#3696ff gui=NONE cterm=NONE
hi DiagnosticHint guifg=#00a6a0 gui=NONE cterm=NONE
hi DiffAdd guifg=#1cad00 gui=NONE cterm=NONE
hi DiffDelete guifg=#ff5f47 gui=NONE cterm=NONE
hi DiffChange guifg=#979b00 gui=NONE cterm=NONE
hi DiffText guifg=#ff2bf9 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#1d0090', '#ff5f47', '#1cad00', '#979b00', '#3696ff', '#ff2bf9', '#00a6a0', '#fea158',
  \ '#5f72ff', '#ff9581', '#24ce00', '#b4b900', '#78b6ff', '#ff85f2', '#00c7c0', '#ffc79f']
