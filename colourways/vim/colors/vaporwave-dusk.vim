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
" rendered by colourway from sources/vaporwave-dusk.css;
" change the source, not this file.
"
"   :colorscheme vaporwave-dusk
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
let g:colors_name = 'vaporwave-dusk'

hi Normal guifg=#b0a295 guibg=#271d3b gui=NONE cterm=NONE
hi NormalFloat guifg=#b0a295 guibg=#332850 gui=NONE cterm=NONE
hi FloatBorder guifg=#432c6b guibg=#332850 gui=NONE cterm=NONE
hi CursorLine guibg=#30264a gui=NONE cterm=NONE
hi CursorLineNr guifg=#b893b4 gui=bold cterm=bold
hi LineNr guifg=#5c4585 gui=NONE cterm=NONE
hi SignColumn guibg=#271d3b gui=NONE cterm=NONE
hi Visual guibg=#3a2260 gui=NONE cterm=NONE
hi VertSplit guifg=#432c6b gui=NONE cterm=NONE
hi WinSeparator guifg=#432c6b gui=NONE cterm=NONE
hi StatusLine guifg=#b0a295 guibg=#332850 gui=NONE cterm=NONE
hi StatusLineNC guifg=#5c4585 guibg=#332850 gui=NONE cterm=NONE
hi Pmenu guifg=#b0a295 guibg=#332850 gui=NONE cterm=NONE
hi PmenuSel guifg=#271d3b guibg=#b893b4 gui=NONE cterm=NONE
hi Search guifg=#271d3b guibg=#c4a57a gui=NONE cterm=NONE
hi IncSearch guifg=#271d3b guibg=#b893b4 gui=NONE cterm=NONE
hi MatchParen guifg=#84aeb0 gui=bold cterm=bold
hi Directory guifg=#84aeb0 gui=NONE cterm=NONE
hi Folded guifg=#9e9590 guibg=#332850 gui=NONE cterm=NONE
hi NonText guifg=#5c4585 gui=NONE cterm=NONE
hi Whitespace guifg=#5c4585 gui=NONE cterm=NONE
hi Conceal guifg=#5c4585 gui=NONE cterm=NONE
hi Title guifg=#b893b4 gui=bold cterm=bold
hi Comment guifg=#9e9590 gui=italic cterm=italic
hi String guifg=#c4a57a gui=NONE cterm=NONE
hi Character guifg=#c4a57a gui=NONE cterm=NONE
hi Number guifg=#c4a57a gui=NONE cterm=NONE
hi Boolean guifg=#c4a57a gui=NONE cterm=NONE
hi Identifier guifg=#b0a295 gui=NONE cterm=NONE
hi Function guifg=#84aeb0 gui=NONE cterm=NONE
hi Statement guifg=#b893b4 gui=NONE cterm=NONE
hi Keyword guifg=#b893b4 gui=NONE cterm=NONE
hi Operator guifg=#9e9590 gui=NONE cterm=NONE
hi PreProc guifg=#94b08e gui=NONE cterm=NONE
hi Type guifg=#94b08e gui=NONE cterm=NONE
hi Constant guifg=#c4a57a gui=NONE cterm=NONE
hi Special guifg=#84aeb0 gui=NONE cterm=NONE
hi Todo guifg=#271d3b guibg=#c4a57a gui=bold cterm=bold
hi Error guifg=#c9958f gui=NONE cterm=NONE
hi markdownH1 guifg=#b893b4 gui=bold cterm=bold
hi markdownH2 guifg=#b893b4 gui=bold cterm=bold
hi markdownH3 guifg=#b893b4 gui=NONE cterm=NONE
hi markdownH4 guifg=#b893b4 gui=NONE cterm=NONE
hi markdownH5 guifg=#84aeb0 gui=NONE cterm=NONE
hi markdownH6 guifg=#84aeb0 gui=NONE cterm=NONE
hi markdownCode guifg=#94b08e gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#94b08e gui=NONE cterm=NONE
hi markdownLinkText guifg=#84aeb0 gui=underline cterm=underline
hi markdownUrl guifg=#9e9590 gui=NONE cterm=NONE
hi markdownListMarker guifg=#b893b4 gui=NONE cterm=NONE
hi markdownRule guifg=#432c6b gui=NONE cterm=NONE
hi markdownBlockquote guifg=#9e9590 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#5c4585 gui=NONE cterm=NONE
hi DiagnosticError guifg=#c9958f gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#c4a57a gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#84aeb0 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#94b08e gui=NONE cterm=NONE
hi DiffAdd guifg=#94b08e gui=NONE cterm=NONE
hi DiffDelete guifg=#c9958f gui=NONE cterm=NONE
hi DiffChange guifg=#c4a57a gui=NONE cterm=NONE
hi DiffText guifg=#b893b4 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#291f3e', '#c9958f', '#8f9580', '#c4a57a', '#8a93a6', '#b893b4', '#84aeb0', '#a3988e',
  \ '#7f746b', '#d1a09a', '#979d88', '#ccae84', '#939caf', '#c19dbd', '#8eb8ba', '#ab9d90']
