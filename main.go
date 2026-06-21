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

// Fallback logo definition used during initialization
const logo = `        A    
        V   
ASCII   G    
N       E    
T      ART    N
O       I     E
N       N     WEB
I       O     M
O       S     A
U             N`

// Global runtime template context reference
var Data = PageData{
	Title:     "ASCII Art Web",
	Header:    "",
	InitInput: "Please type something here...",
	// UserText is used to save user last input text and persist across page reloads
	UserText: "",
	Output:   logo,
}

var fonts []string = []string{"standard", "shadow", "thinkertoy", "extra", "blody", "stylish"}

var Config *arghandler.Config = arghandler.NewConfig()

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

	// Intercept traditional Form POST submissions
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err == nil {
			// Print directly to console
			// Print directly to console without mutating or breaking the blank identifiers (_)
			fmt.Printf("[Form POST Log] UserText: %s | FontWrap: %s | Realtime: %s | FontAlign: %s | font: %s | MaxChars: %s\n",
				r.FormValue("userText"),
				r.FormValue("FontWrap"),
				r.FormValue("Realtime"),
				r.FormValue("FontAlign"),
				r.FormValue("font"),
				r.FormValue("max_chars"),
			)

			// Read form entries straight into the blank identifier (_) using exact HTML layout names
			Data.UserText = r.FormValue("userText")
			Config.NormalizeInput(r.FormValue("userText"))
			Config.Output = r.FormValue("FontWrap")
			_ = r.FormValue("Realtime")
			Config.Align = r.FormValue("FontAlign")
			Config.Font = r.FormValue("font")
			Config.PageCharacterWidth, err = strconv.Atoi(r.FormValue("max_chars"))
			if err != nil {
				Config.PageCharacterWidth = 80
				fmt.Println("Invalid max_chars value. Defaulting to 80 characters per line.")
				w.WriteHeader(http.StatusBadRequest)
			}

			f := font.CreateFont(Config)
			f.RenderResult()
			var rawHTML string
			for i := 0; i < len(f.FinalResult); i++ {
				rawHTML += "<span class=\"line\">" + f.FinalResult[i] + "</span>\n"
			}

			Data.Output = template.HTML(rawHTML)
		}
	}

	// Parse the HTML file from disk
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Execute the template and send Output back to the client using POST
	templ.Execute(w, Data)
}

// SessionHandler unpacks asynchronous dynamic JSON requests submitted by Javascript Fetch
func SessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// DELETE THIS MAYBE AFTER YOUR IMPLEMENTATION (anonymous struct <-LOVE!- to reference JSON properties):
	var incoming struct {
		FontWrap   string `json:"font_wrap"`
		FontAlign  string `json:"font_align"`
		ActiveFont string `json:"active_font"`
		Realtime   bool   `json:"realtime"`
		UserText   string `json:"user_text"`
		MaxChars   int    `json:"max_chars"`
	}

	// Go parses the JSON and populates the struct automatically
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

	// Assign clean properties
	Config.Output = incoming.FontWrap
	Config.Align = incoming.FontAlign
	Config.Font = incoming.ActiveFont
	_ = incoming.Realtime
	Config.NormalizeInput(incoming.UserText)
	Config.PageCharacterWidth = incoming.MaxChars

	f := font.CreateFont(Config)
	f.RenderResult()
	var rawHTML string
	for i := 0; i < len(f.FinalResult); i++ {
		rawHTML += "<span class=\"line\">" + f.FinalResult[i] + "</span>\n"
	}

	Data.Output = template.HTML(rawHTML)
	// Update the global template data with the generated result.
	// TODO: replace this string literal with core ASCII generator function.

	// Inform the client browser that the response body contains data formatted as JSON.
	w.Header().Set("Content-Type", "application/json")

	// Convert the Go struct 'Data' into a JSON string and stream it back to the Javascript Fetch API. This is the way we send Output back using the Fetch API
	json.NewEncoder(w).Encode(Data)
}
