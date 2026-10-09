" aqua -- a pale aqua ground with slate blue ink, 6.0:1, chosen in paratune.
" a bright ground narrows the pupil, and a narrow pupil is a lens stopped
" down: the letters land sharper. cool, not white, so it does not sting.
" accents are dark and cool so they read on the bright ground; cells 17 to
" 254 are the status-line bar's ramp, as tints of the ground.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/aqua.css;
" change the source, not this file.
"
"   :colorscheme aqua-inverted
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
let g:colors_name = 'aqua-inverted'

hi Normal guifg=#afa77f guibg=#5c0000 gui=NONE cterm=NONE
hi NormalFloat guifg=#afa77f guibg=#631208 gui=NONE cterm=NONE
hi FloatBorder guifg=#b59f8f guibg=#631208 gui=NONE cterm=NONE
hi CursorLine guibg=#621006 gui=NONE cterm=NONE
hi CursorLineNr guifg=#95d585 gui=bold cterm=bold
hi LineNr guifg=#b59f8f gui=NONE cterm=NONE
hi SignColumn guibg=#5c0000 gui=NONE cterm=NONE
hi Visual guibg=#801f19 gui=NONE cterm=NONE
hi VertSplit guifg=#b59f8f gui=NONE cterm=NONE
hi WinSeparator guifg=#b59f8f gui=NONE cterm=NONE
hi StatusLine guifg=#afa77f guibg=#631208 gui=NONE cterm=NONE
hi StatusLineNC guifg=#b59f8f guibg=#631208 gui=NONE cterm=NONE
hi Pmenu guifg=#afa77f guibg=#631208 gui=NONE cterm=NONE
hi PmenuSel guifg=#5c0000 guibg=#95d585 gui=NONE cterm=NONE
hi Search guifg=#5c0000 guibg=#95adef gui=NONE cterm=NONE
hi IncSearch guifg=#5c0000 guibg=#95d585 gui=NONE cterm=NONE
hi MatchParen guifg=#e0c575 gui=bold cterm=bold
hi Directory guifg=#e0c575 gui=NONE cterm=NONE
hi Folded guifg=#b59f8f guibg=#631208 gui=NONE cterm=NONE
hi NonText guifg=#b59f8f gui=NONE cterm=NONE
hi Whitespace guifg=#b59f8f gui=NONE cterm=NONE
hi Conceal guifg=#b59f8f gui=NONE cterm=NONE
hi Title guifg=#95d585 gui=bold cterm=bold
hi Comment guifg=#b59f8f gui=italic cterm=italic
hi String guifg=#95adef gui=NONE cterm=NONE
hi Character guifg=#95adef gui=NONE cterm=NONE
hi Number guifg=#95adef gui=NONE cterm=NONE
hi Boolean guifg=#95adef gui=NONE cterm=NONE
hi Identifier guifg=#afa77f gui=NONE cterm=NONE
hi Function guifg=#e0c575 gui=NONE cterm=NONE
hi Statement guifg=#95d585 gui=NONE cterm=NONE
hi Keyword guifg=#95d585 gui=NONE cterm=NONE
hi Operator guifg=#b59f8f gui=NONE cterm=NONE
hi PreProc guifg=#e0a5c5 gui=NONE cterm=NONE
hi Type guifg=#e0a5c5 gui=NONE cterm=NONE
hi Constant guifg=#95adef gui=NONE cterm=NONE
hi Special guifg=#e0c575 gui=NONE cterm=NONE
hi Todo guifg=#5c0000 guibg=#95adef gui=bold cterm=bold
hi Error guifg=#75d5c5 gui=NONE cterm=NONE
hi markdownH1 guifg=#95d585 gui=bold cterm=bold
hi markdownH2 guifg=#95d585 gui=bold cterm=bold
hi markdownH3 guifg=#95d585 gui=NONE cterm=NONE
hi markdownH4 guifg=#95d585 gui=NONE cterm=NONE
hi markdownH5 guifg=#e0c575 gui=NONE cterm=NONE
hi markdownH6 guifg=#e0c575 gui=NONE cterm=NONE
hi markdownCode guifg=#e0a5c5 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#e0a5c5 gui=NONE cterm=NONE
hi markdownLinkText guifg=#e0c575 gui=underline cterm=underline
hi markdownUrl guifg=#b59f8f gui=NONE cterm=NONE
hi markdownListMarker guifg=#95d585 gui=NONE cterm=NONE
hi markdownRule guifg=#b59f8f gui=NONE cterm=NONE
hi markdownBlockquote guifg=#b59f8f gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#b59f8f gui=NONE cterm=NONE
hi DiagnosticError guifg=#75d5c5 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#95adef gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#e0c575 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#f5a595 gui=NONE cterm=NONE
hi DiffAdd guifg=#e0a5c5 gui=NONE cterm=NONE
hi DiffDelete guifg=#75d5c5 gui=NONE cterm=NONE
hi DiffChange guifg=#95adef gui=NONE cterm=NONE
hi DiffText guifg=#95d585 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#e1d5c7', '#75d5c5', '#e0a5c5', '#95adef', '#e0c575', '#95d585', '#f5a595', '#701919',
  \ '#b59f8f', '#5fcab7', '#d595b9', '#859fe7', '#d5b55f', '#85c575', '#eb9585', '#470000']
