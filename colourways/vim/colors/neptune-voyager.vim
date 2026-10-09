" neptune-voyager -- neptune as voyager 2 showed it in august 1989, from
" the green and orange filters of its narrow angle camera (nasa/jpl,
" pia01492). the ground is the limb in shadow, greyed to her hue rule; the
" ink is the white clouds, lit to 11 to 1; the selection and the blues are
" the disc; dim text is the great dark spot. neptune has no red, green or
" yellow, so those are the screen's own wavelengths, as in fermi-long.
" samples in cmd/spectra/data/neptune-voyager.css. made by spectra,
" github.com/janearc/chroma/cmd/spectra, 2026-10-05.
"
" rendered by colourway from sources/neptune-voyager.css;
" change the source, not this file.
"
"   :colorscheme neptune-voyager
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
let g:colors_name = 'neptune-voyager'

hi Normal guifg=#9acfff guibg=#111523 gui=NONE cterm=NONE
hi NormalFloat guifg=#9acfff guibg=#192030 gui=NONE cterm=NONE
hi FloatBorder guifg=#5474ff guibg=#192030 gui=NONE cterm=NONE
hi CursorLine guibg=#181e2e gui=NONE cterm=NONE
hi CursorLineNr guifg=#ff29f9 gui=bold cterm=bold
hi LineNr guifg=#5474ff gui=NONE cterm=NONE
hi SignColumn guibg=#111523 gui=NONE cterm=NONE
hi Visual guibg=#203480 gui=NONE cterm=NONE
hi VertSplit guifg=#5474ff gui=NONE cterm=NONE
hi WinSeparator guifg=#5474ff gui=NONE cterm=NONE
hi StatusLine guifg=#9acfff guibg=#192030 gui=NONE cterm=NONE
hi StatusLineNC guifg=#5474ff guibg=#192030 gui=NONE cterm=NONE
hi Pmenu guifg=#9acfff guibg=#192030 gui=NONE cterm=NONE
hi PmenuSel guifg=#111523 guibg=#ff29f9 gui=NONE cterm=NONE
hi Search guifg=#111523 guibg=#979a00 gui=NONE cterm=NONE
hi IncSearch guifg=#111523 guibg=#ff29f9 gui=NONE cterm=NONE
hi MatchParen guifg=#6a8fff gui=bold cterm=bold
hi Directory guifg=#6a8fff gui=NONE cterm=NONE
hi Folded guifg=#5474ff guibg=#192030 gui=NONE cterm=NONE
hi NonText guifg=#5474ff gui=NONE cterm=NONE
hi Whitespace guifg=#5474ff gui=NONE cterm=NONE
hi Conceal guifg=#5474ff gui=NONE cterm=NONE
hi Title guifg=#ff29f9 gui=bold cterm=bold
hi Comment guifg=#5474ff gui=italic cterm=italic
hi String guifg=#979a00 gui=NONE cterm=NONE
hi Character guifg=#979a00 gui=NONE cterm=NONE
hi Number guifg=#979a00 gui=NONE cterm=NONE
hi Boolean guifg=#979a00 gui=NONE cterm=NONE
hi Identifier guifg=#9acfff gui=NONE cterm=NONE
hi Function guifg=#6a8fff gui=NONE cterm=NONE
hi Statement guifg=#ff29f9 gui=NONE cterm=NONE
hi Keyword guifg=#ff29f9 gui=NONE cterm=NONE
hi Operator guifg=#5474ff gui=NONE cterm=NONE
hi PreProc guifg=#1cac00 gui=NONE cterm=NONE
hi Type guifg=#1cac00 gui=NONE cterm=NONE
hi Constant guifg=#979a00 gui=NONE cterm=NONE
hi Special guifg=#6a8fff gui=NONE cterm=NONE
hi Todo guifg=#111523 guibg=#979a00 gui=bold cterm=bold
hi Error guifg=#ff5e46 gui=NONE cterm=NONE
hi markdownH1 guifg=#ff29f9 gui=bold cterm=bold
hi markdownH2 guifg=#ff29f9 gui=bold cterm=bold
hi markdownH3 guifg=#ff29f9 gui=NONE cterm=NONE
hi markdownH4 guifg=#ff29f9 gui=NONE cterm=NONE
hi markdownH5 guifg=#6a8fff gui=NONE cterm=NONE
hi markdownH6 guifg=#6a8fff gui=NONE cterm=NONE
hi markdownCode guifg=#1cac00 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#1cac00 gui=NONE cterm=NONE
hi markdownLinkText guifg=#6a8fff gui=underline cterm=underline
hi markdownUrl guifg=#5474ff gui=NONE cterm=NONE
hi markdownListMarker guifg=#ff29f9 gui=NONE cterm=NONE
hi markdownRule guifg=#5474ff gui=NONE cterm=NONE
hi markdownBlockquote guifg=#5474ff gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#5474ff gui=NONE cterm=NONE
hi DiagnosticError guifg=#ff5e46 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#979a00 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#6a8fff gui=NONE cterm=NONE
hi DiagnosticHint guifg=#5b9ad2 gui=NONE cterm=NONE
hi DiffAdd guifg=#1cac00 gui=NONE cterm=NONE
hi DiffDelete guifg=#ff5e46 gui=NONE cterm=NONE
hi DiffChange guifg=#979a00 gui=NONE cterm=NONE
hi DiffText guifg=#ff29f9 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#22293e', '#ff5e46', '#1cac00', '#979a00', '#6a8fff', '#ff29f9', '#5b9ad2', '#72bdff',
  \ '#5474ff', '#ff9581', '#24ce00', '#b4b900', '#92b0ff', '#ff84f2', '#6db8fa', '#acd7ff']
