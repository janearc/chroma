" twilight-deep -- twilight's ground taken darker: ink #38241e on ground
" #808a63, 4.0:1, chosen in paratune, with claude code's own colours tuned
" to match. the palette is darker so the accents still read on the deeper
" ground, and the statusline's bar is twilight's ramp moved by the same
" shift.
"
" rendered by colourway from sources/twilight-deep.css;
" change the source, not this file.
"
"   :colorscheme twilight-deep
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
let g:colors_name = 'twilight-deep'

hi Normal guifg=#38241e guibg=#808a63 gui=NONE cterm=NONE
hi NormalFloat guifg=#38241e guibg=#8f8856 gui=NONE cterm=NONE
hi FloatBorder guifg=#6e683e guibg=#8f8856 gui=NONE cterm=NONE
hi CursorLine guibg=#8c8553 gui=NONE cterm=NONE
hi CursorLineNr guifg=#3a1830 gui=bold cterm=bold
hi LineNr guifg=#6e683e gui=NONE cterm=NONE
hi SignColumn guibg=#808a63 gui=NONE cterm=NONE
hi Visual guibg=#7f794b gui=NONE cterm=NONE
hi VertSplit guifg=#6e683e gui=NONE cterm=NONE
hi WinSeparator guifg=#6e683e gui=NONE cterm=NONE
hi StatusLine guifg=#38241e guibg=#8f8856 gui=NONE cterm=NONE
hi StatusLineNC guifg=#6e683e guibg=#8f8856 gui=NONE cterm=NONE
hi Pmenu guifg=#38241e guibg=#8f8856 gui=NONE cterm=NONE
hi PmenuSel guifg=#808a63 guibg=#3a1830 gui=NONE cterm=NONE
hi Search guifg=#808a63 guibg=#3a2c08 gui=NONE cterm=NONE
hi IncSearch guifg=#808a63 guibg=#3a1830 gui=NONE cterm=NONE
hi MatchParen guifg=#141c3a gui=bold cterm=bold
hi Directory guifg=#141c3a gui=NONE cterm=NONE
hi Folded guifg=#2e2a1c guibg=#8f8856 gui=NONE cterm=NONE
hi NonText guifg=#6e683e gui=NONE cterm=NONE
hi Whitespace guifg=#6e683e gui=NONE cterm=NONE
hi Conceal guifg=#6e683e gui=NONE cterm=NONE
hi Title guifg=#3a1830 gui=bold cterm=bold
hi Comment guifg=#2e2a1c gui=italic cterm=italic
hi String guifg=#3a2c08 gui=NONE cterm=NONE
hi Character guifg=#3a2c08 gui=NONE cterm=NONE
hi Number guifg=#3a2c08 gui=NONE cterm=NONE
hi Boolean guifg=#3a2c08 gui=NONE cterm=NONE
hi Identifier guifg=#38241e gui=NONE cterm=NONE
hi Function guifg=#141c3a gui=NONE cterm=NONE
hi Statement guifg=#3a1830 gui=NONE cterm=NONE
hi Keyword guifg=#3a1830 gui=NONE cterm=NONE
hi Operator guifg=#2e2a1c gui=NONE cterm=NONE
hi PreProc guifg=#1c2810 gui=NONE cterm=NONE
hi Type guifg=#1c2810 gui=NONE cterm=NONE
hi Constant guifg=#3a2c08 gui=NONE cterm=NONE
hi Special guifg=#141c3a gui=NONE cterm=NONE
hi Todo guifg=#808a63 guibg=#3a2c08 gui=bold cterm=bold
hi Error guifg=#4a1810 gui=NONE cterm=NONE
hi markdownH1 guifg=#3a1830 gui=bold cterm=bold
hi markdownH2 guifg=#3a1830 gui=bold cterm=bold
hi markdownH3 guifg=#3a1830 gui=NONE cterm=NONE
hi markdownH4 guifg=#3a1830 gui=NONE cterm=NONE
hi markdownH5 guifg=#141c3a gui=NONE cterm=NONE
hi markdownH6 guifg=#141c3a gui=NONE cterm=NONE
hi markdownCode guifg=#1c2810 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#1c2810 gui=NONE cterm=NONE
hi markdownLinkText guifg=#141c3a gui=underline cterm=underline
hi markdownUrl guifg=#2e2a1c gui=NONE cterm=NONE
hi markdownListMarker guifg=#3a1830 gui=NONE cterm=NONE
hi markdownRule guifg=#6e683e gui=NONE cterm=NONE
hi markdownBlockquote guifg=#2e2a1c gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#6e683e gui=NONE cterm=NONE
hi DiagnosticError guifg=#4a1810 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#3a2c08 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#141c3a gui=NONE cterm=NONE
hi DiagnosticHint guifg=#1c2810 gui=NONE cterm=NONE
hi DiffAdd guifg=#1c2810 gui=NONE cterm=NONE
hi DiffDelete guifg=#4a1810 gui=NONE cterm=NONE
hi DiffChange guifg=#3a2c08 gui=NONE cterm=NONE
hi DiffText guifg=#3a1830 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#1c1a10', '#4a1810', '#1c2810', '#3a2c08', '#141c3a', '#3a1830', '#10302e', '#7f794b',
  \ '#2e2a1c', '#5a2015', '#263214', '#4a3a0c', '#1c2850', '#4a2040', '#163c3a', '#a9a269']
