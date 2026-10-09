" fermi-midnight -- a dark colourway, tuned in paratune from fermi-long.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/fermi-midnight.css;
" change the source, not this file.
"
"   :colorscheme fermi-midnight-inverted
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
let g:colors_name = 'fermi-midnight-inverted'

hi Normal guifg=#817b73 guibg=#f5f1cf gui=NONE cterm=NONE
hi NormalFloat guifg=#817b73 guibg=#ece8c8 gui=NONE cterm=NONE
hi FloatBorder guifg=#b4a262 guibg=#ece8c8 gui=NONE cterm=NONE
hi CursorLine guibg=#eeeac9 gui=NONE cterm=NONE
hi CursorLineNr guifg=#57af5c gui=bold cterm=bold
hi LineNr guifg=#b4a262 gui=NONE cterm=NONE
hi SignColumn guibg=#f5f1cf gui=NONE cterm=NONE
hi Visual guibg=#d2cba7 gui=NONE cterm=NONE
hi VertSplit guifg=#b4a262 gui=NONE cterm=NONE
hi WinSeparator guifg=#b4a262 gui=NONE cterm=NONE
hi StatusLine guifg=#817b73 guibg=#ece8c8 gui=NONE cterm=NONE
hi StatusLineNC guifg=#b4a262 guibg=#ece8c8 gui=NONE cterm=NONE
hi Pmenu guifg=#817b73 guibg=#ece8c8 gui=NONE cterm=NONE
hi PmenuSel guifg=#f5f1cf guibg=#57af5c gui=NONE cterm=NONE
hi Search guifg=#f5f1cf guibg=#8582b8 gui=NONE cterm=NONE
hi IncSearch guifg=#f5f1cf guibg=#57af5c gui=NONE cterm=NONE
hi MatchParen guifg=#af8754 gui=bold cterm=bold
hi Directory guifg=#af8754 gui=NONE cterm=NONE
hi Folded guifg=#b4a262 guibg=#ece8c8 gui=NONE cterm=NONE
hi NonText guifg=#b4a262 gui=NONE cterm=NONE
hi Whitespace guifg=#b4a262 gui=NONE cterm=NONE
hi Conceal guifg=#b4a262 gui=NONE cterm=NONE
hi Title guifg=#57af5c gui=bold cterm=bold
hi Comment guifg=#b4a262 gui=italic cterm=italic
hi String guifg=#8582b8 gui=NONE cterm=NONE
hi Character guifg=#8582b8 gui=NONE cterm=NONE
hi Number guifg=#8582b8 gui=NONE cterm=NONE
hi Boolean guifg=#8582b8 gui=NONE cterm=NONE
hi Identifier guifg=#817b73 gui=NONE cterm=NONE
hi Function guifg=#af8754 gui=NONE cterm=NONE
hi Statement guifg=#57af5c gui=NONE cterm=NONE
hi Keyword guifg=#57af5c gui=NONE cterm=NONE
hi Operator guifg=#b4a262 gui=NONE cterm=NONE
hi PreProc guifg=#ac76b4 gui=NONE cterm=NONE
hi Type guifg=#ac76b4 gui=NONE cterm=NONE
hi Constant guifg=#8582b8 gui=NONE cterm=NONE
hi Special guifg=#af8754 gui=NONE cterm=NONE
hi Todo guifg=#f5f1cf guibg=#8582b8 gui=bold cterm=bold
hi Error guifg=#55a0ad gui=NONE cterm=NONE
hi markdownH1 guifg=#57af5c gui=bold cterm=bold
hi markdownH2 guifg=#57af5c gui=bold cterm=bold
hi markdownH3 guifg=#57af5c gui=NONE cterm=NONE
hi markdownH4 guifg=#57af5c gui=NONE cterm=NONE
hi markdownH5 guifg=#af8754 gui=NONE cterm=NONE
hi markdownH6 guifg=#af8754 gui=NONE cterm=NONE
hi markdownCode guifg=#ac76b4 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#ac76b4 gui=NONE cterm=NONE
hi markdownLinkText guifg=#af8754 gui=underline cterm=underline
hi markdownUrl guifg=#b4a262 gui=NONE cterm=NONE
hi markdownListMarker guifg=#57af5c gui=NONE cterm=NONE
hi markdownRule guifg=#b4a262 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#b4a262 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#b4a262 gui=NONE cterm=NONE
hi DiagnosticError guifg=#55a0ad gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#8582b8 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#af8754 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#ac797a gui=NONE cterm=NONE
hi DiffAdd guifg=#ac76b4 gui=NONE cterm=NONE
hi DiffDelete guifg=#55a0ad gui=NONE cterm=NONE
hi DiffChange guifg=#8582b8 gui=NONE cterm=NONE
hi DiffText guifg=#57af5c gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#ece6c6', '#55a0ad', '#ac76b4', '#8582b8', '#af8754', '#57af5c', '#ac797a', '#507b9d',
  \ '#b4a262', '#4e838d', '#a361ab', '#716dad', '#917451', '#548c59', '#a16568', '#416072']
