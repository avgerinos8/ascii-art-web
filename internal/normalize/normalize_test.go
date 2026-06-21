package normalize

import (
	"reflect"
	"testing"
)

func TestNonPrintables(t *testing.T) {
	// Define test cases for NonPrintables function
	tests := []struct {
		name        string
		input       string
		wantCleaned string
		wantNonSupp string
	}{
		{
			name:        "Only printable ASCII characters",
			input:       "Hello World! 123",
			wantCleaned: "Hello World! 123",
			wantNonSupp: "",
		},
		{
			name:        "Contains non-printable characters",
			input:       "Hello\x00World\x01!",
			wantCleaned: "HelloWorld!",
			wantNonSupp: "\x00 \x01 ",
		},
		{
			name:        "Empty string input",
			input:       "",
			wantCleaned: "",
			wantNonSupp: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCleaned, gotNonSupp := NonPrintables(tt.input)

			// Verify the cleaned string output
			if gotCleaned != tt.wantCleaned {
				t.Errorf("NonPrintables() gotCleaned = %q, want %q", gotCleaned, tt.wantCleaned)
			}
			// Verify the non-supported characters output
			if gotNonSupp != tt.wantNonSupp {
				t.Errorf("NonPrintables() gotNonSupp = %q, want %q", gotNonSupp, tt.wantNonSupp)
			}
		})
	}
}

func TestTabsNewLines(t *testing.T) {
	// Define test cases for TabsNewLines function
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "No tabs or newlines present",
			input: "Hello World",
			want:  []string{"Hello World"},
		},
		{
			name:  "Contains escaped tabs",
			input: "Hello\\tWorld",
			want:  []string{"Hello   World"},
		},
		{
			name:  "Contains newline characters",
			input: "Line1\nLine2\nLine3",
			want:  []string{"Line1", "Line2", "Line3"},
		},
		{
			name:  "Combined escaped tabs and newlines",
			input: "Line1\\tData\nLine2",
			want:  []string{"Line1   Data", "Line2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TabsNewLines(tt.input)

			// Use DeepEqual to compare slices
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TabsNewLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintNonSupported(t *testing.T) {
	// This function writes directly to stdout and returns no value.
	// We verify that it executes safely without panicking under various inputs.
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "Empty slice input",
			input: []string{},
		},
		{
			name:  "Single non-supported character",
			input: []string{"\x00 "}, // Length of 2 runes triggers the singular log logic
		},
		{
			name:  "Multiple non-supported characters",
			input: []string{"\x00 ", "\x01 "}, // Triggers the plural log logic
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Confirm the function executes smoothly
			PrintNonSupported(tt.input)
		})
	}
}
