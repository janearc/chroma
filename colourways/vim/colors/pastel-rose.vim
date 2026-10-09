" pastel-rose -- made in paratune from pastel-rose.
"
" rendered by colourway from sources/pastel-rose.css;
" change the source, not this file.
"
"   :colorscheme pastel-rose
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
let g:colors_name = 'pastel-rose'

hi Normal guifg=#020709 guibg=#f082f0 gui=NONE cterm=NONE
hi NormalFloat guifg=#020709 guibg=#dc78dd gui=NONE cterm=NONE
hi FloatBorder guifg=#080609 guibg=#dc78dd gui=NONE cterm=NONE
hi CursorLine guibg=#df7ae0 gui=NONE cterm=NONE
hi CursorLineNr guifg=#79df00 gui=bold cterm=bold
hi LineNr guifg=#080609 gui=NONE cterm=NONE
hi SignColumn guibg=#f082f0 gui=NONE cterm=NONE
hi Visual guibg=#fb58e3 gui=NONE cterm=NONE
hi VertSplit guifg=#080609 gui=NONE cterm=NONE
hi WinSeparator guifg=#080609 gui=NONE cterm=NONE
hi StatusLine guifg=#020709 guibg=#dc78dd gui=NONE cterm=NONE
hi StatusLineNC guifg=#080609 guibg=#dc78dd gui=NONE cterm=NONE
hi Pmenu guifg=#020709 guibg=#dc78dd gui=NONE cterm=NONE
hi PmenuSel guifg=#f082f0 guibg=#79df00 gui=NONE cterm=NONE
hi Search guifg=#f082f0 guibg=#000905 gui=NONE cterm=NONE
hi IncSearch guifg=#f082f0 guibg=#79df00 gui=NONE cterm=NONE
hi MatchParen guifg=#00001a gui=bold cterm=bold
hi Directory guifg=#00001a gui=NONE cterm=NONE
hi Folded guifg=#080609 guibg=#dc78dd gui=NONE cterm=NONE
hi NonText guifg=#080609 gui=NONE cterm=NONE
hi Whitespace guifg=#080609 gui=NONE cterm=NONE
hi Conceal guifg=#080609 gui=NONE cterm=NONE
hi Title guifg=#79df00 gui=bold cterm=bold
hi Comment guifg=#080609 gui=italic cterm=italic
hi String guifg=#000905 gui=NONE cterm=NONE
hi Character guifg=#000905 gui=NONE cterm=NONE
hi Number guifg=#000905 gui=NONE cterm=NONE
hi Boolean guifg=#000905 gui=NONE cterm=NONE
hi Identifier guifg=#020709 gui=NONE cterm=NONE
hi Function guifg=#00001a gui=NONE cterm=NONE
hi Statement guifg=#79df00 gui=NONE cterm=NONE
hi Keyword guifg=#79df00 gui=NONE cterm=NONE
hi Operator guifg=#080609 gui=NONE cterm=NONE
hi PreProc guifg=#ffffff gui=NONE cterm=NONE
hi Type guifg=#ffffff gui=NONE cterm=NONE
hi Constant guifg=#000905 gui=NONE cterm=NONE
hi Special guifg=#00001a gui=NONE cterm=NONE
hi Todo guifg=#f082f0 guibg=#000905 gui=bold cterm=bold
hi Error guifg=#3cdddd gui=NONE cterm=NONE
hi markdownH1 guifg=#79df00 gui=bold cterm=bold
hi markdownH2 guifg=#79df00 gui=bold cterm=bold
hi markdownH3 guifg=#79df00 gui=NONE cterm=NONE
hi markdownH4 guifg=#79df00 gui=NONE cterm=NONE
hi markdownH5 guifg=#00001a gui=NONE cterm=NONE
hi markdownH6 guifg=#00001a gui=NONE cterm=NONE
hi markdownCode guifg=#ffffff gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#ffffff gui=NONE cterm=NONE
hi markdownLinkText guifg=#00001a gui=underline cterm=underline
hi markdownUrl guifg=#080609 gui=NONE cterm=NONE
hi markdownListMarker guifg=#79df00 gui=NONE cterm=NONE
hi markdownRule guifg=#080609 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#080609 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#080609 gui=NONE cterm=NONE
hi DiagnosticError guifg=#3cdddd gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#f0e217 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#00001a gui=NONE cterm=NONE
hi DiagnosticHint guifg=#20ffff gui=NONE cterm=NONE
hi DiffAdd guifg=#ffffff gui=NONE cterm=NONE
hi DiffDelete guifg=#3cdddd gui=NONE cterm=NONE
hi DiffChange guifg=#000905 gui=NONE cterm=NONE
hi DiffText guifg=#79df00 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#e577ff', '#3cdddd', '#ffffff', '#ffff1d', '#00001a', '#89fc00', '#20ffff', '#ffffff',
  \ '#080609', '#a0254b', '#0a6aa0', '#ffff1d', '#08020e', '#00ff0b', '#20ffff', '#ffffff']
