" pluto-night -- a dark colourway, made in paratune from pluto.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/pluto-night.css;
" change the source, not this file.
"
"   :colorscheme pluto-night-inverted
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
let g:colors_name = 'pluto-night-inverted'

hi Normal guifg=#425861 guibg=#f4eecb gui=NONE cterm=NONE
hi NormalFloat guifg=#425861 guibg=#e7e3c3 gui=NONE cterm=NONE
hi FloatBorder guifg=#659fa4 guibg=#e7e3c3 gui=NONE cterm=NONE
hi CursorLine guibg=#e8e4c4 gui=NONE cterm=NONE
hi CursorLineNr guifg=#007300 gui=bold cterm=bold
hi LineNr guifg=#b99755 gui=NONE cterm=NONE
hi SignColumn guibg=#f4eecb gui=NONE cterm=NONE
hi Visual guibg=#f9f9f9 gui=NONE cterm=NONE
hi VertSplit guifg=#659fa4 gui=NONE cterm=NONE
hi WinSeparator guifg=#659fa4 gui=NONE cterm=NONE
hi StatusLine guifg=#425861 guibg=#e7e3c3 gui=NONE cterm=NONE
hi StatusLineNC guifg=#b99755 guibg=#e7e3c3 gui=NONE cterm=NONE
hi Pmenu guifg=#425861 guibg=#e7e3c3 gui=NONE cterm=NONE
hi PmenuSel guifg=#f4eecb guibg=#007300 gui=NONE cterm=NONE
hi Search guifg=#f4eecb guibg=#947cf8 gui=NONE cterm=NONE
hi IncSearch guifg=#f4eecb guibg=#007300 gui=NONE cterm=NONE
hi MatchParen guifg=#c87c00 gui=bold cterm=bold
hi Directory guifg=#c87c00 gui=NONE cterm=NONE
hi Folded guifg=#659fa4 guibg=#e7e3c3 gui=NONE cterm=NONE
hi NonText guifg=#b99755 gui=NONE cterm=NONE
hi Whitespace guifg=#b99755 gui=NONE cterm=NONE
hi Conceal guifg=#b99755 gui=NONE cterm=NONE
hi Title guifg=#007300 gui=bold cterm=bold
hi Comment guifg=#659fa4 gui=italic cterm=italic
hi String guifg=#947cf8 gui=NONE cterm=NONE
hi Character guifg=#947cf8 gui=NONE cterm=NONE
hi Number guifg=#947cf8 gui=NONE cterm=NONE
hi Boolean guifg=#947cf8 gui=NONE cterm=NONE
hi Identifier guifg=#425861 gui=NONE cterm=NONE
hi Function guifg=#c87c00 gui=NONE cterm=NONE
hi Statement guifg=#007300 gui=NONE cterm=NONE
hi Keyword guifg=#007300 gui=NONE cterm=NONE
hi Operator guifg=#659fa4 gui=NONE cterm=NONE
hi PreProc guifg=#8b2b8b gui=NONE cterm=NONE
hi Type guifg=#8b2b8b gui=NONE cterm=NONE
hi Constant guifg=#947cf8 gui=NONE cterm=NONE
hi Special guifg=#c87c00 gui=NONE cterm=NONE
hi Todo guifg=#f4eecb guibg=#947cf8 gui=bold cterm=bold
hi Error guifg=#29f4f2 gui=NONE cterm=NONE
hi markdownH1 guifg=#007300 gui=bold cterm=bold
hi markdownH2 guifg=#007300 gui=bold cterm=bold
hi markdownH3 guifg=#007300 gui=NONE cterm=NONE
hi markdownH4 guifg=#007300 gui=NONE cterm=NONE
hi markdownH5 guifg=#c87c00 gui=NONE cterm=NONE
hi markdownH6 guifg=#c87c00 gui=NONE cterm=NONE
hi markdownCode guifg=#8b2b8b gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#8b2b8b gui=NONE cterm=NONE
hi markdownLinkText guifg=#c87c00 gui=underline cterm=underline
hi markdownUrl guifg=#659fa4 gui=NONE cterm=NONE
hi markdownListMarker guifg=#007300 gui=NONE cterm=NONE
hi markdownRule guifg=#659fa4 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#659fa4 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#b99755 gui=NONE cterm=NONE
hi DiagnosticError guifg=#29f4f2 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#7758f6 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#c87c00 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#a449a4 gui=NONE cterm=NONE
hi DiffAdd guifg=#8b2b8b gui=NONE cterm=NONE
hi DiffDelete guifg=#29f4f2 gui=NONE cterm=NONE
hi DiffChange guifg=#947cf8 gui=NONE cterm=NONE
hi DiffText guifg=#007300 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#e1d2af', '#29f4f2', '#8b2b8b', '#7758f6', '#c87c00', '#007300', '#3a0000', '#007da0',
  \ '#339fba', '#00f6f6', '#c671c6', '#7e7eff', '#df8900', '#36bd36', '#9d7c7c', '#36a3be']
