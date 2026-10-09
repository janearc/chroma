" pluto -- a light colourway, pluto as new horizons first showed it on
" 13 july 2015, the day before closest approach (nasa/jhuapl/swri,
" pia19708). the ground is the heart; the ink is the whale, darkened to 10
" to 1; the selection is the north polar cap, claude's band the midlands,
" dim text the disc. pluto has no red, green or blue to speak of, so the
" sixteen are the screen's own wavelengths, darkened to stand on the heart.
" samples in cmd/spectra/data/pluto-2015.css. made by spectra,
" github.com/janearc/chroma/cmd/spectra, 2026-10-05.
"
" rendered by colourway from sources/pluto.css;
" change the source, not this file.
"
"   :colorscheme pluto
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
let g:colors_name = 'pluto'

hi Normal guifg=#2a1b10 guibg=#e9c29d gui=NONE cterm=NONE
hi NormalFloat guifg=#2a1b10 guibg=#dab592 gui=NONE cterm=NONE
hi FloatBorder guifg=#635040 guibg=#dab592 gui=NONE cterm=NONE
hi CursorLine guibg=#dcb793 gui=NONE cterm=NONE
hi CursorLineNr guifg=#7a0079 gui=bold cterm=bold
hi LineNr guifg=#635040 gui=NONE cterm=NONE
hi SignColumn guibg=#e9c29d gui=NONE cterm=NONE
hi Visual guibg=#d6ae88 gui=NONE cterm=NONE
hi VertSplit guifg=#635040 gui=NONE cterm=NONE
hi WinSeparator guifg=#635040 gui=NONE cterm=NONE
hi StatusLine guifg=#2a1b10 guibg=#dab592 gui=NONE cterm=NONE
hi StatusLineNC guifg=#635040 guibg=#dab592 gui=NONE cterm=NONE
hi Pmenu guifg=#2a1b10 guibg=#dab592 gui=NONE cterm=NONE
hi PmenuSel guifg=#e9c29d guibg=#7a0079 gui=NONE cterm=NONE
hi Search guifg=#e9c29d guibg=#444500 gui=NONE cterm=NONE
hi IncSearch guifg=#e9c29d guibg=#7a0079 gui=NONE cterm=NONE
hi MatchParen guifg=#004282 gui=bold cterm=bold
hi Directory guifg=#004282 gui=NONE cterm=NONE
hi Folded guifg=#635040 guibg=#dab592 gui=NONE cterm=NONE
hi NonText guifg=#635040 gui=NONE cterm=NONE
hi Whitespace guifg=#635040 gui=NONE cterm=NONE
hi Conceal guifg=#635040 gui=NONE cterm=NONE
hi Title guifg=#7a0079 gui=bold cterm=bold
hi Comment guifg=#635040 gui=italic cterm=italic
hi String guifg=#444500 gui=NONE cterm=NONE
hi Character guifg=#444500 gui=NONE cterm=NONE
hi Number guifg=#444500 gui=NONE cterm=NONE
hi Boolean guifg=#444500 gui=NONE cterm=NONE
hi Identifier guifg=#2a1b10 gui=NONE cterm=NONE
hi Function guifg=#004282 gui=NONE cterm=NONE
hi Statement guifg=#7a0079 gui=NONE cterm=NONE
hi Keyword guifg=#7a0079 gui=NONE cterm=NONE
hi Operator guifg=#635040 gui=NONE cterm=NONE
hi PreProc guifg=#074e00 gui=NONE cterm=NONE
hi Type guifg=#074e00 gui=NONE cterm=NONE
hi Constant guifg=#444500 gui=NONE cterm=NONE
hi Special guifg=#004282 gui=NONE cterm=NONE
hi Todo guifg=#e9c29d guibg=#444500 gui=bold cterm=bold
hi Error guifg=#880e00 gui=NONE cterm=NONE
hi markdownH1 guifg=#7a0079 gui=bold cterm=bold
hi markdownH2 guifg=#7a0079 gui=bold cterm=bold
hi markdownH3 guifg=#7a0079 gui=NONE cterm=NONE
hi markdownH4 guifg=#7a0079 gui=NONE cterm=NONE
hi markdownH5 guifg=#004282 gui=NONE cterm=NONE
hi markdownH6 guifg=#004282 gui=NONE cterm=NONE
hi markdownCode guifg=#074e00 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#074e00 gui=NONE cterm=NONE
hi markdownLinkText guifg=#004282 gui=underline cterm=underline
hi markdownUrl guifg=#635040 gui=NONE cterm=NONE
hi markdownListMarker guifg=#7a0079 gui=NONE cterm=NONE
hi markdownRule guifg=#635040 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#635040 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#635040 gui=NONE cterm=NONE
hi DiagnosticError guifg=#880e00 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#444500 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#004282 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#004b49 gui=NONE cterm=NONE
hi DiffAdd guifg=#074e00 gui=NONE cterm=NONE
hi DiffDelete guifg=#880e00 gui=NONE cterm=NONE
hi DiffChange guifg=#444500 gui=NONE cterm=NONE
hi DiffText guifg=#7a0079 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#d5b69a', '#880e00', '#074e00', '#444500', '#004282', '#7a0079', '#004b49', '#2f251d',
  \ '#635040', '#5c0600', '#033300', '#2c2d00', '#002b59', '#530052', '#00312f', '#0f0703']
