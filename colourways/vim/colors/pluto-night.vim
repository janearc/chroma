" pluto-night -- a dark colourway, made in paratune from pluto.
"
" rendered by colourway from sources/pluto-night.css;
" change the source, not this file.
"
"   :colorscheme pluto-night
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
let g:colors_name = 'pluto-night'

hi Normal guifg=#bda79e guibg=#0b1134 gui=NONE cterm=NONE
hi NormalFloat guifg=#bda79e guibg=#151b3c gui=NONE cterm=NONE
hi FloatBorder guifg=#9a605b guibg=#151b3c gui=NONE cterm=NONE
hi CursorLine guibg=#13193a gui=NONE cterm=NONE
hi CursorLineNr guifg=#ff8cff gui=bold cterm=bold
hi LineNr guifg=#4668aa gui=NONE cterm=NONE
hi SignColumn guibg=#0b1134 gui=NONE cterm=NONE
hi Visual guibg=#060606 gui=NONE cterm=NONE
hi VertSplit guifg=#9a605b gui=NONE cterm=NONE
hi WinSeparator guifg=#9a605b gui=NONE cterm=NONE
hi StatusLine guifg=#bda79e guibg=#151b3c gui=NONE cterm=NONE
hi StatusLineNC guifg=#4668aa guibg=#151b3c gui=NONE cterm=NONE
hi Pmenu guifg=#bda79e guibg=#151b3c gui=NONE cterm=NONE
hi PmenuSel guifg=#0b1134 guibg=#ff8cff gui=NONE cterm=NONE
hi Search guifg=#0b1134 guibg=#6b8307 gui=NONE cterm=NONE
hi IncSearch guifg=#0b1134 guibg=#ff8cff gui=NONE cterm=NONE
hi MatchParen guifg=#3783ff gui=bold cterm=bold
hi Directory guifg=#3783ff gui=NONE cterm=NONE
hi Folded guifg=#9a605b guibg=#151b3c gui=NONE cterm=NONE
hi NonText guifg=#4668aa gui=NONE cterm=NONE
hi Whitespace guifg=#4668aa gui=NONE cterm=NONE
hi Conceal guifg=#4668aa gui=NONE cterm=NONE
hi Title guifg=#ff8cff gui=bold cterm=bold
hi Comment guifg=#9a605b gui=italic cterm=italic
hi String guifg=#6b8307 gui=NONE cterm=NONE
hi Character guifg=#6b8307 gui=NONE cterm=NONE
hi Number guifg=#6b8307 gui=NONE cterm=NONE
hi Boolean guifg=#6b8307 gui=NONE cterm=NONE
hi Identifier guifg=#bda79e gui=NONE cterm=NONE
hi Function guifg=#3783ff gui=NONE cterm=NONE
hi Statement guifg=#ff8cff gui=NONE cterm=NONE
hi Keyword guifg=#ff8cff gui=NONE cterm=NONE
hi Operator guifg=#9a605b gui=NONE cterm=NONE
hi PreProc guifg=#74d474 gui=NONE cterm=NONE
hi Type guifg=#74d474 gui=NONE cterm=NONE
hi Constant guifg=#6b8307 gui=NONE cterm=NONE
hi Special guifg=#3783ff gui=NONE cterm=NONE
hi Todo guifg=#0b1134 guibg=#6b8307 gui=bold cterm=bold
hi Error guifg=#d60b0d gui=NONE cterm=NONE
hi markdownH1 guifg=#ff8cff gui=bold cterm=bold
hi markdownH2 guifg=#ff8cff gui=bold cterm=bold
hi markdownH3 guifg=#ff8cff gui=NONE cterm=NONE
hi markdownH4 guifg=#ff8cff gui=NONE cterm=NONE
hi markdownH5 guifg=#3783ff gui=NONE cterm=NONE
hi markdownH6 guifg=#3783ff gui=NONE cterm=NONE
hi markdownCode guifg=#74d474 gui=NONE cterm=NONE
hi markdownCodeBlock guifg=#74d474 gui=NONE cterm=NONE
hi markdownLinkText guifg=#3783ff gui=underline cterm=underline
hi markdownUrl guifg=#9a605b gui=NONE cterm=NONE
hi markdownListMarker guifg=#ff8cff gui=NONE cterm=NONE
hi markdownRule guifg=#9a605b gui=NONE cterm=NONE
hi markdownBlockquote guifg=#9a605b gui=italic cterm=italic
hi markdownHeadingDelimiter guifg=#4668aa gui=NONE cterm=NONE
hi DiagnosticError guifg=#d60b0d gui=NONE cterm=NONE
hi DiagnosticWarn guifg=#88a709 gui=NONE cterm=NONE
hi DiagnosticInfo guifg=#3783ff gui=NONE cterm=NONE
hi DiagnosticHint guifg=#5bb65b gui=NONE cterm=NONE
hi DiffAdd guifg=#74d474 gui=NONE cterm=NONE
hi DiffDelete guifg=#d60b0d gui=NONE cterm=NONE
hi DiffChange guifg=#6b8307 gui=NONE cterm=NONE
hi DiffText guifg=#ff8cff gui=bold cterm=bold

let g:terminal_ansi_colors = [
  \ '#1e2d50', '#d60b0d', '#74d474', '#88a709', '#3783ff', '#ff8cff', '#c5ffff', '#ff825f',
  \ '#cc6045', '#ff0909', '#398e39', '#818100', '#2076ff', '#c942c9', '#628383', '#c95c41']
