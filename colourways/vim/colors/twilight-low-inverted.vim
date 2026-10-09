" twilight-low -- made in paratune from twilight.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/twilight-low.css;
" change the source, not this file.
"
"   :colorscheme twilight-low-inverted
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
let g:colors_name = 'twilight-low-inverted'

hi Normal guifg=#392100 guibg=#14c0ed gui=NONE cterm=NONE
hi NormalFloat guifg=#392100 guibg=#002c4e gui=NONE cterm=NONE
hi FloatBorder guifg=#000000 guibg=#002c4e gui=NONE cterm=NONE
hi CursorLine guibg=#00315d gui=NONE cterm=NONE
hi CursorLineNr guifg=#00a22c gui=bold cterm=bold
hi LineNr guifg=#000000 gui=NONE cterm=NONE
hi SignColumn guibg=#14c0ed gui=NONE cterm=NONE
hi Visual guibg=#0085e5 gui=NONE cterm=NONE
hi VertSplit guifg=#000000 gui=NONE cterm=NONE
hi WinSeparator guifg=#000000 gui=NONE cterm=NONE
hi StatusLine guifg=#392100 guibg=#002c4e gui=NONE cterm=NONE
hi StatusLineNC guifg=#000000 guibg=#002c4e gui=NONE cterm=NONE
hi Pmenu guifg=#392100 guibg=#002c4e gui=NONE cterm=NONE
hi PmenuSel guifg=#14c0ed guibg=#00a22c gui=NONE cterm=NONE
hi Search guifg=#14c0ed guibg=#2551c4 gui=NONE cterm=NONE
hi IncSearch guifg=#14c0ed guibg=#00a22c gui=NONE cterm=NONE
hi MatchParen guifg=#835c00 gui=bold cterm=bold
hi Directory guifg=#835c00 gui=NONE cterm=NONE
hi Folded guifg=#2b5dc6 guibg=#002c4e gui=NONE cterm=NONE
hi NonText guifg=#000000 gui=NONE cterm=NONE
hi Whitespace guifg=#000000 gui=NONE cterm=NONE
hi Conceal guifg=#000000 gui=NONE cterm=NONE
hi Title guifg=#00a22c gui=bold cterm=bold
hi Comment guifg=#2b5dc6 gui=italic cterm=italic
hi String guifg=#2551c4 gui=NONE cterm=NONE
hi Character guifg=#2551c4 gui=NONE cterm=NONE
hi Number guifg=#2551c4 gui=NONE cterm=NONE
hi Boolean guifg=#2551c4 gui=NONE cterm=NONE
hi Identifier guifg=#392100 gui=NONE cterm=NONE
hi Function guifg=#835c00 gui=NONE cterm=NONE
hi Statement guifg=#00a22c gui=NONE cterm=NONE
hi Keyword guifg=#00a22c gui=NONE cterm=NONE
hi Operator guifg=#2b5dc6 gui=NONE cterm=NONE
hi PreProc guifg=#804cc8 gui=NONE cterm=NONE
hi Type guifg=#804cc8 gui=NONE cterm=NONE
hi Constant guifg=#2551c4 gui=NONE cterm=NONE
hi Special guifg=#835c00 gui=NONE cterm=NONE
hi Todo guifg=#14c0ed guibg=#2551c4 gui=bold cterm=bold
hi Error guifg=#00747a gui=NONE cterm=NONE
hi markdownH1 guifg=#00a22c gui=bold cterm=bold
hi markdownH2 guifg=#00a22c gui=bold cterm=bold
hi markdownH3 guifg=#00a22c gui=NONE cterm=NONE
hi markdownH4 guifg=#00a22c gui=NONE cterm=NONE
hi markdownH5 guifg=#835c00 gui=NONE cterm=NONE
hi markdownH6 guifg=#835c00 gui=NONE cterm=NONE
hi markdownCode guifg=#804cc8 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#804cc8 gui=NONE cterm=NONE
hi markdownLinkText guifg=#835c00 gui=underline cterm=underline
hi markdownUrl guifg=#2b5dc6 gui=NONE cterm=NONE
hi markdownListMarker guifg=#00a22c gui=NONE cterm=NONE
hi markdownRule guifg=#000000 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#2b5dc6 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#000000 gui=NONE cterm=NONE
hi DiagnosticError guifg=#00747a gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#2551c4 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#835c00 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#804cc8 gui=NONE cterm=NONE
hi DiffAdd guifg=#804cc8 gui=NONE cterm=NONE
hi DiffDelete guifg=#00747a gui=NONE cterm=NONE
hi DiffChange guifg=#2551c4 gui=NONE cterm=NONE
hi DiffText guifg=#00a22c gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#ef9286', '#000000', '#2d0055', '#000c2c', '#110b00', '#000000', '#440200', '#000305',
  \ '#5c5c51', '#000000', '#2e005e', '#000c2a', '#0f0a00', '#000000', '#4d0001', '#000000']
