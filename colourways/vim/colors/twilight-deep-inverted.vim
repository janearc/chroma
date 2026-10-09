" twilight-deep -- twilight's ground taken darker: ink #38241e on ground
" #808a63, 4.0:1, chosen in paratune, with claude code's own colours tuned
" to match. the palette is darker so the accents still read on the deeper
" ground, and the statusline's bar is twilight's ramp moved by the same
" shift.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/twilight-deep.css;
" change the source, not this file.
"
"   :colorscheme twilight-deep-inverted
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
let g:colors_name = 'twilight-deep-inverted'

hi Normal guifg=#c7dbe1 guibg=#7f759c gui=NONE cterm=NONE
hi NormalFloat guifg=#c7dbe1 guibg=#7077a9 gui=NONE cterm=NONE
hi FloatBorder guifg=#9197c1 guibg=#7077a9 gui=NONE cterm=NONE
hi CursorLine guibg=#737aac gui=NONE cterm=NONE
hi CursorLineNr guifg=#c5e7cf gui=bold cterm=bold
hi LineNr guifg=#9197c1 gui=NONE cterm=NONE
hi SignColumn guibg=#7f759c gui=NONE cterm=NONE
hi Visual guibg=#8086b4 gui=NONE cterm=NONE
hi VertSplit guifg=#9197c1 gui=NONE cterm=NONE
hi WinSeparator guifg=#9197c1 gui=NONE cterm=NONE
hi StatusLine guifg=#c7dbe1 guibg=#7077a9 gui=NONE cterm=NONE
hi StatusLineNC guifg=#9197c1 guibg=#7077a9 gui=NONE cterm=NONE
hi Pmenu guifg=#c7dbe1 guibg=#7077a9 gui=NONE cterm=NONE
hi PmenuSel guifg=#7f759c guibg=#c5e7cf gui=NONE cterm=NONE
hi Search guifg=#7f759c guibg=#c5d3f7 gui=NONE cterm=NONE
hi IncSearch guifg=#7f759c guibg=#c5e7cf gui=NONE cterm=NONE
hi MatchParen guifg=#ebe3c5 gui=bold cterm=bold
hi Directory guifg=#ebe3c5 gui=NONE cterm=NONE
hi Folded guifg=#d1d5e3 guibg=#7077a9 gui=NONE cterm=NONE
hi NonText guifg=#9197c1 gui=NONE cterm=NONE
hi Whitespace guifg=#9197c1 gui=NONE cterm=NONE
hi Conceal guifg=#9197c1 gui=NONE cterm=NONE
hi Title guifg=#c5e7cf gui=bold cterm=bold
hi Comment guifg=#d1d5e3 gui=italic cterm=italic
hi String guifg=#c5d3f7 gui=NONE cterm=NONE
hi Character guifg=#c5d3f7 gui=NONE cterm=NONE
hi Number guifg=#c5d3f7 gui=NONE cterm=NONE
hi Boolean guifg=#c5d3f7 gui=NONE cterm=NONE
hi Identifier guifg=#c7dbe1 gui=NONE cterm=NONE
hi Function guifg=#ebe3c5 gui=NONE cterm=NONE
hi Statement guifg=#c5e7cf gui=NONE cterm=NONE
hi Keyword guifg=#c5e7cf gui=NONE cterm=NONE
hi Operator guifg=#d1d5e3 gui=NONE cterm=NONE
hi PreProc guifg=#e3d7ef gui=NONE cterm=NONE
hi Type guifg=#e3d7ef gui=NONE cterm=NONE
hi Constant guifg=#c5d3f7 gui=NONE cterm=NONE
hi Special guifg=#ebe3c5 gui=NONE cterm=NONE
hi Todo guifg=#7f759c guibg=#c5d3f7 gui=bold cterm=bold
hi Error guifg=#b5e7ef gui=NONE cterm=NONE
hi markdownH1 guifg=#c5e7cf gui=bold cterm=bold
hi markdownH2 guifg=#c5e7cf gui=bold cterm=bold
hi markdownH3 guifg=#c5e7cf gui=NONE cterm=NONE
hi markdownH4 guifg=#c5e7cf gui=NONE cterm=NONE
hi markdownH5 guifg=#ebe3c5 gui=NONE cterm=NONE
hi markdownH6 guifg=#ebe3c5 gui=NONE cterm=NONE
hi markdownCode guifg=#e3d7ef gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#e3d7ef gui=NONE cterm=NONE
hi markdownLinkText guifg=#ebe3c5 gui=underline cterm=underline
hi markdownUrl guifg=#d1d5e3 gui=NONE cterm=NONE
hi markdownListMarker guifg=#c5e7cf gui=NONE cterm=NONE
hi markdownRule guifg=#9197c1 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#d1d5e3 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#9197c1 gui=NONE cterm=NONE
hi DiagnosticError guifg=#b5e7ef gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#c5d3f7 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#ebe3c5 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#e3d7ef gui=NONE cterm=NONE
hi DiffAdd guifg=#e3d7ef gui=NONE cterm=NONE
hi DiffDelete guifg=#b5e7ef gui=NONE cterm=NONE
hi DiffChange guifg=#c5d3f7 gui=NONE cterm=NONE
hi DiffText guifg=#c5e7cf gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#e3e5ef', '#b5e7ef', '#e3d7ef', '#c5d3f7', '#ebe3c5', '#c5e7cf', '#efcfd1', '#8086b4',
  \ '#d1d5e3', '#a5dfea', '#d9cdeb', '#b5c5f3', '#e3d7af', '#b5dfbf', '#e9c3c5', '#565d96']
