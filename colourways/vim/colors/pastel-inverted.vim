" pastel -- made in paratune from pastel.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/pastel.css;
" change the source, not this file.
"
"   :colorscheme pastel-inverted
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
let g:colors_name = 'pastel-inverted'

hi Normal guifg=#353535 guibg=#b9bb5b gui=NONE cterm=NONE
hi NormalFloat guifg=#353535 guibg=#afb159 gui=NONE cterm=NONE
hi FloatBorder guifg=#836409 guibg=#afb159 gui=NONE cterm=NONE
hi CursorLine guibg=#b0b259 gui=NONE cterm=NONE
hi CursorLineNr guifg=#06720c gui=bold cterm=bold
hi LineNr guifg=#836409 gui=NONE cterm=NONE
hi SignColumn guibg=#b9bb5b gui=NONE cterm=NONE
hi Visual guibg=#bcaa55 gui=NONE cterm=NONE
hi VertSplit guifg=#836409 gui=NONE cterm=NONE
hi WinSeparator guifg=#836409 gui=NONE cterm=NONE
hi StatusLine guifg=#353535 guibg=#afb159 gui=NONE cterm=NONE
hi StatusLineNC guifg=#836409 guibg=#afb159 gui=NONE cterm=NONE
hi Pmenu guifg=#353535 guibg=#afb159 gui=NONE cterm=NONE
hi PmenuSel guifg=#b9bb5b guibg=#06720c gui=NONE cterm=NONE
hi Search guifg=#b9bb5b guibg=#612479 gui=NONE cterm=NONE
hi IncSearch guifg=#b9bb5b guibg=#06720c gui=NONE cterm=NONE
hi MatchParen guifg=#9d4d00 gui=bold cterm=bold
hi Directory guifg=#9d4d00 gui=NONE cterm=NONE
hi Folded guifg=#836409 guibg=#afb159 gui=NONE cterm=NONE
hi NonText guifg=#836409 gui=NONE cterm=NONE
hi Whitespace guifg=#836409 gui=NONE cterm=NONE
hi Conceal guifg=#836409 gui=NONE cterm=NONE
hi Title guifg=#06720c gui=bold cterm=bold
hi Comment guifg=#836409 gui=italic cterm=italic
hi String guifg=#612479 gui=NONE cterm=NONE
hi Character guifg=#612479 gui=NONE cterm=NONE
hi Number guifg=#612479 gui=NONE cterm=NONE
hi Boolean guifg=#612479 gui=NONE cterm=NONE
hi Identifier guifg=#353535 gui=NONE cterm=NONE
hi Function guifg=#9d4d00 gui=NONE cterm=NONE
hi Statement guifg=#06720c gui=NONE cterm=NONE
hi Keyword guifg=#06720c gui=NONE cterm=NONE
hi Operator guifg=#836409 gui=NONE cterm=NONE
hi PreProc guifg=#aa3cb3 gui=NONE cterm=NONE
hi Type guifg=#aa3cb3 gui=NONE cterm=NONE
hi Constant guifg=#612479 gui=NONE cterm=NONE
hi Special guifg=#9d4d00 gui=NONE cterm=NONE
hi Todo guifg=#b9bb5b guibg=#612479 gui=bold cterm=bold
hi Error guifg=#066c7c gui=NONE cterm=NONE
hi markdownH1 guifg=#06720c gui=bold cterm=bold
hi markdownH2 guifg=#06720c gui=bold cterm=bold
hi markdownH3 guifg=#06720c gui=NONE cterm=NONE
hi markdownH4 guifg=#06720c gui=NONE cterm=NONE
hi markdownH5 guifg=#9d4d00 gui=NONE cterm=NONE
hi markdownH6 guifg=#9d4d00 gui=NONE cterm=NONE
hi markdownCode guifg=#aa3cb3 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#aa3cb3 gui=NONE cterm=NONE
hi markdownLinkText guifg=#9d4d00 gui=underline cterm=underline
hi markdownUrl guifg=#836409 gui=NONE cterm=NONE
hi markdownListMarker guifg=#06720c gui=NONE cterm=NONE
hi markdownRule guifg=#836409 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#836409 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#836409 gui=NONE cterm=NONE
hi DiagnosticError guifg=#066c7c gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#514cb2 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#9d4d00 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#ab4547 gui=NONE cterm=NONE
hi DiffAdd guifg=#aa3cb3 gui=NONE cterm=NONE
hi DiffDelete guifg=#066c7c gui=NONE cterm=NONE
hi DiffChange guifg=#612479 gui=NONE cterm=NONE
hi DiffText guifg=#06720c gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#d7ce9a', '#066c7c', '#aa3cb3', '#514cb2', '#9d4d00', '#06720c', '#ab4547', '#0a3359',
  \ '#836409', '#043e4a', '#9d1aa6', '#322fa6', '#513405', '#04430f', '#9d212c', '#1c1c1c']
