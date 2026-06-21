package font

import (
	"fmt"
	"log/slog"
	"strings"

	color "asciiartweb/internal/color"
)

// ── Global Variables ───────────────────────────────────────────────────────

// Reset is a constant that returns the terminal formatting to its default state.
var Reset = "\033[0m"

// ── Struct ──────────────────────────────────────────────────────────────────

type Token struct {
	Value      rune   // 'h' ή ' '
	ColorValue string // "" (no color) ή "\033[0;31m"
}

type Line struct {
	Content []Token
	Length  int
	Spaces  int
}

func (f *Font) ComplexPrintLines(text string) {
	// If no color is selected
	if len(f.Con.Color) < 1 {
		Reset = ""
	}

	// STEP A : TOKENIZE (words, spaces)
	Characters := tokenize(text)

	// STEP Β : Add ColorValue
	f.addColorValue(Characters)

	Lines := []Line{}
	if f.Con.Output != "" {
		// STEP C : Predict output size to fit terminal width. Where needed split text to new lines.
		Lines = f.fitToTerminal(Characters)

		// STEP D : Distribute spaces to align correctly (right or center or justify) using measurements from STEP C.
		if f.Con.Align != "" {
			for i, line := range Lines {
				Lines[i].Content = f.alignLine(line)
			}
		}
	} else {
		Lines = []Line{{
			Content: Characters,
			Length:  len(Characters),
			Spaces:  0,
		}}
	}

	// STEP E : take tokenized slice and print ascii, line by line, and when you encounter colorcode tokentype add the ANSI ESCAPE CODE
	for _, line := range Lines {
		tempResult := ""
		tempResult = f.toAscii(line.Content)
		f.FinalResult = append(f.FinalResult, tempResult)
	}

}

/*
STEP A : TOKENIZATION
═════════════════════════════════════════════════════════════════════════
Input: text string +
Output: []Token
*/
// tokenize converts a raw input string into a slice of Token structures,
// preserving every character as a rune while initializing empty color states.
func tokenize(s string) []Token {
	result := make([]Token, 0, len(s))
	for _, letter := range s {
		result = append(result, Token{Value: letter, ColorValue: ""})
	}
	return result
}

/*
STEP Β : ColorValue
═════════════════════════════════════════════════════════════════════════
Read f.Con.Color ([]ColorInfo) and loop for every element
LOOP SUBSTRINGS with sorted priority (longer strings are already first in the slice f.Con.Color)
AND ADD ColorValue FOR EVERY TOKEN
*/

func (f *Font) addColorValue(Characters []Token) {
	for _, colorFlag := range f.Con.Color {
		if colorFlag.Substring == "" {
			slog.Error(fmt.Sprintf("Color %v has an empty substring! Skipping...", colorFlag.Num))
			continue
		}
		convertedColorToANSIescape, err := color.ColorConvert(colorFlag.ColorCode)
		if err != nil {
			slog.Error(fmt.Sprintf("Could not understand Color %v ColorCode! What kind of color is %s ? Skipping substring %s", colorFlag.Num, colorFlag.ColorCode, colorFlag.Substring))
			fmt.Println(err)
			continue
		}
		f.singleColorValue(Characters, colorFlag.Substring, convertedColorToANSIescape) // Character is being passed by reference in order to keep the changes without returning the slice and making copies
	}
}

/*
STEP B > Helper Function
Adds ColorValue to the []Token for the given substring (only for this one substring!)
This is called by AddColorValue for every substring in f.Con.Color and overwrites previous ColorValue if there is an overlap between substrings
Priority is sorted (longer strings are first)
*/
func (f *Font) singleColorValue(Characters []Token, Substring string, ANSIescape string) {
	subS := []rune(Substring)

	phraseLength := len(Characters)
	subsLength := len(subS)

	// Empty substring Early return
	if subsLength == 0 {
		slog.Error("Empty substring!", "substring", Substring)
		return
	}

	for i := 0; i < phraseLength; i++ {
		// Stop looking if the remaining slice elements are fewer than the substring length
		if i > phraseLength-subsLength {
			break
		}
		// Check if the current token value matches the first rune of the substring
		if Characters[i].Value == subS[0] {
			matched := 1
			// Scan subsequent tokens to confirm a continuous string match
			for j := 1; j < subsLength && j+i < phraseLength; j++ {
				if Characters[j+i].Value == subS[j] {
					matched++
				} else {
					break
				}
			}
			// If the complete substring is successfully matched, apply the ANSI color
			if matched == subsLength {
				Characters[i].ColorValue = ANSIescape
				for j := 1; j < subsLength && j+i < phraseLength; j++ {
					Characters[i+j].ColorValue = ANSIescape
				}
				// Skip the processed substring tokens in the outer loop execution
				// i += subsLength - 1
			}
		}
	}
}

/*
STEP C : fitToTerminal splits the provided input into an array of Lines along with measurements
regarding length of result and spaces needed in the align case.
*/
func (f *Font) fitToTerminal(Characters []Token) (Lines []Line) {
	// Break tokens into lines that fit within terminal width
	width := f.Con.PageCharacterWidth
	var currentWidth, start, lastspace, space int

	for i, r := range Characters {
		// Track last space for word boundary breaks
		if r.Value == ' ' {
			lastspace = i
			space++
		}

		// Accumulate width of current token
		currentWidth += f.BannerWidth[r.Value]

		// Create line when width exceeded or at end of input
		if currentWidth >= width || i == len(Characters)-1 {
			remainder := 0
			var newLine Line
			if i == len(Characters)-1 && !(currentWidth > width) {
				newLine.Content = Characters[start:]
			} else if r.Value == ' ' {
				newLine.Content = Characters[start:i]
				start = i + 1
				remainder = f.BannerWidth[' ']
				space--
			} else if lastspace == 0 {
				newLine.Content = Characters[start:i]
				start = i
			} else {
				newLine.Content = Characters[start:lastspace]
				start = lastspace + 1
				remainder = f.BannerWidth[' ']
				space--
			}
			lastspace = 0
			// Calculate line metrics (width and space count)
			for _, ch := range newLine.Content {
				newLine.Length += f.BannerWidth[ch.Value]
			}
			newLine.Spaces = space
			space = 0
			Lines = append(Lines, newLine)

			// Handle overflow at end: create additional line for remaining text
			if i == len(Characters)-1 && (currentWidth > width) {
				newLine.Content = Characters[start:]
				newLine.Length = 0
				for _, ch := range newLine.Content {
					if ch.Value == ' ' {
						space++
					}
					newLine.Length += f.BannerWidth[ch.Value]
				}
				newLine.Spaces = space
				Lines = append(Lines, newLine)
			}
			currentWidth -= (newLine.Length + remainder)
		}
	}
	return Lines
}

func (f *Font) alignLine(line Line) []Token {
	// Apply alignment by inserting padding tokens (rune(0) = space)
	Align := f.Con.Align
	ConsoleWidth := f.Con.PageCharacterWidth
	ExtraSpace := ConsoleWidth - line.Length
	if ExtraSpace < 0 {
		ExtraSpace = 0
	}
	var Aligned []Token
	// Fall back to left align if justify requested but no spaces exist
	if line.Spaces < 1 && Align == "justify" {
		Align = "left"
	}
	switch Align {
	case "center":
		// Add padding to left half
		for i := 0; i < ExtraSpace/2; i++ {
			Aligned = append(Aligned, Token{Value: rune(0)})
		}
		Aligned = append(Aligned, line.Content...)
	case "right":
		// Add all padding to left
		for i := 0; i < ExtraSpace; i++ {
			Aligned = append(Aligned, Token{Value: rune(0)})
		}
		Aligned = append(Aligned, line.Content...)
	case "justify":
		// Distribute extra space evenly among word spaces
		SpacesWidth := make([]int, line.Spaces)
		for i := 0; i < ExtraSpace; i++ {
			SpacesWidth[i%line.Spaces]++
		}
		SpaceIndex := 0
		for _, t := range line.Content {
			if t.Value != ' ' {
				Aligned = append(Aligned, t)
			} else {
				// Add space and its extra padding
				Aligned = append(Aligned, t)
				for i := 0; i < SpacesWidth[SpaceIndex]; i++ {
					Aligned = append(Aligned, Token{Value: rune(0)})
				}
				if SpaceIndex < len(SpacesWidth)-1 {
					SpaceIndex++
				}
			}
		}
	default:
		// Left align: no padding added
		return line.Content
	}
	return Aligned
}

func (f *Font) toAscii(input []Token) string {
	var result strings.Builder

	for i := 0; i < 8; i++ {
		// Since we append Reset at the end of every row,
		// the terminal always starts the next row in a Reset state.
		lastColor := Reset

		for _, t := range input {
			// Resolve the actual ANSI string we want for this token
			currentColor := t.ColorValue
			if currentColor == "" {
				currentColor = Reset
			}

			// If the required color is different from the terminal's current state, emit it
			if currentColor != lastColor {
				result.WriteString(currentColor)
				lastColor = currentColor
			}

			result.WriteString(f.Banner[t.Value][i])
		}
		// Always clear color at the end of the line to prevent trailing bleed
		result.WriteString(Reset + "\n")
	}
	return result.String()
}
