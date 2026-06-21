package normalize

import (
	"fmt"
	"log/slog"
	"strings"
)

// ── Functions ──────────────────────────────────────────────────────────────

// NonPrintables cleans the input string by replacing tabs with spaces,
// handling newlines, and filtering out any non-printable ASCII characters.
// It logs the changes made and returns a slice of strings split by newlines.
func NonPrintables(s string) (string, string) {

	var nonSupported strings.Builder
	var cleaned strings.Builder

	// Iterate through each rune to filter out non-printable characters (outside ASCII space to tilde)
	for i, r := range s {
		if r == '\n' {
			cleaned.WriteRune(r)
			continue
		}
		if !(r >= ' ' && r <= '~') {
			nonSupported.WriteRune(r)
			nonSupported.WriteRune(' ')
			if i+1 >= len(s) {
				continue
			}
			continue

		}
		cleaned.WriteRune(r)
	}

	return cleaned.String(), nonSupported.String()

}

func TabsNewLines(s string) []string {
	//beforetabsize := len(s)

	// Replace escaped tab characters with 3 spaces for consistent alignment
	s = strings.ReplaceAll(s, "\\t", "   ")
	//aftertabsize := len(s)
	//slog.Info(fmt.Sprintf("There were %v tab spaces (\t) found.", aftertabsize-beforetabsize))

	// Normalize newline characters and prepare builders for efficient string manipulation
	s = strings.ReplaceAll(s, "\\n", "\n")

	var res []string

	// Split the cleaned text back into a slice based on the newline markers
	res = strings.Split(s, "\n")
	slog.Info(fmt.Sprintf("User given input contained %v new lines (\n).", len(res)-1))

	return res
}

func PrintNonSupported(allnonsupported []string) {
	accumulated := ""

	for _, letters := range allnonsupported {
		accumulated += letters
	}
	// Handle the reporting of any omitted characters found during the loop
	if accumulated != "" {
		// Determine the correct grammar (singular/plural) for the log message
		plural1 := "these"
		plural2 := "s"
		plural3 := "were"
		if len([]rune(accumulated)) == 2 {
			plural1 = "this"
			plural2 = ""
			plural3 = "was"
		}
		// Construct the final strings for output
		partOne := "Input contained " + plural1 + " non-supported character" + plural2
		partTwo := "which " + plural3 + " omitted."

		// Log the information via slog
		slog.Info(fmt.Sprintf(partOne+" %s"+partTwo, accumulated))

		// Define ANSI escape codes for stylized console output
		infoColor := "\033[1;97;48;2;89;0;126m "
		infoReset := "\033[0m"

		// Print the formatted message directly to the console
		fmt.Printf(partOne+" %s%s%s "+partTwo, infoColor, accumulated, infoReset)
		fmt.Println("")
	}

}
