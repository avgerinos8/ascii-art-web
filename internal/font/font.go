package font

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	arghandler "asciiartweb/internal/arghandler"
)

// ── Global Variables ───────────────────────────────────────────────────────

// ── Structs ────────────────────────────────────────────────────────────────

// Font represents the engine for ASCII art generation.
// It holds the configuration, the mapped banner characters,
// and the final slice of strings to be printed.

type Font struct {
	Con         *arghandler.Config
	Banner      map[rune]([]string)
	BannerWidth map[rune](int)
	FinalResult []string
}

// ── Functions ──────────────────────────────────────────────────────────────

// CreateFont initializes a new Font instance with the provided config.
// It automatically triggers the banner reading process before returning.
func CreateFont(config *arghandler.Config) *Font {
	f := &Font{Con: config}
	f.ReadBanner()
	f.Banner[rune(0)] = []string{" ", " ", " ", " ", " ", " ", " ", " "}
	f.CountBannerWidth()
	// f.PrintBannerWidthDebug() // ONLY FOR DEBUGGING

	return f
}

// ── Methods ────────────────────────────────────────────────────────────────

// ReadBanner opens the specified banner file and parses it into a map.
// It maps each ASCII character starting from ' ' to its 8-line representation.
func (f *Font) ReadBanner() {
	filename := "banners/" + f.Con.Font + ".txt"

	// Check if the file exists in the current working directory
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// If not found, try the path for execution from within the /cmd folder
		filename = "../" + filename

		if _, err := os.Stat(filename); os.IsNotExist(err) {
			// If still not found, try the path for execution from outside the repo folder
			// Note: This assumes the repository folder is named 'ascii-art'
			filename = "ascii-art/banners/" + f.Con.Font + ".txt"
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		slog.Error("The file required to read ascii art banners is missing", "filename", filename)
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var banner map[rune][]string
	banner = make(map[rune][]string)
	characterCounter := 0
	counter := 0

	// Standard banners have 8 lines per character, usually separated by an empty line
	for scanner.Scan() {
		if counter > 0 && counter < 9 {
			character := rune(' ' + characterCounter)
			banner[character] = append(banner[character], scanner.Text())
		}
		if counter == 8 {
			counter = 0
			characterCounter++
			continue
		}
		counter++
	}
	f.Banner = banner
}

func (f *Font) CountBannerWidth() {
	f.BannerWidth = make(map[rune]int)
	for r, sl := range f.Banner {
		f.BannerWidth[r] = len(sl[0])
	}
}

// RenderResult processes the input text, handles word-wrapping based on
// terminal width, and populates the finalresult slice with ASCII art lines.
func (f *Font) RenderResult() {
	fileOutput := false
	// Iterate over the input text
	if fileOutput {
		for _, textToBePrinted := range f.Con.Text {
			f.SimplePrintLines(textToBePrinted) // simple method for FILE output WITHOUT colors, alignment
		}
	} else {
		for _, textToBePrinted := range f.Con.Text {
			f.ComplexPrintLines(textToBePrinted) // complex method for CONSOLE output WITH colors, alignment, etc
		}
	}
}

func (f *Font) SimplePrintLines(text string) {
	result := ""
	for i := 0; i < 8; i++ {
		for _, letter := range text {
			result += f.Banner[letter][i]
		}
		result += "\n"
	}
	if result != "" {
		f.FinalResult = append(f.FinalResult, result)
	}
}

// ────────────────────────────────────────────────────────────────────────────
// PrintResult writes all lines stored in finalresult to the given io.Writer.
// This allows flexible output to stdout, files, or buffers.
func (f *Font) PrintResult() {
	mode := os.Stdout
	slog.Info("Output mode is set to console. Writing...")
	// Changes output based on output flag
	if f.Con.Output != "" {
		slog.Info("Output mode is set to file. Writing...", "filename", f.Con.Output)
		destination, er := os.Create(f.Con.Output)
		if er != nil {
			fmt.Println(er)
			return
		}
		mode = destination
		fmt.Printf("Operation complete! Result can be found in %s\n", f.Con.Output)
	}
	for _, finalLine := range f.FinalResult {
		fmt.Fprint(mode, finalLine)
	}
	slog.Info("ASCII Art printed.")
}

func (f *Font) PrintBannerWidthDebug() {
	for i := ' '; i <= '~'; i++ {
		fmt.Printf(" %s : %d\t", string(i), f.BannerWidth[i])
	}
	fmt.Printf(" %s : %d\n", string(rune(0)), f.BannerWidth[rune(0)])
}
