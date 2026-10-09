" vaporwave -- matched to ~/.config/nvim/colors/vaporwave.lua
"
" Every text colour is placed against a measured comfortable band rather than
" picked. Contrast has a ceiling as well as a floor: past it, light text on a
" dark ground halates and the strokes smear. Plain white here is 18.6:1, which
" is far past it. Nothing below exceeds about 12:1.
"
" The band is narrow, so the normal and bright halves of the palette cannot
" separate by brightness alone -- they would either be indistinguishable or
" push bright back through the ceiling. They separate by saturation instead,
" with only a small step in lightness.
"
" rendered by colourway from sources/vaporwave.css;
" change the source, not this file.
"
"   :colorscheme vaporwave
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
let g:colors_name = 'vaporwave'

hi Normal guifg=#cdc3d8 guibg=#160d2b gui=NONE cterm=NONE
hi NormalFloat guifg=#cdc3d8 guibg=#2a1a4d gui=NONE cterm=NONE
hi FloatBorder guifg=#432c6b guibg=#2a1a4d gui=NONE cterm=NONE
hi CursorLine guibg=#221540 gui=NONE cterm=NONE
hi CursorLineNr guifg=#ffa1ff gui=bold cterm=bold
hi LineNr guifg=#5c4585 gui=NONE cterm=NONE
hi SignColumn guibg=#160d2b gui=NONE cterm=NONE
hi Visual guibg=#3a2260 gui=NONE cterm=NONE
hi VertSplit guifg=#432c6b gui=NONE cterm=NONE
hi WinSeparator guifg=#432c6b gui=NONE cterm=NONE
hi StatusLine guifg=#cdc3d8 guibg=#2a1a4d gui=NONE cterm=NONE
hi StatusLineNC guifg=#5c4585 guibg=#2a1a4d gui=NONE cterm=NONE
hi Pmenu guifg=#cdc3d8 guibg=#2a1a4d gui=NONE cterm=NONE
hi PmenuSel guifg=#160d2b guibg=#ffa1ff gui=NONE cterm=NONE
hi Search guifg=#160d2b guibg=#e2bd7a gui=NONE cterm=NONE
hi IncSearch guifg=#160d2b guibg=#ffa1ff gui=NONE cterm=NONE
hi MatchParen guifg=#00dff9 gui=bold cterm=bold
hi Directory guifg=#00dff9 gui=NONE cterm=NONE
hi Folded guifg=#d0afff guibg=#2a1a4d gui=NONE cterm=NONE
hi NonText guifg=#5c4585 gui=NONE cterm=NONE
hi Whitespace guifg=#5c4585 gui=NONE cterm=NONE
hi Conceal guifg=#5c4585 gui=NONE cterm=NONE
hi Title guifg=#ffa1ff gui=bold cterm=bold
hi Comment guifg=#d0afff gui=italic cterm=italic
hi String guifg=#e2bd7a gui=NONE cterm=NONE
hi Character guifg=#e2bd7a gui=NONE cterm=NONE
hi Number guifg=#e2bd7a gui=NONE cterm=NONE
hi Boolean guifg=#e2bd7a gui=NONE cterm=NONE
hi Identifier guifg=#cdc3d8 gui=NONE cterm=NONE
hi Function guifg=#00dff9 gui=NONE cterm=NONE
hi Statement guifg=#ffa1ff gui=NONE cterm=NONE
hi Keyword guifg=#ffa1ff gui=NONE cterm=NONE
hi Operator guifg=#d0afff gui=NONE cterm=NONE
hi PreProc guifg=#9bcfa5 gui=NONE cterm=NONE
hi Type guifg=#9bcfa5 gui=NONE cterm=NONE
hi Constant guifg=#e2bd7a gui=NONE cterm=NONE
hi Special guifg=#00dff9 gui=NONE cterm=NONE
hi Todo guifg=#160d2b guibg=#e2bd7a gui=bold cterm=bold
hi Error guifg=#ff8a9b gui=NONE cterm=NONE
hi markdownH1 guifg=#ffa1ff gui=bold cterm=bold
hi markdownH2 guifg=#ffa1ff gui=bold cterm=bold
hi markdownH3 guifg=#ffa1ff gui=NONE cterm=NONE
hi markdownH4 guifg=#ffa1ff gui=NONE cterm=NONE
hi markdownH5 guifg=#00dff9 gui=NONE cterm=NONE
hi markdownH6 guifg=#00dff9 gui=NONE cterm=NONE
hi markdownCode guifg=#9bcfa5 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#9bcfa5 gui=NONE cterm=NONE
hi markdownLinkText guifg=#00dff9 gui=underline cterm=underline
hi markdownUrl guifg=#d0afff gui=NONE cterm=NONE
hi markdownListMarker guifg=#ffa1ff gui=NONE cterm=NONE
hi markdownRule guifg=#432c6b gui=NONE cterm=NONE
hi markdownBlockquote guifg=#d0afff gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#5c4585 gui=NONE cterm=NONE
hi DiagnosticError guifg=#ff8a9b gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#e2bd7a gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#00dff9 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#9bcfa5 gui=NONE cterm=NONE
hi DiffAdd guifg=#9bcfa5 gui=NONE cterm=NONE
hi DiffDelete guifg=#ff8a9b gui=NONE cterm=NONE
hi DiffChange guifg=#e2bd7a gui=NONE cterm=NONE
hi DiffText guifg=#ffa1ff gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#241245', '#ffa4d4', '#84d09c', '#e2b97b', '#94c3ff', '#ff9cfe', '#6ccee2', '#c6bbd9',
  \ '#5c4585', '#ffb8ee', '#8fdfa8', '#f3c684', '#a0d3ff', '#ffb5ff', '#74dcf2', '#d3c9e8']
