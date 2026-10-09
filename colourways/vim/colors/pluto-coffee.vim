" pluto-coffee -- made in paratune from pluto-night.
"
" rendered by colourway from sources/pluto-coffee.css;
" change the source, not this file.
"
"   :colorscheme pluto-coffee
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
let g:colors_name = 'pluto-coffee'

hi Normal guifg=#67717c guibg=#0d0d10 gui=NONE cterm=NONE
hi NormalFloat guifg=#67717c guibg=#121316 gui=NONE cterm=NONE
hi FloatBorder guifg=#382b29 guibg=#121316 gui=NONE cterm=NONE
hi CursorLine guibg=#121215 gui=NONE cterm=NONE
hi CursorLineNr guifg=#605160 gui=bold cterm=bold
hi LineNr guifg=#232b3a gui=NONE cterm=NONE
hi SignColumn guibg=#0d0d10 gui=NONE cterm=NONE
hi Visual guibg=#383838 gui=NONE cterm=NONE
hi VertSplit guifg=#382b29 gui=NONE cterm=NONE
hi WinSeparator guifg=#382b29 gui=NONE cterm=NONE
hi StatusLine guifg=#67717c guibg=#121316 gui=NONE cterm=NONE
hi StatusLineNC guifg=#232b3a guibg=#121316 gui=NONE cterm=NONE
hi Pmenu guifg=#67717c guibg=#121316 gui=NONE cterm=NONE
hi PmenuSel guifg=#0d0d10 guibg=#605160 gui=NONE cterm=NONE
hi Search guifg=#0d0d10 guibg=#303e28 gui=NONE cterm=NONE
hi IncSearch guifg=#0d0d10 guibg=#605160 gui=NONE cterm=NONE
hi MatchParen guifg=#474f61 gui=bold cterm=bold
hi Directory guifg=#474f61 gui=NONE cterm=NONE
hi Folded guifg=#382b29 guibg=#121316 gui=NONE cterm=NONE
hi NonText guifg=#232b3a gui=NONE cterm=NONE
hi Whitespace guifg=#232b3a gui=NONE cterm=NONE
hi Conceal guifg=#232b3a gui=NONE cterm=NONE
hi Title guifg=#605160 gui=bold cterm=bold
hi Comment guifg=#382b29 gui=italic cterm=italic
hi String guifg=#303e28 gui=NONE cterm=NONE
hi Character guifg=#303e28 gui=NONE cterm=NONE
hi Number guifg=#303e28 gui=NONE cterm=NONE
hi Boolean guifg=#303e28 gui=NONE cterm=NONE
hi Identifier guifg=#67717c gui=NONE cterm=NONE
hi Function guifg=#474f61 gui=NONE cterm=NONE
hi Statement guifg=#605160 gui=NONE cterm=NONE
hi Keyword guifg=#605160 gui=NONE cterm=NONE
hi Operator guifg=#382b29 gui=NONE cterm=NONE
hi PreProc guifg=#485d47 gui=NONE cterm=NONE
hi Type guifg=#485d47 gui=NONE cterm=NONE
hi Constant guifg=#303e28 gui=NONE cterm=NONE
hi Special guifg=#474f61 gui=NONE cterm=NONE
hi Todo guifg=#0d0d10 guibg=#303e28 gui=bold cterm=bold
hi Error guifg=#6b4842 gui=NONE cterm=NONE
hi markdownH1 guifg=#605160 gui=bold cterm=bold
hi markdownH2 guifg=#605160 gui=bold cterm=bold
hi markdownH3 guifg=#605160 gui=NONE cterm=NONE
hi markdownH4 guifg=#605160 gui=NONE cterm=NONE
hi markdownH5 guifg=#474f61 gui=NONE cterm=NONE
hi markdownH6 guifg=#474f61 gui=NONE cterm=NONE
hi markdownCode guifg=#485d47 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#485d47 gui=NONE cterm=NONE
hi markdownLinkText guifg=#474f61 gui=underline cterm=underline
hi markdownUrl guifg=#382b29 gui=NONE cterm=NONE
hi markdownListMarker guifg=#605160 gui=NONE cterm=NONE
hi markdownRule guifg=#382b29 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#382b29 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#232b3a gui=NONE cterm=NONE
hi DiagnosticError guifg=#6b4842 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#48563b gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#474f61 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#3e533e gui=NONE cterm=NONE
hi DiffAdd guifg=#485d47 gui=NONE cterm=NONE
hi DiffDelete guifg=#6b4842 gui=NONE cterm=NONE
hi DiffChange guifg=#303e28 gui=NONE cterm=NONE
hi DiffText guifg=#605160 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#76808c', '#6b4842', '#485d47', '#48563b', '#474f61', '#605160', '#5c78a4', '#62514e',
  \ '#634a44', '#644a46', '#40583f', '#53533b', '#464f64', '#624862', '#485454', '#644a43']
