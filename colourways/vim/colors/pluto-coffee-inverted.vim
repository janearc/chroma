" pluto-coffee -- made in paratune from pluto-night.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/pluto-coffee.css;
" change the source, not this file.
"
"   :colorscheme pluto-coffee-inverted
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
let g:colors_name = 'pluto-coffee-inverted'

hi Normal guifg=#988e83 guibg=#f2f2ef gui=NONE cterm=NONE
hi NormalFloat guifg=#988e83 guibg=#ebebe7 gui=NONE cterm=NONE
hi FloatBorder guifg=#c7d4d6 guibg=#ebebe7 gui=NONE cterm=NONE
hi CursorLine guibg=#ecece8 gui=NONE cterm=NONE
hi CursorLineNr guifg=#9fae9f gui=bold cterm=bold
hi LineNr guifg=#dcd4c5 gui=NONE cterm=NONE
hi SignColumn guibg=#f2f2ef gui=NONE cterm=NONE
hi Visual guibg=#c7c7c7 gui=NONE cterm=NONE
hi VertSplit guifg=#c7d4d6 gui=NONE cterm=NONE
hi WinSeparator guifg=#c7d4d6 gui=NONE cterm=NONE
hi StatusLine guifg=#988e83 guibg=#ebebe7 gui=NONE cterm=NONE
hi StatusLineNC guifg=#dcd4c5 guibg=#ebebe7 gui=NONE cterm=NONE
hi Pmenu guifg=#988e83 guibg=#ebebe7 gui=NONE cterm=NONE
hi PmenuSel guifg=#f2f2ef guibg=#9fae9f gui=NONE cterm=NONE
hi Search guifg=#f2f2ef guibg=#cfc1d7 gui=NONE cterm=NONE
hi IncSearch guifg=#f2f2ef guibg=#9fae9f gui=NONE cterm=NONE
hi MatchParen guifg=#b8b09e gui=bold cterm=bold
hi Directory guifg=#b8b09e gui=NONE cterm=NONE
hi Folded guifg=#c7d4d6 guibg=#ebebe7 gui=NONE cterm=NONE
hi NonText guifg=#dcd4c5 gui=NONE cterm=NONE
hi Whitespace guifg=#dcd4c5 gui=NONE cterm=NONE
hi Conceal guifg=#dcd4c5 gui=NONE cterm=NONE
hi Title guifg=#9fae9f gui=bold cterm=bold
hi Comment guifg=#c7d4d6 gui=italic cterm=italic
hi String guifg=#cfc1d7 gui=NONE cterm=NONE
hi Character guifg=#cfc1d7 gui=NONE cterm=NONE
hi Number guifg=#cfc1d7 gui=NONE cterm=NONE
hi Boolean guifg=#cfc1d7 gui=NONE cterm=NONE
hi Identifier guifg=#988e83 gui=NONE cterm=NONE
hi Function guifg=#b8b09e gui=NONE cterm=NONE
hi Statement guifg=#9fae9f gui=NONE cterm=NONE
hi Keyword guifg=#9fae9f gui=NONE cterm=NONE
hi Operator guifg=#c7d4d6 gui=NONE cterm=NONE
hi PreProc guifg=#b7a2b8 gui=NONE cterm=NONE
hi Type guifg=#b7a2b8 gui=NONE cterm=NONE
hi Constant guifg=#cfc1d7 gui=NONE cterm=NONE
hi Special guifg=#b8b09e gui=NONE cterm=NONE
hi Todo guifg=#f2f2ef guibg=#cfc1d7 gui=bold cterm=bold
hi Error guifg=#94b7bd gui=NONE cterm=NONE
hi markdownH1 guifg=#9fae9f gui=bold cterm=bold
hi markdownH2 guifg=#9fae9f gui=bold cterm=bold
hi markdownH3 guifg=#9fae9f gui=NONE cterm=NONE
hi markdownH4 guifg=#9fae9f gui=NONE cterm=NONE
hi markdownH5 guifg=#b8b09e gui=NONE cterm=NONE
hi markdownH6 guifg=#b8b09e gui=NONE cterm=NONE
hi markdownCode guifg=#b7a2b8 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#b7a2b8 gui=NONE cterm=NONE
hi markdownLinkText guifg=#b8b09e gui=underline cterm=underline
hi markdownUrl guifg=#c7d4d6 gui=NONE cterm=NONE
hi markdownListMarker guifg=#9fae9f gui=NONE cterm=NONE
hi markdownRule guifg=#c7d4d6 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#c7d4d6 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#dcd4c5 gui=NONE cterm=NONE
hi DiagnosticError guifg=#94b7bd gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#b7a9c4 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#b8b09e gui=NONE cterm=NONE
hi DiagnosticHint guifg=#c1acc1 gui=NONE cterm=NONE
hi DiffAdd guifg=#b7a2b8 gui=NONE cterm=NONE
hi DiffDelete guifg=#94b7bd gui=NONE cterm=NONE
hi DiffChange guifg=#cfc1d7 gui=NONE cterm=NONE
hi DiffText guifg=#9fae9f gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#897f73', '#94b7bd', '#b7a2b8', '#b7a9c4', '#b8b09e', '#9fae9f', '#a3875b', '#9daeb1',
  \ '#9cb5bb', '#9bb5b9', '#bfa7c0', '#acacc4', '#b9b09b', '#9db79d', '#b7abab', '#9bb5bc']
