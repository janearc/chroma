" twilight -- a ground that is not dark and letters that are not bright,
" so the ground cannot shine up through the text in rivers: an amber
" ground with deep navy ink, 8.9:1, chosen in paratune. dark text on a
" lightish ground.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/twilight.css;
" change the source, not this file.
"
"   :colorscheme twilight-inverted
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
let g:colors_name = 'twilight-inverted'

hi Normal guifg=#fae8a6 guibg=#0053b1 gui=NONE cterm=NONE
hi NormalFloat guifg=#fae8a6 guibg=#0959b4 gui=NONE cterm=NONE
hi FloatBorder guifg=#3b7db5 guibg=#0959b4 gui=NONE cterm=NONE
hi CursorLine guibg=#0f5db6 gui=NONE cterm=NONE
hi CursorLineNr guifg=#a5d5b5 gui=bold cterm=bold
hi LineNr guifg=#4789bc gui=NONE cterm=NONE
hi SignColumn guibg=#0053b1 gui=NONE cterm=NONE
hi Visual guibg=#1769c5 gui=NONE cterm=NONE
hi VertSplit guifg=#3b7db5 gui=NONE cterm=NONE
hi WinSeparator guifg=#3b7db5 gui=NONE cterm=NONE
hi StatusLine guifg=#fae8a6 guibg=#0959b4 gui=NONE cterm=NONE
hi StatusLineNC guifg=#4789bc guibg=#0959b4 gui=NONE cterm=NONE
hi Pmenu guifg=#fae8a6 guibg=#0959b4 gui=NONE cterm=NONE
hi PmenuSel guifg=#0053b1 guibg=#a5d5b5 gui=NONE cterm=NONE
hi Search guifg=#0053b1 guibg=#a5bbef gui=NONE cterm=NONE
hi IncSearch guifg=#0053b1 guibg=#a5d5b5 gui=NONE cterm=NONE
hi MatchParen guifg=#dbc9a5 gui=bold cterm=bold
hi Directory guifg=#dbc9a5 gui=NONE cterm=NONE
hi Folded guifg=#b5c2d4 guibg=#0959b4 gui=NONE cterm=NONE
hi NonText guifg=#4789bc gui=NONE cterm=NONE
hi Whitespace guifg=#4789bc gui=NONE cterm=NONE
hi Conceal guifg=#4789bc gui=NONE cterm=NONE
hi Title guifg=#a5d5b5 gui=bold cterm=bold
hi Comment guifg=#b5c2d4 gui=italic cterm=italic
hi String guifg=#a5bbef gui=NONE cterm=NONE
hi Character guifg=#a5bbef gui=NONE cterm=NONE
hi Number guifg=#a5bbef gui=NONE cterm=NONE
hi Boolean guifg=#a5bbef gui=NONE cterm=NONE
hi Identifier guifg=#fae8a6 gui=NONE cterm=NONE
hi Function guifg=#dbc9a5 gui=NONE cterm=NONE
hi Statement guifg=#a5d5b5 gui=NONE cterm=NONE
hi Keyword guifg=#a5d5b5 gui=NONE cterm=NONE
hi Operator guifg=#b5c2d4 gui=NONE cterm=NONE
hi PreProc guifg=#d0c5e0 gui=NONE cterm=NONE
hi Type guifg=#d0c5e0 gui=NONE cterm=NONE
hi Constant guifg=#a5bbef gui=NONE cterm=NONE
hi Special guifg=#dbc9a5 gui=NONE cterm=NONE
hi Todo guifg=#0053b1 guibg=#a5bbef gui=bold cterm=bold
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
hi markdownRule guifg=#3b7db5 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#b5c2d4 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#4789bc gui=NONE cterm=NONE
hi DiagnosticError guifg=#91d5e0 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#a5bbef gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#dbc9a5 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#d0c5e0 gui=NONE cterm=NONE
hi DiffAdd guifg=#d0c5e0 gui=NONE cterm=NONE
hi DiffDelete guifg=#91d5e0 gui=NONE cterm=NONE
hi DiffChange guifg=#a5bbef gui=NONE cterm=NONE
hi DiffText guifg=#a5d5b5 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#d5dbe7', '#91d5e0', '#d0c5e0', '#a5bbef', '#dbc9a5', '#a5d5b5', '#e0b5b5', '#1363bd',
  \ '#b5c2d4', '#7bc9d5', '#c4b7d9', '#91abe7', '#d0bb8f', '#91cba6', '#d7a3a3', '#004291']
