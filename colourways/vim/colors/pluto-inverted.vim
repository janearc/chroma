" pluto -- a light colourway, pluto as new horizons first showed it on
" 13 july 2015, the day before closest approach (nasa/jhuapl/swri,
" pia19708). the ground is the heart; the ink is the whale, darkened to 10
" to 1; the selection is the north polar cap, claude's band the midlands,
" dim text the disc. pluto has no red, green or blue to speak of, so the
" sixteen are the screen's own wavelengths, darkened to stand on the heart.
" samples in cmd/spectra/data/pluto-2015.css. made by spectra,
" github.com/janearc/chroma/cmd/spectra, 2026-10-05.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/pluto.css;
" change the source, not this file.
"
"   :colorscheme pluto-inverted
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
let g:colors_name = 'pluto-inverted'

hi Normal guifg=#d5e4ef guibg=#163d62 gui=NONE cterm=NONE
hi NormalFloat guifg=#d5e4ef guibg=#23486b gui=NONE cterm=NONE
hi FloatBorder guifg=#9cafbf guibg=#23486b gui=NONE cterm=NONE
hi CursorLine guibg=#21466a gui=NONE cterm=NONE
hi CursorLineNr guifg=#85ff86 gui=bold cterm=bold
hi LineNr guifg=#9cafbf gui=NONE cterm=NONE
hi SignColumn guibg=#163d62 gui=NONE cterm=NONE
hi Visual guibg=#295177 gui=NONE cterm=NONE
hi VertSplit guifg=#9cafbf gui=NONE cterm=NONE
hi WinSeparator guifg=#9cafbf gui=NONE cterm=NONE
hi StatusLine guifg=#d5e4ef guibg=#23486b gui=NONE cterm=NONE
hi StatusLineNC guifg=#9cafbf guibg=#23486b gui=NONE cterm=NONE
hi Pmenu guifg=#d5e4ef guibg=#23486b gui=NONE cterm=NONE
hi PmenuSel guifg=#163d62 guibg=#85ff86 gui=NONE cterm=NONE
hi Search guifg=#163d62 guibg=#bbbaff gui=NONE cterm=NONE
hi IncSearch guifg=#163d62 guibg=#85ff86 gui=NONE cterm=NONE
hi MatchParen guifg=#ffbd7d gui=bold cterm=bold
hi Directory guifg=#ffbd7d gui=NONE cterm=NONE
hi Folded guifg=#9cafbf guibg=#23486b gui=NONE cterm=NONE
hi NonText guifg=#9cafbf gui=NONE cterm=NONE
hi Whitespace guifg=#9cafbf gui=NONE cterm=NONE
hi Conceal guifg=#9cafbf gui=NONE cterm=NONE
hi Title guifg=#85ff86 gui=bold cterm=bold
hi Comment guifg=#9cafbf gui=italic cterm=italic
hi String guifg=#bbbaff gui=NONE cterm=NONE
hi Character guifg=#bbbaff gui=NONE cterm=NONE
hi Number guifg=#bbbaff gui=NONE cterm=NONE
hi Boolean guifg=#bbbaff gui=NONE cterm=NONE
hi Identifier guifg=#d5e4ef gui=NONE cterm=NONE
hi Function guifg=#ffbd7d gui=NONE cterm=NONE
hi Statement guifg=#85ff86 gui=NONE cterm=NONE
hi Keyword guifg=#85ff86 gui=NONE cterm=NONE
hi Operator guifg=#9cafbf gui=NONE cterm=NONE
hi PreProc guifg=#f8b1ff gui=NONE cterm=NONE
hi Type guifg=#f8b1ff gui=NONE cterm=NONE
hi Constant guifg=#bbbaff gui=NONE cterm=NONE
hi Special guifg=#ffbd7d gui=NONE cterm=NONE
hi Todo guifg=#163d62 guibg=#bbbaff gui=bold cterm=bold
hi Error guifg=#77f1ff gui=NONE cterm=NONE
hi markdownH1 guifg=#85ff86 gui=bold cterm=bold
hi markdownH2 guifg=#85ff86 gui=bold cterm=bold
hi markdownH3 guifg=#85ff86 gui=NONE cterm=NONE
hi markdownH4 guifg=#85ff86 gui=NONE cterm=NONE
hi markdownH5 guifg=#ffbd7d gui=NONE cterm=NONE
hi markdownH6 guifg=#ffbd7d gui=NONE cterm=NONE
hi markdownCode guifg=#f8b1ff gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#f8b1ff gui=NONE cterm=NONE
hi markdownLinkText guifg=#ffbd7d gui=underline cterm=underline
hi markdownUrl guifg=#9cafbf gui=NONE cterm=NONE
hi markdownListMarker guifg=#85ff86 gui=NONE cterm=NONE
hi markdownRule guifg=#9cafbf gui=NONE cterm=NONE
hi markdownBlockquote guifg=#9cafbf gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#9cafbf gui=NONE cterm=NONE
hi DiagnosticError guifg=#77f1ff gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#bbbaff gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#ffbd7d gui=NONE cterm=NONE
hi DiagnosticHint guifg=#ffb4b6 gui=NONE cterm=NONE
hi DiffAdd guifg=#f8b1ff gui=NONE cterm=NONE
hi DiffDelete guifg=#77f1ff gui=NONE cterm=NONE
hi DiffChange guifg=#bbbaff gui=NONE cterm=NONE
hi DiffText guifg=#85ff86 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#2a4965', '#77f1ff', '#f8b1ff', '#bbbaff', '#ffbd7d', '#85ff86', '#ffb4b6', '#d0dae2',
  \ '#9cafbf', '#a3f9ff', '#fcccff', '#d3d2ff', '#ffd4a6', '#acffad', '#ffced0', '#f0f8fc']
