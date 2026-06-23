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
	mux.HandleFunc("/", Handler)

	fmt.Println("Server ready. Use http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

// Standard UI Template Handler for index loading and traditional form POSTs
func Handler(w http.ResponseWriter, r *http.Request) {
	// Secure routing definitions for both endpoints mapped to this handler
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		w.WriteHeader(http.StatusNotFound)
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
			fmt.Printf("[Form POST Log] UserText: %s | FontWrap: %s | Realtime: %s | FontAlign: %s | font: %s | MaxChars: %s\n",
				r.FormValue("userText"),
				r.FormValue("FontWrap"),
				r.FormValue("Realtime"),
				r.FormValue("FontAlign"),
				r.FormValue("font"),
				r.FormValue("max_chars"),
			)

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

	var incoming struct {
		FontWrap   string `json:"font_wrap"`
		FontAlign  string `json:"font_align"`
		ActiveFont string `json:"active_font"`
		Realtime   bool   `json:"realtime"`
		UserText   string `json:"user_text"`
		MaxChars   int    `json:"max_chars"`
	}

	if err := json.NewDecoder(r.Body).Decode(&incoming); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	fmt.Printf("[Fetch JSON Log] UserText: %s | FontWrap: %s | Realtime: %t | FontAlign: %s | ActiveFont: %s | MaxChars: %d\n",
		incoming.UserText,
		incoming.FontWrap,
		incoming.Realtime,
		incoming.FontAlign,
		incoming.ActiveFont,
		incoming.MaxChars,
	)

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
	localConfig.NormalizeInput(incoming.UserText)
	localConfig.PageCharacterWidth = incoming.MaxChars
	localData.UserText = incoming.UserText

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
