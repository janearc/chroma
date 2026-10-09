" pastel-invert -- pastel, inverted (the screen negative), its ground
" and the colours on it turned greener and darker by eye.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/pastel-invert.css;
" change the source, not this file.
"
"   :colorscheme pastel-invert-inverted
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
let g:colors_name = 'pastel-invert-inverted'

hi Normal guifg=#cacaca guibg=#77447d gui=NONE cterm=NONE
hi NormalFloat guifg=#cacaca guibg=#7d4d82 gui=NONE cterm=NONE
hi FloatBorder guifg=#7c9bf6 guibg=#7d4d82 gui=NONE cterm=NONE
hi CursorLine guibg=#7c4c82 gui=NONE cterm=NONE
hi CursorLineNr guifg=#f98df3 gui=bold cterm=bold
hi LineNr guifg=#7c9bf6 gui=NONE cterm=NONE
hi SignColumn guibg=#77447d gui=NONE cterm=NONE
hi Visual guibg=#70538e gui=NONE cterm=NONE
hi VertSplit guifg=#7c9bf6 gui=NONE cterm=NONE
hi WinSeparator guifg=#7c9bf6 gui=NONE cterm=NONE
hi StatusLine guifg=#cacaca guibg=#7d4d82 gui=NONE cterm=NONE
hi StatusLineNC guifg=#7c9bf6 guibg=#7d4d82 gui=NONE cterm=NONE
hi Pmenu guifg=#cacaca guibg=#7d4d82 gui=NONE cterm=NONE
hi PmenuSel guifg=#77447d guibg=#f98df3 gui=NONE cterm=NONE
hi Search guifg=#77447d guibg=#9edb86 gui=NONE cterm=NONE
hi IncSearch guifg=#77447d guibg=#f98df3 gui=NONE cterm=NONE
hi MatchParen guifg=#62b2ff gui=bold cterm=bold
hi Directory guifg=#62b2ff gui=NONE cterm=NONE
hi Folded guifg=#7c9bf6 guibg=#7d4d82 gui=NONE cterm=NONE
hi NonText guifg=#7c9bf6 gui=NONE cterm=NONE
hi Whitespace guifg=#7c9bf6 gui=NONE cterm=NONE
hi Conceal guifg=#7c9bf6 gui=NONE cterm=NONE
hi Title guifg=#f98df3 gui=bold cterm=bold
hi Comment guifg=#7c9bf6 gui=italic cterm=italic
hi String guifg=#9edb86 gui=NONE cterm=NONE
hi Character guifg=#9edb86 gui=NONE cterm=NONE
hi Number guifg=#9edb86 gui=NONE cterm=NONE
hi Boolean guifg=#9edb86 gui=NONE cterm=NONE
hi Identifier guifg=#cacaca gui=NONE cterm=NONE
hi Function guifg=#62b2ff gui=NONE cterm=NONE
hi Statement guifg=#f98df3 gui=NONE cterm=NONE
hi Keyword guifg=#f98df3 gui=NONE cterm=NONE
hi Operator guifg=#7c9bf6 gui=NONE cterm=NONE
hi PreProc guifg=#55c34c gui=NONE cterm=NONE
hi Type guifg=#55c34c gui=NONE cterm=NONE
hi Constant guifg=#9edb86 gui=NONE cterm=NONE
hi Special guifg=#62b2ff gui=NONE cterm=NONE
hi Todo guifg=#77447d guibg=#9edb86 gui=bold cterm=bold
hi Error guifg=#f99383 gui=NONE cterm=NONE
hi markdownH1 guifg=#f98df3 gui=bold cterm=bold
hi markdownH2 guifg=#f98df3 gui=bold cterm=bold
hi markdownH3 guifg=#f98df3 gui=NONE cterm=NONE
hi markdownH4 guifg=#f98df3 gui=NONE cterm=NONE
hi markdownH5 guifg=#62b2ff gui=NONE cterm=NONE
hi markdownH6 guifg=#62b2ff gui=NONE cterm=NONE
hi markdownCode guifg=#55c34c gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#55c34c gui=NONE cterm=NONE
hi markdownLinkText guifg=#62b2ff gui=underline cterm=underline
hi markdownUrl guifg=#7c9bf6 gui=NONE cterm=NONE
hi markdownListMarker guifg=#f98df3 gui=NONE cterm=NONE
hi markdownRule guifg=#7c9bf6 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#7c9bf6 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#7c9bf6 gui=NONE cterm=NONE
hi DiagnosticError guifg=#f99383 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#aeb34d gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#62b2ff gui=NONE cterm=NONE
hi DiagnosticHint guifg=#54bab8 gui=NONE cterm=NONE
hi DiffAdd guifg=#55c34c gui=NONE cterm=NONE
hi DiffDelete guifg=#f99383 gui=NONE cterm=NONE
hi DiffChange guifg=#9edb86 gui=NONE cterm=NONE
hi DiffText guifg=#f98df3 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#283165', '#f99383', '#55c34c', '#aeb34d', '#62b2ff', '#f98df3', '#54bab8', '#f5cca6',
  \ '#7c9bf6', '#fbc1b5', '#62e559', '#cdd059', '#aecbfa', '#fbbcf0', '#62ded3', '#e3e3e3']
