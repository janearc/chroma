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
" rendered by colourway from sources/vaporwave.css;
" change the source, not this file.
"
"   :colorscheme vaporwave-inverted
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
let g:colors_name = 'vaporwave-inverted'

hi Normal guifg=#323c27 guibg=#e9f2d4 gui=NONE cterm=NONE
hi NormalFloat guifg=#323c27 guibg=#d5e5b2 gui=NONE cterm=NONE
hi FloatBorder guifg=#bcd394 guibg=#d5e5b2 gui=NONE cterm=NONE
hi CursorLine guibg=#ddeabf gui=NONE cterm=NONE
hi CursorLineNr guifg=#005e00 gui=bold cterm=bold
hi LineNr guifg=#a3ba7a gui=NONE cterm=NONE
hi SignColumn guibg=#e9f2d4 gui=NONE cterm=NONE
hi Visual guibg=#c5dd9f gui=NONE cterm=NONE
hi VertSplit guifg=#bcd394 gui=NONE cterm=NONE
hi WinSeparator guifg=#bcd394 gui=NONE cterm=NONE
hi StatusLine guifg=#323c27 guibg=#d5e5b2 gui=NONE cterm=NONE
hi StatusLineNC guifg=#a3ba7a guibg=#d5e5b2 gui=NONE cterm=NONE
hi Pmenu guifg=#323c27 guibg=#d5e5b2 gui=NONE cterm=NONE
hi PmenuSel guifg=#e9f2d4 guibg=#005e00 gui=NONE cterm=NONE
hi Search guifg=#e9f2d4 guibg=#1d4285 gui=NONE cterm=NONE
hi IncSearch guifg=#e9f2d4 guibg=#005e00 gui=NONE cterm=NONE
hi MatchParen guifg=#ff2006 gui=bold cterm=bold
hi Directory guifg=#ff2006 gui=NONE cterm=NONE
hi Folded guifg=#2f5000 guibg=#d5e5b2 gui=NONE cterm=NONE
hi NonText guifg=#a3ba7a gui=NONE cterm=NONE
hi Whitespace guifg=#a3ba7a gui=NONE cterm=NONE
hi Conceal guifg=#a3ba7a gui=NONE cterm=NONE
hi Title guifg=#005e00 gui=bold cterm=bold
hi Comment guifg=#2f5000 gui=italic cterm=italic
hi String guifg=#1d4285 gui=NONE cterm=NONE
hi Character guifg=#1d4285 gui=NONE cterm=NONE
hi Number guifg=#1d4285 gui=NONE cterm=NONE
hi Boolean guifg=#1d4285 gui=NONE cterm=NONE
hi Identifier guifg=#323c27 gui=NONE cterm=NONE
hi Function guifg=#ff2006 gui=NONE cterm=NONE
hi Statement guifg=#005e00 gui=NONE cterm=NONE
hi Keyword guifg=#005e00 gui=NONE cterm=NONE
hi Operator guifg=#2f5000 gui=NONE cterm=NONE
hi PreProc guifg=#64305a gui=NONE cterm=NONE
hi Type guifg=#64305a gui=NONE cterm=NONE
hi Constant guifg=#1d4285 gui=NONE cterm=NONE
hi Special guifg=#ff2006 gui=NONE cterm=NONE
hi Todo guifg=#e9f2d4 guibg=#1d4285 gui=bold cterm=bold
hi Error guifg=#007564 gui=NONE cterm=NONE
hi markdownH1 guifg=#005e00 gui=bold cterm=bold
hi markdownH2 guifg=#005e00 gui=bold cterm=bold
hi markdownH3 guifg=#005e00 gui=NONE cterm=NONE
hi markdownH4 guifg=#005e00 gui=NONE cterm=NONE
hi markdownH5 guifg=#ff2006 gui=NONE cterm=NONE
hi markdownH6 guifg=#ff2006 gui=NONE cterm=NONE
hi markdownCode guifg=#64305a gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#64305a gui=NONE cterm=NONE
hi markdownLinkText guifg=#ff2006 gui=underline cterm=underline
hi markdownUrl guifg=#2f5000 gui=NONE cterm=NONE
hi markdownListMarker guifg=#005e00 gui=NONE cterm=NONE
hi markdownRule guifg=#bcd394 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#2f5000 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#a3ba7a gui=NONE cterm=NONE
hi DiagnosticError guifg=#007564 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#1d4285 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#ff2006 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#64305a gui=NONE cterm=NONE
hi DiffAdd guifg=#64305a gui=NONE cterm=NONE
hi DiffDelete guifg=#007564 gui=NONE cterm=NONE
hi DiffChange guifg=#1d4285 gui=NONE cterm=NONE
hi DiffText guifg=#005e00 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#dbedba', '#005b2b', '#7b2f63', '#1d4684', '#6b3c00', '#006301', '#93311d', '#394426',
  \ '#a3ba7a', '#004711', '#702057', '#0c397b', '#5f2c00', '#004a00', '#8b230d', '#2c3617']
