" fermi-midnight -- a dark colourway, tuned in paratune from fermi-long.
"
" rendered by colourway from sources/fermi-midnight.css;
" change the source, not this file.
"
"   :colorscheme fermi-midnight
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
let g:colors_name = 'fermi-midnight'

hi Normal guifg=#7e848c guibg=#0a0e30 gui=NONE cterm=NONE
hi NormalFloat guifg=#7e848c guibg=#101636 gui=NONE cterm=NONE
hi FloatBorder guifg=#4b5d9d guibg=#101636 gui=NONE cterm=NONE
hi CursorLine guibg=#0f1535 gui=NONE cterm=NONE
hi CursorLineNr guifg=#a850a3 gui=bold cterm=bold
hi LineNr guifg=#4b5d9d gui=NONE cterm=NONE
hi SignColumn guibg=#0a0e30 gui=NONE cterm=NONE
hi Visual guibg=#2d3458 gui=NONE cterm=NONE
hi VertSplit guifg=#4b5d9d gui=NONE cterm=NONE
hi WinSeparator guifg=#4b5d9d gui=NONE cterm=NONE
hi StatusLine guifg=#7e848c guibg=#101636 gui=NONE cterm=NONE
hi StatusLineNC guifg=#4b5d9d guibg=#101636 gui=NONE cterm=NONE
hi Pmenu guifg=#7e848c guibg=#101636 gui=NONE cterm=NONE
hi PmenuSel guifg=#0a0e30 guibg=#a850a3 gui=NONE cterm=NONE
hi Search guifg=#0a0e30 guibg=#7a7d47 gui=NONE cterm=NONE
hi IncSearch guifg=#0a0e30 guibg=#a850a3 gui=NONE cterm=NONE
hi MatchParen guifg=#5078ab gui=bold cterm=bold
hi Directory guifg=#5078ab gui=NONE cterm=NONE
hi Folded guifg=#4b5d9d guibg=#101636 gui=NONE cterm=NONE
hi NonText guifg=#4b5d9d gui=NONE cterm=NONE
hi Whitespace guifg=#4b5d9d gui=NONE cterm=NONE
hi Conceal guifg=#4b5d9d gui=NONE cterm=NONE
hi Title guifg=#a850a3 gui=bold cterm=bold
hi Comment guifg=#4b5d9d gui=italic cterm=italic
hi String guifg=#7a7d47 gui=NONE cterm=NONE
hi Character guifg=#7a7d47 gui=NONE cterm=NONE
hi Number guifg=#7a7d47 gui=NONE cterm=NONE
hi Boolean guifg=#7a7d47 gui=NONE cterm=NONE
hi Identifier guifg=#7e848c gui=NONE cterm=NONE
hi Function guifg=#5078ab gui=NONE cterm=NONE
hi Statement guifg=#a850a3 gui=NONE cterm=NONE
hi Keyword guifg=#a850a3 gui=NONE cterm=NONE
hi Operator guifg=#4b5d9d gui=NONE cterm=NONE
hi PreProc guifg=#53894b gui=NONE cterm=NONE
hi Type guifg=#53894b gui=NONE cterm=NONE
hi Constant guifg=#7a7d47 gui=NONE cterm=NONE
hi Special guifg=#5078ab gui=NONE cterm=NONE
hi Todo guifg=#0a0e30 guibg=#7a7d47 gui=bold cterm=bold
hi Error guifg=#aa5f52 gui=NONE cterm=NONE
hi markdownH1 guifg=#a850a3 gui=bold cterm=bold
hi markdownH2 guifg=#a850a3 gui=bold cterm=bold
hi markdownH3 guifg=#a850a3 gui=NONE cterm=NONE
hi markdownH4 guifg=#a850a3 gui=NONE cterm=NONE
hi markdownH5 guifg=#5078ab gui=NONE cterm=NONE
hi markdownH6 guifg=#5078ab gui=NONE cterm=NONE
hi markdownCode guifg=#53894b gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#53894b gui=NONE cterm=NONE
hi markdownLinkText guifg=#5078ab gui=underline cterm=underline
hi markdownUrl guifg=#4b5d9d gui=NONE cterm=NONE
hi markdownListMarker guifg=#a850a3 gui=NONE cterm=NONE
hi markdownRule guifg=#4b5d9d gui=NONE cterm=NONE
hi markdownBlockquote guifg=#4b5d9d gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#4b5d9d gui=NONE cterm=NONE
hi DiagnosticError guifg=#aa5f52 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#7a7d47 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#5078ab gui=NONE cterm=NONE
hi DiagnosticHint guifg=#538685 gui=NONE cterm=NONE
hi DiffAdd guifg=#53894b gui=NONE cterm=NONE
hi DiffDelete guifg=#aa5f52 gui=NONE cterm=NONE
hi DiffChange guifg=#7a7d47 gui=NONE cterm=NONE
hi DiffText guifg=#a850a3 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#131939', '#aa5f52', '#53894b', '#7a7d47', '#5078ab', '#a850a3', '#538685', '#af8462',
  \ '#4b5d9d', '#b17c72', '#5c9e54', '#8e9252', '#6e8bae', '#ab73a6', '#5e9a97', '#be9f8d']
