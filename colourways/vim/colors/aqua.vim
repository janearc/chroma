" aqua -- a pale aqua ground with slate blue ink, 6.0:1, chosen in paratune.
" a bright ground narrows the pupil, and a narrow pupil is a lens stopped
" down: the letters land sharper. cool, not white, so it does not sting.
" accents are dark and cool so they read on the bright ground; cells 17 to
" 254 are the status-line bar's ramp, as tints of the ground.
"
" rendered by colourway from sources/aqua.css;
" change the source, not this file.
"
"   :colorscheme aqua
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
let g:colors_name = 'aqua'

hi Normal guifg=#505880 guibg=#a3ffff gui=NONE cterm=NONE
hi NormalFloat guifg=#505880 guibg=#9df2f6 gui=NONE cterm=NONE
hi FloatBorder guifg=#4a6070 guibg=#9df2f6 gui=NONE cterm=NONE
hi CursorLine guibg=#9ef4f7 gui=NONE cterm=NONE
hi CursorLineNr guifg=#6a2a7a gui=bold cterm=bold
hi LineNr guifg=#4a6070 gui=NONE cterm=NONE
hi SignColumn guibg=#a3ffff gui=NONE cterm=NONE
hi Visual guibg=#7fe0e6 gui=NONE cterm=NONE
hi VertSplit guifg=#4a6070 gui=NONE cterm=NONE
hi WinSeparator guifg=#4a6070 gui=NONE cterm=NONE
hi StatusLine guifg=#505880 guibg=#9df2f6 gui=NONE cterm=NONE
hi StatusLineNC guifg=#4a6070 guibg=#9df2f6 gui=NONE cterm=NONE
hi Pmenu guifg=#505880 guibg=#9df2f6 gui=NONE cterm=NONE
hi PmenuSel guifg=#a3ffff guibg=#6a2a7a gui=NONE cterm=NONE
hi Search guifg=#a3ffff guibg=#6a5210 gui=NONE cterm=NONE
hi IncSearch guifg=#a3ffff guibg=#6a2a7a gui=NONE cterm=NONE
hi MatchParen guifg=#1f3a8a gui=bold cterm=bold
hi Directory guifg=#1f3a8a gui=NONE cterm=NONE
hi Folded guifg=#4a6070 guibg=#9df2f6 gui=NONE cterm=NONE
hi NonText guifg=#4a6070 gui=NONE cterm=NONE
hi Whitespace guifg=#4a6070 gui=NONE cterm=NONE
hi Conceal guifg=#4a6070 gui=NONE cterm=NONE
hi Title guifg=#6a2a7a gui=bold cterm=bold
hi Comment guifg=#4a6070 gui=italic cterm=italic
hi String guifg=#6a5210 gui=NONE cterm=NONE
hi Character guifg=#6a5210 gui=NONE cterm=NONE
hi Number guifg=#6a5210 gui=NONE cterm=NONE
hi Boolean guifg=#6a5210 gui=NONE cterm=NONE
hi Identifier guifg=#505880 gui=NONE cterm=NONE
hi Function guifg=#1f3a8a gui=NONE cterm=NONE
hi Statement guifg=#6a2a7a gui=NONE cterm=NONE
hi Keyword guifg=#6a2a7a gui=NONE cterm=NONE
hi Operator guifg=#4a6070 gui=NONE cterm=NONE
hi PreProc guifg=#1f5a3a gui=NONE cterm=NONE
hi Type guifg=#1f5a3a gui=NONE cterm=NONE
hi Constant guifg=#6a5210 gui=NONE cterm=NONE
hi Special guifg=#1f3a8a gui=NONE cterm=NONE
hi Todo guifg=#a3ffff guibg=#6a5210 gui=bold cterm=bold
hi Error guifg=#8a2a3a gui=NONE cterm=NONE
hi markdownH1 guifg=#6a2a7a gui=bold cterm=bold
hi markdownH2 guifg=#6a2a7a gui=bold cterm=bold
hi markdownH3 guifg=#6a2a7a gui=NONE cterm=NONE
hi markdownH4 guifg=#6a2a7a gui=NONE cterm=NONE
hi markdownH5 guifg=#1f3a8a gui=NONE cterm=NONE
hi markdownH6 guifg=#1f3a8a gui=NONE cterm=NONE
hi markdownCode guifg=#1f5a3a gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#1f5a3a gui=NONE cterm=NONE
hi markdownLinkText guifg=#1f3a8a gui=underline cterm=underline
hi markdownUrl guifg=#4a6070 gui=NONE cterm=NONE
hi markdownListMarker guifg=#6a2a7a gui=NONE cterm=NONE
hi markdownRule guifg=#4a6070 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#4a6070 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#4a6070 gui=NONE cterm=NONE
hi DiagnosticError guifg=#8a2a3a gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#6a5210 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#1f3a8a gui=NONE cterm=NONE
hi DiagnosticHint guifg=#0a5a6a gui=NONE cterm=NONE
hi DiffAdd guifg=#1f5a3a gui=NONE cterm=NONE
hi DiffDelete guifg=#8a2a3a gui=NONE cterm=NONE
hi DiffChange guifg=#6a5210 gui=NONE cterm=NONE
hi DiffText guifg=#6a2a7a gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#1e2a38', '#8a2a3a', '#1f5a3a', '#6a5210', '#1f3a8a', '#6a2a7a', '#0a5a6a', '#8fe6e6',
  \ '#4a6070', '#a03548', '#2a6a46', '#7a6018', '#2a4aa0', '#7a3a8a', '#146a7a', '#b8ffff']
