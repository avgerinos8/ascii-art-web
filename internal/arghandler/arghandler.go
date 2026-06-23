package arghandler

import (
	normalize "asciiartweb/internal/normalize"
	"sort"
)

// Config stores all the settings derived from command-line arguments.
type Config struct {
	Font               string      `json:"font"`
	Color              []ColorInfo `json:"colors"`
	Text               []string    `json:"text"`
	Output             string      `json:"output"`
	PageCharacterWidth int         `json:"page_character_width"`
	Align              string      `json:"align"`
}

// ColorInfo holds specific data for text coloring.
type ColorInfo struct {
	Num       int    `json:"num"`
	ColorCode string `json:"color_code"`
	Substring string `json:"substring"`
}

// NewConfig returns a pointer to a fresh Config instance.
func NewConfig() *Config {
	c := &Config{}
	return c
}

func (c *Config) NormalizeInput(input string) {
	input, _ = normalize.NonPrintables(input)
	c.Text = normalize.TabsNewLines(input)
}

// SortColors orders the colors slice by substring length descending.
// Since "All Text" fields were already expanded into full text lines by the handlers,
// sorting by maximum length automatically places "All Text" rows first.
func (c *Config) SortColors() {
	sort.Slice(c.Color, func(i, j int) bool {
		// Sort by string length descending (longest string first)
		return len(c.Color[i].Substring) > len(c.Color[j].Substring)
	})

	// Re-assign the running sequence indexes from 1 to N after sorting is complete
	for idx := range c.Color {
		c.Color[idx].Num = idx + 1
	}
}
