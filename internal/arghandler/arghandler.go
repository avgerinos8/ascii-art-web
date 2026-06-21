package arghandler

import (
	normalize "asciiartweb/internal/normalize"
)

// ── Structs ────────────────────────────────────────────────────────────────⊃

// Config stores all the settings derived from command-line arguments.
// It acts as the central state for the ASCII art generation process.
type Config struct {
	Font               string
	Color              []ColorInfo
	Text               []string
	Output             string
	PageCharacterWidth int
	Align              string
}

// ColorInfo holds specific data for text coloring, including which
// part of the text to color and the corresponding ANSI/color code.
type ColorInfo struct {
	Num       int
	ColorCode string
	Substring string
}

// ── Functions ──────────────────────────────────────────────────────────────⊃

// NewConfig returns a pointer to a fresh Config instance.
func NewConfig() *Config {
	c := &Config{}
	return c
}

func (c *Config) NormalizeInput(input string) {
	input, _ = normalize.NonPrintables(input)
	c.Text = normalize.TabsNewLines(input)
}
