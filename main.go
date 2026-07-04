// ── main ────────────────────────────────────────────────────────────────────⊃

package main

import (
	"asciiartweb/internal/arghandler"
	"asciiartweb/internal/font"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var PageWidth = 80

// PageData maps out template variable configurations injected straight into index.html
type PageData struct {
	Title     string
	Header    string
	InitInput string
	Output    template.HTML
	UserText  string
}

var fonts []string = []string{"standard", "shadow", "thinkertoy", "extra", "blody", "stylish"}

// Entrypoint configuration routing all endpoints
func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/api/session-state", SessionHandler)

	mux.HandleFunc("/ascii-art", Handler)
	mux.HandleFunc("/", Handler)

	fmt.Println("Server ready. Use http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

// Standard UI Template Handler for index loading and traditional form POSTs
func Handler(w http.ResponseWriter, r *http.Request) {
	// Secure routing definitions for both endpoints mapped to this handler
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		http.NotFound(w, r)
		return
	}

	// Create a FRESH, isolated instance of PageData and Config for THIS request only
	localData := PageData{
		Title:     "ASCII Art Web",
		Header:    "",
		InitInput: "Please type something here...",
		UserText:  "",
		Output:    "",
	}
	localConfig := arghandler.NewConfig()

	// Intercept traditional Form POST submissions
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err == nil {

			// Populate our local data structures dynamically
			localData.UserText = r.FormValue("userText")
			localConfig.NormalizeInput(r.FormValue("userText"))
			localConfig.Output = r.FormValue("FontWrap")
			_ = r.FormValue("Realtime")
			localConfig.Align = r.FormValue("FontAlign")
			localConfig.Font = r.FormValue("font")

			var err error
			localConfig.PageCharacterWidth, err = strconv.Atoi(r.FormValue("max_chars"))
			if err != nil {
				localConfig.PageCharacterWidth = 80
				fmt.Println("Invalid max_chars value. Defaulting to 80 characters per line.")
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Extract the parallel form arrays submitted from the DOM color blocks
			formSubstrings := r.PostForm["substring[]"]
			formHexCodes := r.PostForm["hexcolorcode[]"]

			var colors []arghandler.ColorInfo
			colorCounter := 1

			// Loop through arrays using the substrings count length as our limit boundary
			for i := 0; i < len(formSubstrings); i++ {
				subStr := formSubstrings[i]
				hexStr := ""
				if i < len(formHexCodes) {
					hexStr = formHexCodes[i]
				}

				// If JS flagged this form element via readOnly string, expand it to all text lines
				if subStr == "_ALL_TEXT_" {
					for _, textLine := range localConfig.Text {
						// Skip empty lines to prevent unnecessary color processing overhead
						if textLine == "" {
							continue
						}
						colors = append(colors, arghandler.ColorInfo{
							Num:       colorCounter,
							ColorCode: hexStr,
							Substring: textLine, // Target the full contents of this row line directly
						})
						colorCounter++
					}
				} else {
					// Standard explicit single substring mapping process branch
					colors = append(colors, arghandler.ColorInfo{
						Num:       colorCounter,
						ColorCode: hexStr,
						Substring: subStr,
					})
					colorCounter++
				}
			}

			// Assign parsed dynamic arrays right into your Config object context state
			localConfig.Color = colors
			localConfig.SortColors()

			// Enhanced Debug Logging for Form POST
			fmt.Printf("[Form POST Log] UserText: %s | FontWrap: %s | Realtime: %s | FontAlign: %s | font: %s | MaxChars: %s | Colors Count: %d\n",
				r.FormValue("userText"),
				r.FormValue("FontWrap"),
				r.FormValue("Realtime"),
				r.FormValue("FontAlign"),
				r.FormValue("font"),
				r.FormValue("max_chars"),
				len(localConfig.Color),
			)
			if len(localConfig.Color) > 0 {
				fmt.Println("--- Sorted Colors Detail (POST) ---")
				for _, c := range localConfig.Color {
					fmt.Printf("  -> Num: %d | Substring: %q | ColorCode/HEX: %s\n", c.Num, c.Substring, c.ColorCode)
				}
				fmt.Println("-----------------------------------")
			}

			f := font.CreateFont(localConfig)
			f.RenderResult()
			var rawHTML string
			for i := 0; i < len(f.FinalResult); i++ {
				rawHTML += "<span class=\"line\">" + f.FinalResult[i] + "</span>"
			}

			localData.Output = template.HTML(rawHTML)
		}
	}

	// Parse the HTML file from disk
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Execute using the isolated localData context safely
	templ.Execute(w, localData)
}

// SessionHandler unpacks asynchronous dynamic JSON requests submitted by Javascript Fetch
func SessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Mirror anonymous schema structure containing explicit slice mappings for colors array JSON data
	var incoming struct {
		FontWrap   string                 `json:"font_wrap"`
		FontAlign  string                 `json:"font_align"`
		ActiveFont string                 `json:"active_font"`
		Realtime   bool                   `json:"realtime"`
		UserText   string                 `json:"user_text"`
		MaxChars   int                    `json:"max_chars"`
		Colors     []arghandler.ColorInfo `json:"colors"` // Bound directly to external arghandler package struct
	}

	if err := json.NewDecoder(r.Body).Decode(&incoming); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Create isolated metrics for this specific async transaction
	localData := PageData{
		Title:     "ASCII Art Web",
		InitInput: "Please type something here...",
	}
	localConfig := arghandler.NewConfig()

	localConfig.Output = incoming.FontWrap
	localConfig.Align = incoming.FontAlign
	localConfig.Font = incoming.ActiveFont
	_ = incoming.Realtime

	// Normalize text input first so localConfig.Text ([]string) gets populated with rows
	localConfig.NormalizeInput(incoming.UserText)
	localConfig.PageCharacterWidth = incoming.MaxChars
	localData.UserText = incoming.UserText

	var finalColors []arghandler.ColorInfo
	colorCounter := 1

	// Loop through the incoming JSON colors array sent by fetch.js
	for _, incomingColor := range incoming.Colors {
		// FIXED: Check for explicit _ALL_TEXT_ flag instead of empty string
		if incomingColor.Substring == "_ALL_TEXT_" {
			for _, textLine := range localConfig.Text {
				if textLine == "" {
					continue
				}
				finalColors = append(finalColors, arghandler.ColorInfo{
					Num:       colorCounter,
					ColorCode: incomingColor.ColorCode,
					Substring: textLine,
				})
				colorCounter++
			}
		} else {
			// Standard single substring matching pass execution
			finalColors = append(finalColors, arghandler.ColorInfo{
				Num:       colorCounter,
				ColorCode: incomingColor.ColorCode,
				Substring: incomingColor.Substring,
			})
			colorCounter++
		}
	}

	// Pass the expanded color slices configuration down into internal operational config states
	localConfig.Color = finalColors

	localConfig.SortColors()

	// Enhanced Debug Logging for Fetch JSON
	fmt.Printf("[Fetch JSON Log] UserText: %s | FontWrap: %s | Realtime: %t | FontAlign: %s | ActiveFont: %s | MaxChars: %d | Expanded Colors Count: %d\n",
		incoming.UserText,
		incoming.FontWrap,
		incoming.Realtime,
		incoming.FontAlign,
		incoming.ActiveFont,
		incoming.MaxChars,
		len(localConfig.Color),
	)
	if len(localConfig.Color) > 0 {
		fmt.Println("--- Sorted Colors Detail (FETCH) ---")
		for _, c := range localConfig.Color {
			fmt.Printf("  -> Num: %d | Substring: %q | ColorCode/HEX: %s\n", c.Num, c.Substring, c.ColorCode)
		}
		fmt.Println("------------------------------------")
	}

	f := font.CreateFont(localConfig)
	f.RenderResult()
	var rawHTML string
	for i := 0; i < len(f.FinalResult); i++ {
		rawHTML += "<span class=\"line\">" + f.FinalResult[i] + "</span>"
	}

	localData.Output = template.HTML(rawHTML)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(localData)
}
