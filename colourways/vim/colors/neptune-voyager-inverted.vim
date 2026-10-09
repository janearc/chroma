" neptune-voyager -- neptune as voyager 2 showed it in august 1989, from
" the green and orange filters of its narrow angle camera (nasa/jpl,
" pia01492). the ground is the limb in shadow, greyed to her hue rule; the
" ink is the white clouds, lit to 11 to 1; the selection and the blues are
" the disc; dim text is the great dark spot. neptune has no red, green or
" yellow, so those are the screen's own wavelengths, as in fermi-long.
" samples in cmd/spectra/data/neptune-voyager.css. made by spectra,
" github.com/janearc/chroma/cmd/spectra, 2026-10-05.
"
" inverted, for a screen whose colours the system inverts, so it shows as meant.
"
" rendered by colourway from sources/neptune-voyager.css;
" change the source, not this file.
"
"   :colorscheme neptune-voyager-inverted
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
let g:colors_name = 'neptune-voyager-inverted'

hi Normal guifg=#653000 guibg=#eeeadc gui=NONE cterm=NONE
hi NormalFloat guifg=#653000 guibg=#e4dccc gui=NONE cterm=NONE
hi FloatBorder guifg=#ab8b00 guibg=#e4dccc gui=NONE cterm=NONE
hi CursorLine guibg=#e6dece gui=NONE cterm=NONE
hi CursorLineNr guifg=#00d606 gui=bold cterm=bold
hi LineNr guifg=#ab8b00 gui=NONE cterm=NONE
hi SignColumn guibg=#eeeadc gui=NONE cterm=NONE
hi Visual guibg=#dfcb7f gui=NONE cterm=NONE
hi VertSplit guifg=#ab8b00 gui=NONE cterm=NONE
hi WinSeparator guifg=#ab8b00 gui=NONE cterm=NONE
hi StatusLine guifg=#653000 guibg=#e4dccc gui=NONE cterm=NONE
hi StatusLineNC guifg=#ab8b00 guibg=#e4dccc gui=NONE cterm=NONE
hi Pmenu guifg=#653000 guibg=#e4dccc gui=NONE cterm=NONE
hi PmenuSel guifg=#eeeadc guibg=#00d606 gui=NONE cterm=NONE
hi Search guifg=#eeeadc guibg=#6865ff gui=NONE cterm=NONE
hi IncSearch guifg=#eeeadc guibg=#00d606 gui=NONE cterm=NONE
hi MatchParen guifg=#957000 gui=bold cterm=bold
hi Directory guifg=#957000 gui=NONE cterm=NONE
hi Folded guifg=#ab8b00 guibg=#e4dccc gui=NONE cterm=NONE
hi NonText guifg=#ab8b00 gui=NONE cterm=NONE
hi Whitespace guifg=#ab8b00 gui=NONE cterm=NONE
hi Conceal guifg=#ab8b00 gui=NONE cterm=NONE
hi Title guifg=#00d606 gui=bold cterm=bold
hi Comment guifg=#ab8b00 gui=italic cterm=italic
hi String guifg=#6865ff gui=NONE cterm=NONE
hi Character guifg=#6865ff gui=NONE cterm=NONE
hi Number guifg=#6865ff gui=NONE cterm=NONE
hi Boolean guifg=#6865ff gui=NONE cterm=NONE
hi Identifier guifg=#653000 gui=NONE cterm=NONE
hi Function guifg=#957000 gui=NONE cterm=NONE
hi Statement guifg=#00d606 gui=NONE cterm=NONE
hi Keyword guifg=#00d606 gui=NONE cterm=NONE
hi Operator guifg=#ab8b00 gui=NONE cterm=NONE
hi PreProc guifg=#e353ff gui=NONE cterm=NONE
hi Type guifg=#e353ff gui=NONE cterm=NONE
hi Constant guifg=#6865ff gui=NONE cterm=NONE
hi Special guifg=#957000 gui=NONE cterm=NONE
hi Todo guifg=#eeeadc guibg=#6865ff gui=bold cterm=bold
hi Error guifg=#00a1b9 gui=NONE cterm=NONE
hi markdownH1 guifg=#00d606 gui=bold cterm=bold
hi markdownH2 guifg=#00d606 gui=bold cterm=bold
hi markdownH3 guifg=#00d606 gui=NONE cterm=NONE
hi markdownH4 guifg=#00d606 gui=NONE cterm=NONE
hi markdownH5 guifg=#957000 gui=NONE cterm=NONE
hi markdownH6 guifg=#957000 gui=NONE cterm=NONE
hi markdownCode guifg=#e353ff gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#e353ff gui=NONE cterm=NONE
hi markdownLinkText guifg=#957000 gui=underline cterm=underline
hi markdownUrl guifg=#ab8b00 gui=NONE cterm=NONE
hi markdownListMarker guifg=#00d606 gui=NONE cterm=NONE
hi markdownRule guifg=#ab8b00 gui=NONE cterm=NONE
hi markdownBlockquote guifg=#ab8b00 gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#ab8b00 gui=NONE cterm=NONE
hi DiagnosticError guifg=#00a1b9 gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#6865ff gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#957000 gui=NONE cterm=NONE
hi DiagnosticHint guifg=#a4652d gui=NONE cterm=NONE
hi DiffAdd guifg=#e353ff gui=NONE cterm=NONE
hi DiffDelete guifg=#00a1b9 gui=NONE cterm=NONE
hi DiffChange guifg=#6865ff gui=NONE cterm=NONE
hi DiffText guifg=#00d606 gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#ddd6c1', '#00a1b9', '#e353ff', '#6865ff', '#957000', '#00d606', '#a4652d', '#8d4200',
  \ '#ab8b00', '#006a7e', '#db31ff', '#4b46ff', '#6d4f00', '#007b0d', '#924705', '#532800']
