package color

import (
	"testing"
)

// ── test structures and data models ──────────────────────────────────────── ⊃

// ColorConvTests defines the schema for table-driven test cases,
// holding the raw color input string, its expected ANSI escape sequence,
// and a boolean flag indicating if an error is anticipated.
type ColorConvTests struct {
	Input       string
	Expected    string
	ExpectError bool
}

// ── test execution suites ────────────────────────────────────────────────── ⊃

// TestColorConvert executes table-driven tests against the ColorConvert function.
// It validates both valid and invalid scenarios for basic text, HEX, RGB, and HSL
// inputs, verifying that strings match and error states are handled correctly.
func TestColorConvert(t *testing.T) {
	var AllTests []ColorConvTests = WriteTests()
	for _, test := range AllTests {
		value, err := ColorConvert(test.Input)

		if test.ExpectError {
			if err == nil {
				t.Errorf("ColorConvert(%q) expected an error, but got none", test.Input)
			}
			continue
		}

		if err != nil {
			t.Errorf("ColorConvert(%q) returned unexpected error: %v", test.Input, err)
			continue
		}

		if value != test.Expected {
			t.Errorf("ColorConvert(%q) -> Expected %q, got %q", test.Input, test.Expected, value)
		}
	}
}

// WriteTests initializes and returns a slice of ColorConvTests containing
// various color format test cases (basic names, hex, rgb, hsl) for validation.
func WriteTests() []ColorConvTests {
	return []ColorConvTests{
		// Basic Text Colors
		{Input: "black", Expected: "\033[0;30m", ExpectError: false},
		{Input: "red", Expected: "\033[0;31m", ExpectError: false},
		{Input: "bold", Expected: "\033[1m", ExpectError: false},
		{Input: "orange", Expected: "\033[38;5;208m", ExpectError: false},
		{Input: "invalidcolor", Expected: "\033[0m", ExpectError: true},

		// HEX Colors (3-digit & 6-digit)
		{Input: "#FFF", Expected: "\033[38;2;255;255;255m", ExpectError: false},
		{Input: "#000000", Expected: "\033[38;2;0;0;0m", ExpectError: false},
		{Input: "#FF0000", Expected: "\033[38;2;255;0;0m", ExpectError: false},
		{Input: "#12345", Expected: "\033[0m", ExpectError: true},  // Invalid length
		{Input: "#GHIJKL", Expected: "\033[0m", ExpectError: true}, // Invalid hex characters

		// RGB Colors
		{Input: "rgb(0,0,0)", Expected: "\033[38;2;0;0;0m", ExpectError: false},
		{Input: "rgb(255,105,180)", Expected: "\033[38;2;255;105;180m", ExpectError: false},
		{Input: "rgb(256,0,0)", Expected: "\033[0m", ExpectError: true}, // Out of range (256)
		{Input: "rgb(100,100)", Expected: "\033[0m", ExpectError: true}, // Missing components

		// HSL Colors
		{Input: "hsl(0,100,50)", Expected: "\033[38;2;255;0;0m", ExpectError: false},   // Pure Red
		{Input: "hsl(120,100,50)", Expected: "\033[38;2;0;255;0m", ExpectError: false}, // Pure Green
		{Input: "hsl(240,100,50)", Expected: "\033[38;2;0;0;255m", ExpectError: false}, // Pure Blue
		{Input: "hsl(361,50,50)", Expected: "\033[0m", ExpectError: true},              // Hue out of range
		{Input: "hsl(180,105,50)", Expected: "\033[0m", ExpectError: true},             // Saturation out of range
	}
}
