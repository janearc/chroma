" twilight-burnt -- burnt green ink on a darker orange, twilight as it began
" on 2026-10-01, before the ground went bright; with claude code's colours
" dark on the orange.
"
" rendered by colourway from sources/twilight-burnt.css;
" change the source, not this file.
"
"   :colorscheme twilight-burnt
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
let g:colors_name = 'twilight-burnt'

hi Normal guifg=#1c2210 guibg=#b37f53 gui=NONE cterm=NONE
hi NormalFloat guifg=#1c2210 guibg=#a7784e gui=NONE cterm=NONE
hi FloatBorder guifg=#4a3d2b guibg=#a7784e gui=NONE cterm=NONE
hi CursorLine guibg=#a9794f gui=NONE cterm=NONE
hi CursorLineNr guifg=#5a2a4a gui=bold cterm=bold
hi LineNr guifg=#4a3d2b gui=NONE cterm=NONE
hi SignColumn guibg=#b37f53 gui=NONE cterm=NONE
hi Visual guibg=#d4a070 gui=NONE cterm=NONE
hi VertSplit guifg=#4a3d2b gui=NONE cterm=NONE
hi WinSeparator guifg=#4a3d2b gui=NONE cterm=NONE
hi StatusLine guifg=#1c2210 guibg=#a7784e gui=NONE cterm=NONE
hi StatusLineNC guifg=#4a3d2b guibg=#a7784e gui=NONE cterm=NONE
hi Pmenu guifg=#1c2210 guibg=#a7784e gui=NONE cterm=NONE
hi PmenuSel guifg=#b37f53 guibg=#5a2a4a gui=NONE cterm=NONE
hi Search guifg=#b37f53 guibg=#5a4410 gui=NONE cterm=NONE
hi IncSearch guifg=#b37f53 guibg=#5a2a4a gui=NONE cterm=NONE
hi MatchParen guifg=#24365a gui=bold cterm=bold
hi Directory guifg=#24365a gui=NONE cterm=NONE
hi Folded guifg=#4a3d2b guibg=#a7784e gui=NONE cterm=NONE
hi NonText guifg=#4a3d2b gui=NONE cterm=NONE
hi Whitespace guifg=#4a3d2b gui=NONE cterm=NONE
hi Conceal guifg=#4a3d2b gui=NONE cterm=NONE
hi Title guifg=#5a2a4a gui=bold cterm=bold
hi Comment guifg=#4a3d2b gui=italic cterm=italic
hi String guifg=#5a4410 gui=NONE cterm=NONE
hi Character guifg=#5a4410 gui=NONE cterm=NONE
hi Number guifg=#5a4410 gui=NONE cterm=NONE
hi Boolean guifg=#5a4410 gui=NONE cterm=NONE
hi Identifier guifg=#1c2210 gui=NONE cterm=NONE
hi Function guifg=#24365a gui=NONE cterm=NONE
hi Statement guifg=#5a2a4a gui=NONE cterm=NONE
hi Keyword guifg=#5a2a4a gui=NONE cterm=NONE
hi Operator guifg=#4a3d2b gui=NONE cterm=NONE
hi PreProc guifg=#2f3a1f gui=NONE cterm=NONE
hi Type guifg=#2f3a1f gui=NONE cterm=NONE
hi Constant guifg=#5a4410 gui=NONE cterm=NONE
hi Special guifg=#24365a gui=NONE cterm=NONE
hi Todo guifg=#b37f53 guibg=#5a4410 gui=bold cterm=bold
hi Error guifg=#6e2a1f gui=NONE cterm=NONE
hi markdownH1 guifg=#5a2a4a gui=bold cterm=bold
hi markdownH2 guifg=#5a2a4a gui=bold cterm=bold
hi markdownH3 guifg=#5a2a4a gui=NONE cterm=NONE
hi markdownH4 guifg=#5a2a4a gui=NONE cterm=NONE
hi markdownH5 guifg=#24365a gui=NONE cterm=NONE
hi markdownH6 guifg=#24365a gui=NONE cterm=NONE
hi markdownCode guifg=#2f3a1f gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#2f3a1f gui=NONE cterm=NONE
hi markdownLinkText guifg=#24365a gui=underline cterm=underline
hi markdownUrl guifg=#4a3d2b gui=NONE cterm=NONE
hi markdownListMarker guifg=#5a2a4a gui=NONE cterm=NONE
hi markdownRule guifg=#4a3d2b gui=NONE cterm=NONE
hi markdownBlockquote guifg=#4a3d2b gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#4a3d2b gui=NONE cterm=NONE
hi DiagnosticError guifg=#6e2a1f gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#5a4410 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#24365a gui=NONE cterm=NONE
hi DiagnosticHint guifg=#1f4a4a gui=NONE cterm=NONE
hi DiffAdd guifg=#2f3a1f gui=NONE cterm=NONE
hi DiffDelete guifg=#6e2a1f gui=NONE cterm=NONE
hi DiffChange guifg=#5a4410 gui=NONE cterm=NONE
hi DiffText guifg=#5a2a4a gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#2a2418', '#6e2a1f', '#2f3a1f', '#5a4410', '#24365a', '#5a2a4a', '#1f4a4a', '#b07a4e',
  \ '#4a3d2b', '#84362a', '#3b4826', '#6e5418', '#2f4470', '#6e3459', '#285c5c', '#cc976a']
