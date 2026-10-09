" twilight-low -- made in paratune from twilight.
"
" rendered by colourway from sources/twilight-low.css;
" change the source, not this file.
"
"   :colorscheme twilight-low
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
let g:colors_name = 'twilight-low'

hi Normal guifg=#c6deff guibg=#eb3f12 gui=NONE cterm=NONE
hi NormalFloat guifg=#c6deff guibg=#ffd3b1 gui=NONE cterm=NONE
hi FloatBorder guifg=#ffffff guibg=#ffd3b1 gui=NONE cterm=NONE
hi CursorLine guibg=#ffcea2 gui=NONE cterm=NONE
hi CursorLineNr guifg=#ff5dd3 gui=bold cterm=bold
hi LineNr guifg=#ffffff gui=NONE cterm=NONE
hi SignColumn guibg=#eb3f12 gui=NONE cterm=NONE
hi Visual guibg=#ff7a1a gui=NONE cterm=NONE
hi VertSplit guifg=#ffffff gui=NONE cterm=NONE
hi WinSeparator guifg=#ffffff gui=NONE cterm=NONE
hi StatusLine guifg=#c6deff guibg=#ffd3b1 gui=NONE cterm=NONE
hi StatusLineNC guifg=#ffffff guibg=#ffd3b1 gui=NONE cterm=NONE
hi Pmenu guifg=#c6deff guibg=#ffd3b1 gui=NONE cterm=NONE
hi PmenuSel guifg=#eb3f12 guibg=#ff5dd3 gui=NONE cterm=NONE
hi Search guifg=#eb3f12 guibg=#daae3b gui=NONE cterm=NONE
hi IncSearch guifg=#eb3f12 guibg=#ff5dd3 gui=NONE cterm=NONE
hi MatchParen guifg=#7ca3ff gui=bold cterm=bold
hi Directory guifg=#7ca3ff gui=NONE cterm=NONE
hi Folded guifg=#d4a239 guibg=#ffd3b1 gui=NONE cterm=NONE
hi NonText guifg=#ffffff gui=NONE cterm=NONE
hi Whitespace guifg=#ffffff gui=NONE cterm=NONE
hi Conceal guifg=#ffffff gui=NONE cterm=NONE
hi Title guifg=#ff5dd3 gui=bold cterm=bold
hi Comment guifg=#d4a239 gui=italic cterm=italic
hi String guifg=#daae3b gui=NONE cterm=NONE
hi Character guifg=#daae3b gui=NONE cterm=NONE
hi Number guifg=#daae3b gui=NONE cterm=NONE
hi Boolean guifg=#daae3b gui=NONE cterm=NONE
hi Identifier guifg=#c6deff gui=NONE cterm=NONE
hi Function guifg=#7ca3ff gui=NONE cterm=NONE
hi Statement guifg=#ff5dd3 gui=NONE cterm=NONE
hi Keyword guifg=#ff5dd3 gui=NONE cterm=NONE
hi Operator guifg=#d4a239 gui=NONE cterm=NONE
hi PreProc guifg=#7fb337 gui=NONE cterm=NONE
hi Type guifg=#7fb337 gui=NONE cterm=NONE
hi Constant guifg=#daae3b gui=NONE cterm=NONE
hi Special guifg=#7ca3ff gui=NONE cterm=NONE
hi Todo guifg=#eb3f12 guibg=#daae3b gui=bold cterm=bold
hi Error guifg=#ff8b85 gui=NONE cterm=NONE
hi markdownH1 guifg=#ff5dd3 gui=bold cterm=bold
hi markdownH2 guifg=#ff5dd3 gui=bold cterm=bold
hi markdownH3 guifg=#ff5dd3 gui=NONE cterm=NONE
hi markdownH4 guifg=#ff5dd3 gui=NONE cterm=NONE
hi markdownH5 guifg=#7ca3ff gui=NONE cterm=NONE
hi markdownH6 guifg=#7ca3ff gui=NONE cterm=NONE
hi markdownCode guifg=#7fb337 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#7fb337 gui=NONE cterm=NONE
hi markdownLinkText guifg=#7ca3ff gui=underline cterm=underline
hi markdownUrl guifg=#d4a239 gui=NONE cterm=NONE
hi markdownListMarker guifg=#ff5dd3 gui=NONE cterm=NONE
hi markdownRule guifg=#ffffff gui=NONE cterm=NONE
hi markdownBlockquote guifg=#d4a239 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#ffffff gui=NONE cterm=NONE
hi DiagnosticError guifg=#ff8b85 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#daae3b gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#7ca3ff gui=NONE cterm=NONE
hi DiagnosticHint guifg=#7fb337 gui=NONE cterm=NONE
hi DiffAdd guifg=#7fb337 gui=NONE cterm=NONE
hi DiffDelete guifg=#ff8b85 gui=NONE cterm=NONE
hi DiffChange guifg=#daae3b gui=NONE cterm=NONE
hi DiffText guifg=#ff5dd3 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#106d79', '#ffffff', '#d2ffaa', '#fff3d3', '#eef4ff', '#ffffff', '#bbfdff', '#fffcfa',
  \ '#a3a3ae', '#ffffff', '#d1ffa1', '#fff3d5', '#f0f5ff', '#ffffff', '#b2fffe', '#ffffff']
