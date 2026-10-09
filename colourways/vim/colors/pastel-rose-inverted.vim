" pastel-rose -- made in paratune from pastel-rose.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/pastel-rose.css;
" change the source, not this file.
"
"   :colorscheme pastel-rose-inverted
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
let g:colors_name = 'pastel-rose-inverted'

hi Normal guifg=#fdf8f6 guibg=#0f7d0f gui=NONE cterm=NONE
hi NormalFloat guifg=#fdf8f6 guibg=#2c8628 gui=NONE cterm=NONE
hi FloatBorder guifg=#f7f9f6 guibg=#2c8628 gui=NONE cterm=NONE
hi CursorLine guibg=#298425 gui=NONE cterm=NONE
hi CursorLineNr guifg=#8620ff gui=bold cterm=bold
hi LineNr guifg=#f7f9f6 gui=NONE cterm=NONE
hi SignColumn guibg=#0f7d0f gui=NONE cterm=NONE
hi Visual guibg=#04a71c gui=NONE cterm=NONE
hi VertSplit guifg=#f7f9f6 gui=NONE cterm=NONE
hi WinSeparator guifg=#f7f9f6 gui=NONE cterm=NONE
hi StatusLine guifg=#fdf8f6 guibg=#2c8628 gui=NONE cterm=NONE
hi StatusLineNC guifg=#f7f9f6 guibg=#2c8628 gui=NONE cterm=NONE
hi Pmenu guifg=#fdf8f6 guibg=#2c8628 gui=NONE cterm=NONE
hi PmenuSel guifg=#0f7d0f guibg=#8620ff gui=NONE cterm=NONE
hi Search guifg=#0f7d0f guibg=#fff6fa gui=NONE cterm=NONE
hi IncSearch guifg=#0f7d0f guibg=#8620ff gui=NONE cterm=NONE
hi MatchParen guifg=#ffffe5 gui=bold cterm=bold
hi Directory guifg=#ffffe5 gui=NONE cterm=NONE
hi Folded guifg=#f7f9f6 guibg=#2c8628 gui=NONE cterm=NONE
hi NonText guifg=#f7f9f6 gui=NONE cterm=NONE
hi Whitespace guifg=#f7f9f6 gui=NONE cterm=NONE
hi Conceal guifg=#f7f9f6 gui=NONE cterm=NONE
hi Title guifg=#8620ff gui=bold cterm=bold
hi Comment guifg=#f7f9f6 gui=italic cterm=italic
hi String guifg=#fff6fa gui=NONE cterm=NONE
hi Character guifg=#fff6fa gui=NONE cterm=NONE
hi Number guifg=#fff6fa gui=NONE cterm=NONE
hi Boolean guifg=#fff6fa gui=NONE cterm=NONE
hi Identifier guifg=#fdf8f6 gui=NONE cterm=NONE
hi Function guifg=#ffffe5 gui=NONE cterm=NONE
hi Statement guifg=#8620ff gui=NONE cterm=NONE
hi Keyword guifg=#8620ff gui=NONE cterm=NONE
hi Operator guifg=#f7f9f6 gui=NONE cterm=NONE
hi PreProc guifg=#000000 gui=NONE cterm=NONE
hi Type guifg=#000000 gui=NONE cterm=NONE
hi Constant guifg=#fff6fa gui=NONE cterm=NONE
hi Special guifg=#ffffe5 gui=NONE cterm=NONE
hi Todo guifg=#0f7d0f guibg=#fff6fa gui=bold cterm=bold
hi Error guifg=#c32222 gui=NONE cterm=NONE
hi markdownH1 guifg=#8620ff gui=bold cterm=bold
hi markdownH2 guifg=#8620ff gui=bold cterm=bold
hi markdownH3 guifg=#8620ff gui=NONE cterm=NONE
hi markdownH4 guifg=#8620ff gui=NONE cterm=NONE
hi markdownH5 guifg=#ffffe5 gui=NONE cterm=NONE
hi markdownH6 guifg=#ffffe5 gui=NONE cterm=NONE
hi markdownCode guifg=#000000 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#000000 gui=NONE cterm=NONE
hi markdownLinkText guifg=#ffffe5 gui=underline cterm=underline
hi markdownUrl guifg=#f7f9f6 gui=NONE cterm=NONE
hi markdownListMarker guifg=#8620ff gui=NONE cterm=NONE
hi markdownRule guifg=#f7f9f6 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#f7f9f6 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#f7f9f6 gui=NONE cterm=NONE
hi DiagnosticError guifg=#c32222 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#0f1de8 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#ffffe5 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#df0000 gui=NONE cterm=NONE
hi DiffAdd guifg=#000000 gui=NONE cterm=NONE
hi DiffDelete guifg=#c32222 gui=NONE cterm=NONE
hi DiffChange guifg=#fff6fa gui=NONE cterm=NONE
hi DiffText guifg=#8620ff gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#1a8800', '#c32222', '#000000', '#0000e2', '#ffffe5', '#7603ff', '#df0000', '#000000',
  \ '#f7f9f6', '#5fdab4', '#f5955f', '#0000e2', '#f7fdf1', '#ff00f4', '#df0000', '#000000']
