" twilight-burnt -- burnt green ink on a darker orange, twilight as it began
" on 2026-10-01, before the ground went bright; with claude code's colours
" dark on the orange.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/twilight-burnt.css;
" change the source, not this file.
"
"   :colorscheme twilight-burnt-inverted
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
let g:colors_name = 'twilight-burnt-inverted'

hi Normal guifg=#e3ddef guibg=#4c80ac gui=NONE cterm=NONE
hi NormalFloat guifg=#e3ddef guibg=#5786b1 gui=NONE cterm=NONE
hi FloatBorder guifg=#b5c2d4 guibg=#5786b1 gui=NONE cterm=NONE
hi CursorLine guibg=#5586b0 gui=NONE cterm=NONE
hi CursorLineNr guifg=#a5d5b5 gui=bold cterm=bold
hi LineNr guifg=#b5c2d4 gui=NONE cterm=NONE
hi SignColumn guibg=#4c80ac gui=NONE cterm=NONE
hi Visual guibg=#2b5f8f gui=NONE cterm=NONE
hi VertSplit guifg=#b5c2d4 gui=NONE cterm=NONE
hi WinSeparator guifg=#b5c2d4 gui=NONE cterm=NONE
hi StatusLine guifg=#e3ddef guibg=#5786b1 gui=NONE cterm=NONE
hi StatusLineNC guifg=#b5c2d4 guibg=#5786b1 gui=NONE cterm=NONE
hi Pmenu guifg=#e3ddef guibg=#5786b1 gui=NONE cterm=NONE
hi PmenuSel guifg=#4c80ac guibg=#a5d5b5 gui=NONE cterm=NONE
hi Search guifg=#4c80ac guibg=#a5bbef gui=NONE cterm=NONE
hi IncSearch guifg=#4c80ac guibg=#a5d5b5 gui=NONE cterm=NONE
hi MatchParen guifg=#dbc9a5 gui=bold cterm=bold
hi Directory guifg=#dbc9a5 gui=NONE cterm=NONE
hi Folded guifg=#b5c2d4 guibg=#5786b1 gui=NONE cterm=NONE
hi NonText guifg=#b5c2d4 gui=NONE cterm=NONE
hi Whitespace guifg=#b5c2d4 gui=NONE cterm=NONE
hi Conceal guifg=#b5c2d4 gui=NONE cterm=NONE
hi Title guifg=#a5d5b5 gui=bold cterm=bold
hi Comment guifg=#b5c2d4 gui=italic cterm=italic
hi String guifg=#a5bbef gui=NONE cterm=NONE
hi Character guifg=#a5bbef gui=NONE cterm=NONE
hi Number guifg=#a5bbef gui=NONE cterm=NONE
hi Boolean guifg=#a5bbef gui=NONE cterm=NONE
hi Identifier guifg=#e3ddef gui=NONE cterm=NONE
hi Function guifg=#dbc9a5 gui=NONE cterm=NONE
hi Statement guifg=#a5d5b5 gui=NONE cterm=NONE
hi Keyword guifg=#a5d5b5 gui=NONE cterm=NONE
hi Operator guifg=#b5c2d4 gui=NONE cterm=NONE
hi PreProc guifg=#d0c5e0 gui=NONE cterm=NONE
hi Type guifg=#d0c5e0 gui=NONE cterm=NONE
hi Constant guifg=#a5bbef gui=NONE cterm=NONE
hi Special guifg=#dbc9a5 gui=NONE cterm=NONE
hi Todo guifg=#4c80ac guibg=#a5bbef gui=bold cterm=bold
hi Error guifg=#91d5e0 gui=NONE cterm=NONE
hi markdownH1 guifg=#a5d5b5 gui=bold cterm=bold
hi markdownH2 guifg=#a5d5b5 gui=bold cterm=bold
hi markdownH3 guifg=#a5d5b5 gui=NONE cterm=NONE
hi markdownH4 guifg=#a5d5b5 gui=NONE cterm=NONE
hi markdownH5 guifg=#dbc9a5 gui=NONE cterm=NONE
hi markdownH6 guifg=#dbc9a5 gui=NONE cterm=NONE
hi markdownCode guifg=#d0c5e0 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#d0c5e0 gui=NONE cterm=NONE
hi markdownLinkText guifg=#dbc9a5 gui=underline cterm=underline
hi markdownUrl guifg=#b5c2d4 gui=NONE cterm=NONE
hi markdownListMarker guifg=#a5d5b5 gui=NONE cterm=NONE
hi markdownRule guifg=#b5c2d4 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#b5c2d4 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#b5c2d4 gui=NONE cterm=NONE
hi DiagnosticError guifg=#91d5e0 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#a5bbef gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#dbc9a5 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#e0b5b5 gui=NONE cterm=NONE
hi DiffAdd guifg=#d0c5e0 gui=NONE cterm=NONE
hi DiffDelete guifg=#91d5e0 gui=NONE cterm=NONE
hi DiffChange guifg=#a5bbef gui=NONE cterm=NONE
hi DiffText guifg=#a5d5b5 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#d5dbe7', '#91d5e0', '#d0c5e0', '#a5bbef', '#dbc9a5', '#a5d5b5', '#e0b5b5', '#4f85b1',
  \ '#b5c2d4', '#7bc9d5', '#c4b7d9', '#91abe7', '#d0bb8f', '#91cba6', '#d7a3a3', '#336895']
