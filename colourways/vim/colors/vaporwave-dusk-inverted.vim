" vaporwave-dusk -- a small step from vaporwave: the ground a little
" lighter; the ink warm, dim, about 6:1; every colour muted, warmer and
" under 7:1, so nothing smears.
"
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
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/vaporwave-dusk.css;
" change the source, not this file.
"
"   :colorscheme vaporwave-dusk-inverted
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
let g:colors_name = 'vaporwave-dusk-inverted'

hi Normal guifg=#4f5d6a guibg=#d8e2c4 gui=NONE cterm=NONE
hi NormalFloat guifg=#4f5d6a guibg=#ccd7af gui=NONE cterm=NONE
hi FloatBorder guifg=#bcd394 guibg=#ccd7af gui=NONE cterm=NONE
hi CursorLine guibg=#cfd9b5 gui=NONE cterm=NONE
hi CursorLineNr guifg=#476c4b gui=bold cterm=bold
hi LineNr guifg=#a3ba7a gui=NONE cterm=NONE
hi SignColumn guibg=#d8e2c4 gui=NONE cterm=NONE
hi Visual guibg=#c5dd9f gui=NONE cterm=NONE
hi VertSplit guifg=#bcd394 gui=NONE cterm=NONE
hi WinSeparator guifg=#bcd394 gui=NONE cterm=NONE
hi StatusLine guifg=#4f5d6a guibg=#ccd7af gui=NONE cterm=NONE
hi StatusLineNC guifg=#a3ba7a guibg=#ccd7af gui=NONE cterm=NONE
hi Pmenu guifg=#4f5d6a guibg=#ccd7af gui=NONE cterm=NONE
hi PmenuSel guifg=#d8e2c4 guibg=#476c4b gui=NONE cterm=NONE
hi Search guifg=#d8e2c4 guibg=#3b5a85 gui=NONE cterm=NONE
hi IncSearch guifg=#d8e2c4 guibg=#476c4b gui=NONE cterm=NONE
hi MatchParen guifg=#7b514f gui=bold cterm=bold
hi Directory guifg=#7b514f gui=NONE cterm=NONE
hi Folded guifg=#616a6f guibg=#ccd7af gui=NONE cterm=NONE
hi NonText guifg=#a3ba7a gui=NONE cterm=NONE
hi Whitespace guifg=#a3ba7a gui=NONE cterm=NONE
hi Conceal guifg=#a3ba7a gui=NONE cterm=NONE
hi Title guifg=#476c4b gui=bold cterm=bold
hi Comment guifg=#616a6f gui=italic cterm=italic
hi String guifg=#3b5a85 gui=NONE cterm=NONE
hi Character guifg=#3b5a85 gui=NONE cterm=NONE
hi Number guifg=#3b5a85 gui=NONE cterm=NONE
hi Boolean guifg=#3b5a85 gui=NONE cterm=NONE
hi Identifier guifg=#4f5d6a gui=NONE cterm=NONE
hi Function guifg=#7b514f gui=NONE cterm=NONE
hi Statement guifg=#476c4b gui=NONE cterm=NONE
hi Keyword guifg=#476c4b gui=NONE cterm=NONE
hi Operator guifg=#616a6f gui=NONE cterm=NONE
hi PreProc guifg=#6b4f71 gui=NONE cterm=NONE
hi Type guifg=#6b4f71 gui=NONE cterm=NONE
hi Constant guifg=#3b5a85 gui=NONE cterm=NONE
hi Special guifg=#7b514f gui=NONE cterm=NONE
hi Todo guifg=#d8e2c4 guibg=#3b5a85 gui=bold cterm=bold
hi Error guifg=#366a70 gui=NONE cterm=NONE
hi markdownH1 guifg=#476c4b gui=bold cterm=bold
hi markdownH2 guifg=#476c4b gui=bold cterm=bold
hi markdownH3 guifg=#476c4b gui=NONE cterm=NONE
hi markdownH4 guifg=#476c4b gui=NONE cterm=NONE
hi markdownH5 guifg=#7b514f gui=NONE cterm=NONE
hi markdownH6 guifg=#7b514f gui=NONE cterm=NONE
hi markdownCode guifg=#6b4f71 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#6b4f71 gui=NONE cterm=NONE
hi markdownLinkText guifg=#7b514f gui=underline cterm=underline
hi markdownUrl guifg=#616a6f gui=NONE cterm=NONE
hi markdownListMarker guifg=#476c4b gui=NONE cterm=NONE
hi markdownRule guifg=#bcd394 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#616a6f gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#a3ba7a gui=NONE cterm=NONE
hi DiagnosticError guifg=#366a70 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#3b5a85 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#7b514f gui=NONE cterm=NONE
hi DiagnosticHint guifg=#6b4f71 gui=NONE cterm=NONE
hi DiffAdd guifg=#6b4f71 gui=NONE cterm=NONE
hi DiffDelete guifg=#366a70 gui=NONE cterm=NONE
hi DiffChange guifg=#3b5a85 gui=NONE cterm=NONE
hi DiffText guifg=#476c4b gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#d6e0c1', '#366a70', '#706a7f', '#3b5a85', '#756c59', '#476c4b', '#7b514f', '#5c6771',
  \ '#808b94', '#2e5f65', '#686277', '#33517b', '#6c6350', '#3e6242', '#714745', '#54626f']
