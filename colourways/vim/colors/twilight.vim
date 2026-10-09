" twilight -- a ground that is not dark and letters that are not bright,
" so the ground cannot shine up through the text in rivers: an amber
" ground with deep navy ink, 8.9:1, chosen in paratune. dark text on a
" lightish ground.
"
" rendered by colourway from sources/twilight.css;
" change the source, not this file.
"
"   :colorscheme twilight
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
let g:colors_name = 'twilight'

hi Normal guifg=#051759 guibg=#ffac4e gui=NONE cterm=NONE
hi NormalFloat guifg=#051759 guibg=#f6a64b gui=NONE cterm=NONE
hi FloatBorder guifg=#c4824a guibg=#f6a64b gui=NONE cterm=NONE
hi CursorLine guibg=#f0a249 gui=NONE cterm=NONE
hi CursorLineNr guifg=#5a2a4a gui=bold cterm=bold
hi LineNr guifg=#b87643 gui=NONE cterm=NONE
hi SignColumn guibg=#ffac4e gui=NONE cterm=NONE
hi Visual guibg=#e8963a gui=NONE cterm=NONE
hi VertSplit guifg=#c4824a gui=NONE cterm=NONE
hi WinSeparator guifg=#c4824a gui=NONE cterm=NONE
hi StatusLine guifg=#051759 guibg=#f6a64b gui=NONE cterm=NONE
hi StatusLineNC guifg=#b87643 guibg=#f6a64b gui=NONE cterm=NONE
hi Pmenu guifg=#051759 guibg=#f6a64b gui=NONE cterm=NONE
hi PmenuSel guifg=#ffac4e guibg=#5a2a4a gui=NONE cterm=NONE
hi Search guifg=#ffac4e guibg=#5a4410 gui=NONE cterm=NONE
hi IncSearch guifg=#ffac4e guibg=#5a2a4a gui=NONE cterm=NONE
hi MatchParen guifg=#24365a gui=bold cterm=bold
hi Directory guifg=#24365a gui=NONE cterm=NONE
hi Folded guifg=#4a3d2b guibg=#f6a64b gui=NONE cterm=NONE
hi NonText guifg=#b87643 gui=NONE cterm=NONE
hi Whitespace guifg=#b87643 gui=NONE cterm=NONE
hi Conceal guifg=#b87643 gui=NONE cterm=NONE
hi Title guifg=#5a2a4a gui=bold cterm=bold
hi Comment guifg=#4a3d2b gui=italic cterm=italic
hi String guifg=#5a4410 gui=NONE cterm=NONE
hi Character guifg=#5a4410 gui=NONE cterm=NONE
hi Number guifg=#5a4410 gui=NONE cterm=NONE
hi Boolean guifg=#5a4410 gui=NONE cterm=NONE
hi Identifier guifg=#051759 gui=NONE cterm=NONE
hi Function guifg=#24365a gui=NONE cterm=NONE
hi Statement guifg=#5a2a4a gui=NONE cterm=NONE
hi Keyword guifg=#5a2a4a gui=NONE cterm=NONE
hi Operator guifg=#4a3d2b gui=NONE cterm=NONE
hi PreProc guifg=#2f3a1f gui=NONE cterm=NONE
hi Type guifg=#2f3a1f gui=NONE cterm=NONE
hi Constant guifg=#5a4410 gui=NONE cterm=NONE
hi Special guifg=#24365a gui=NONE cterm=NONE
hi Todo guifg=#ffac4e guibg=#5a4410 gui=bold cterm=bold
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
hi markdownRule guifg=#c4824a gui=NONE cterm=NONE
hi markdownBlockquote guifg=#4a3d2b gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#b87643 gui=NONE cterm=NONE
hi DiagnosticError guifg=#6e2a1f gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#5a4410 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#24365a gui=NONE cterm=NONE
hi DiagnosticHint guifg=#2f3a1f gui=NONE cterm=NONE
hi DiffAdd guifg=#2f3a1f gui=NONE cterm=NONE
hi DiffDelete guifg=#6e2a1f gui=NONE cterm=NONE
hi DiffChange guifg=#5a4410 gui=NONE cterm=NONE
hi DiffText guifg=#5a2a4a gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#2a2418', '#6e2a1f', '#2f3a1f', '#5a4410', '#24365a', '#5a2a4a', '#1f4a4a', '#ec9c42',
  \ '#4a3d2b', '#84362a', '#3b4826', '#6e5418', '#2f4470', '#6e3459', '#285c5c', '#ffbd6e']
