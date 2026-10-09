-- twilight-low -- made in paratune from twilight.
--
-- inverted, for a screen whose colours the system inverts, so it shows as meant.
--
-- rendered by colourway from sources/twilight-low.css;
-- change the source, not this file.
--
--   :colorscheme twilight-low-inverted

local g = vim.api.nvim_set_hl

vim.cmd("highlight clear")
if vim.fn.exists("syntax_on") == 1 then vim.cmd("syntax reset") end
vim.o.background = "light"
vim.g.colors_name = "twilight-low-inverted"

local c = {
  bg       = "#14c0ed", -- the ground
  ink      = "#392100", --  7.04:1  body prose
  dim      = "#2b5dc6", --  2.81:1  comments
  heading  = "#00a22c", --  1.58:1  headings and keywords
  link     = "#835c00", --  2.80:1  links and functions
  code     = "#804cc8", --  2.59:1  code and types
  string   = "#2551c4", --  3.21:1  strings and numbers
  gutter   = "#000000", --  9.79:1  line numbers, furniture you look past
  cursorln = "#00315d", --  6.13:1  the cursor's line
  visual   = "#0085e5", --  1.78:1  the selection
  panel    = "#002c4e", --  6.66:1  floats and the status line
  border   = "#000000", --  9.79:1  borders and rules
  err      = "#00747a", --  2.59:1  errors
  warn     = "#2551c4", --  3.21:1  warnings
  info     = "#835c00", --  2.80:1  notes
  hint     = "#804cc8", --  2.59:1  hints
}


-- editor furniture
g(0, "Normal",        { fg = c.ink, bg = c.bg })
g(0, "NormalFloat",   { fg = c.ink, bg = c.panel })
g(0, "FloatBorder",   { fg = c.border, bg = c.panel })
g(0, "CursorLine",    { bg = c.cursorln })
g(0, "CursorLineNr",  { fg = c.heading, bold = true })
g(0, "LineNr",        { fg = c.gutter })
g(0, "SignColumn",    { bg = c.bg })
g(0, "Visual",        { bg = c.visual })
g(0, "VertSplit",     { fg = c.border })
g(0, "WinSeparator",  { fg = c.border })
g(0, "StatusLine",    { fg = c.ink, bg = c.panel })
g(0, "StatusLineNC",  { fg = c.gutter, bg = c.panel })
g(0, "Pmenu",         { fg = c.ink, bg = c.panel })
g(0, "PmenuSel",      { fg = c.bg, bg = c.heading })
g(0, "Search",        { fg = c.bg, bg = c.string })
g(0, "IncSearch",     { fg = c.bg, bg = c.heading })
g(0, "MatchParen",    { fg = c.link, bold = true })
g(0, "Directory",     { fg = c.link })
g(0, "Folded",        { fg = c.dim, bg = c.panel })
g(0, "NonText",       { fg = c.gutter })
g(0, "Whitespace",    { fg = c.gutter })
g(0, "Conceal",       { fg = c.gutter })
g(0, "Title",         { fg = c.heading, bold = true })

-- code
g(0, "Comment",    { fg = c.dim, italic = true })
g(0, "String",     { fg = c.string })
g(0, "Character",  { fg = c.string })
g(0, "Number",     { fg = c.string })
g(0, "Boolean",    { fg = c.string })
g(0, "Identifier", { fg = c.ink })
g(0, "Function",   { fg = c.link })
g(0, "Statement",  { fg = c.heading })
g(0, "Keyword",    { fg = c.heading })
g(0, "Operator",   { fg = c.dim })
g(0, "PreProc",    { fg = c.code })
g(0, "Type",       { fg = c.code })
g(0, "Constant",   { fg = c.string })
g(0, "Special",    { fg = c.link })
g(0, "Todo",       { fg = c.bg, bg = c.string, bold = true })
g(0, "Error",      { fg = c.err })

-- markdown, which is the reason this file exists.
--
-- Headings step by weight and not by brightness. A six-level brightness ramp
-- inside a 2.5-point band is invisible, and a ramp wide enough to see puts H1
-- back through the ceiling -- which is the problem being fixed.
g(0, "@markup.heading.1.markdown", { fg = c.heading, bold = true })
g(0, "@markup.heading.2.markdown", { fg = c.heading, bold = true })
g(0, "@markup.heading.3.markdown", { fg = c.heading })
g(0, "@markup.heading.4.markdown", { fg = c.heading })
g(0, "@markup.heading.5.markdown", { fg = c.link })
g(0, "@markup.heading.6.markdown", { fg = c.link })
g(0, "@markup.strong",   { fg = c.ink, bold = true })
g(0, "@markup.italic",   { fg = c.ink, italic = true })
g(0, "@markup.strikethrough", { fg = c.gutter, strikethrough = true })
g(0, "@markup.raw",           { fg = c.code })   -- `inline code`
g(0, "@markup.raw.block",     { fg = c.code })   -- fenced blocks
g(0, "@markup.link",          { fg = c.link })
g(0, "@markup.link.label",    { fg = c.link, underline = true })
g(0, "@markup.link.url",      { fg = c.dim, underline = true })
g(0, "@markup.list",          { fg = c.heading })
g(0, "@markup.list.checked",  { fg = c.code })
g(0, "@markup.list.unchecked",{ fg = c.dim })
g(0, "@markup.quote",         { fg = c.dim, italic = true })
g(0, "@markup.math",          { fg = c.string })
g(0, "@punctuation.special.markdown", { fg = c.gutter })
g(0, "@label.markdown",       { fg = c.heading })

-- the pre-treesitter group names, for anything not running a parser
g(0, "markdownH1",         { fg = c.heading, bold = true })
g(0, "markdownH2",         { fg = c.heading, bold = true })
g(0, "markdownH3",         { fg = c.heading })
g(0, "markdownH4",         { fg = c.heading })
g(0, "markdownH5",         { fg = c.link })
g(0, "markdownH6",         { fg = c.link })
g(0, "markdownCode",       { fg = c.code })
g(0, "markdownCodeBlock",  { fg = c.code })
g(0, "markdownLinkText",   { fg = c.link, underline = true })
g(0, "markdownUrl",        { fg = c.dim })
g(0, "markdownListMarker", { fg = c.heading })
g(0, "markdownRule",       { fg = c.border })
g(0, "markdownBlockquote", { fg = c.dim, italic = true })
g(0, "markdownHeadingDelimiter", { fg = c.gutter })

-- diagnostics
g(0, "DiagnosticError", { fg = c.err })
g(0, "DiagnosticWarn",  { fg = c.warn })
g(0, "DiagnosticInfo",  { fg = c.info })
g(0, "DiagnosticHint",  { fg = c.hint })

-- diffs
g(0, "DiffAdd",    { fg = c.code })
g(0, "DiffDelete", { fg = c.err })
g(0, "DiffChange", { fg = c.string })
g(0, "DiffText",   { fg = c.heading, bold = true })


-- a shell inside vim (:terminal) draws in these
-- sixteen, the same the terminal theme from this colourway uses,
-- so the two look alike
local term = {
  "#ef9286", "#000000", "#2d0055", "#000c2c", "#110b00", "#000000", "#440200", "#000305",
  "#5c5c51", "#000000", "#2e005e", "#000c2a", "#0f0a00", "#000000", "#4d0001", "#000000",
}
for i, v in ipairs(term) do vim.g["terminal_color_" .. (i - 1)] = v end
