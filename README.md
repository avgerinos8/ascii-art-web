
<h1 align="center">Ascii Art Web</h1>

<p align="center">
    <img src="_docs/icon.png" alt="logo" height="256px" width="256px" />
</p>

<p align="center">
    <a href="https://go.dev/dl/"><img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go" alt="Go Version" /></a>
    <img src="https://img.shields.io/badge/build-passing-brightgreen?style=for-the-badge" alt="Build Status" />
    <img src="https://img.shields.io/badge/Campus-zone01.gr-blue?style=for-the-badge&logo=gitea" alt="License" />
    <img src="https://img.shields.io/badge/Cohort-2.3-ed1c24?style=for-the-badge&labelColor=grey&logo=medusa" alt="Cohort" />
</p>

---

## 📖 Overview

**Ascii Art Web** is a fully interactive, responsive web application that upgrades the classic ASCII art generation tool into a feature-rich graphical dashboard. Built with a robust **Go server** backend and an advanced native **JavaScript/CSS frontend**, it enables users to generate, style, color, and customize ASCII art dynamically through their browser.

---

## 🚀 Getting Started

To spin up the web server locally, navigate to the project directory and execute:

```bash
go run .
```

Once the server initializes, open your browser and navigate to:
`http://localhost:8080` *(or your configured server port)*

---

## 🎨 Key Web Features

### 🔤 Font Matrix Selection
Choose from 6 uniquely styled banner fonts instantly via the stylized control bar:
* `Standard` (Default block text)
* `Shadow` (Dropped layout styling)
* `Thinkertoy` (Playful rounded look)
* `Bloody` (Horror/Graffiti type)
* `Extra` (Extended custom character sheet)
* `Stylish` (Condensed minimalist aesthetic)

### 📊 Interactive Visual Effects (`OPTIONS`)
Fine-tune the ambient UI background mechanics using responsive hardware-accelerated sliders:
* **Size & Distance:** Adjust background matrix element boundaries.
* **Speed & Depth:** Scale the biological procedural breath animations.
* **Grain:** Controls independent visual digital noise levels.

### 📝 Smart Text Input Terminal
* **Dynamic Row Sizing:** Scale your text terminal area between `1` and `6` rows natively using the custom `▲` / `▼` toggle buttons.
* **Realtime Processing:** Enable immediate live server requests as you type, or fall back to standard form submissions.
* **Formatting Suite:** Full support for custom `Word Wrap` thresholds and geometric positioning (`Left`, `Center`, `Right`, `Justify`).

### 🌈 Advanced HSL Substring Coloring
Build an unlimited stack of complex target-matching rules:
* Feed specific substrings into the dynamic color cards.
* Manipulate fine-grain **Hue**, **Saturation**, and **Lightness** (HSL) values using custom-compiled inline track thumbs.
* Match isolated phrases or fallback to an **"All Text"** global override canvas rule.

### 💾 LocalState Persistence
The interface contains memory triggers mapped to your browser's `localStorage`. Your active fonts, wrap configs, background metrics, and multi-row textarea structural sizes survive standard soft page reloads seamlessly. 
* *To clear your layout profile and return to factory defaults, simply click the **Reset** button.*

---

## 📁 Project Structure

```text
ascii-art-web/
├── cmd/
│   └── main.go                 # HTTP Web Server router and controller logic
├── internal/
│   ├── font/                   # Banner file parsers and ASCII render matrix engines
│   └── normalize/              # Multi-line text sanitization utilities
├── banners/                    # Font source assets (.txt format)
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│   ├── bloody.txt
│   ├── extra.txt
│   └── stylish.txt
└── static/                     # High-fidelity frontend asset directory
    ├── style.css               # Dynamic dark/light layout stylesheet variables
    ├── bg.js                   # Background particle animation and breath ticks
    ├── fetch.js                # Async server requests handling
    ├── darkMode.js             # Layout color profile toggles
    ├── fontWrap.js             # Multi-line word wrapping constraints
    ├── persistence.js          # LocalStorage state save/restore bindings
    ├── colorPreview.js         # Interactive HSL generator card logic
    ├── aScrollSync.js          # Synchronized layout positioning
    ├── footerCollapse.js       # Master footer panel animations
    └── featuresCollapse.js     # Granular modular configuration collagers
```

---

## 🛠️ Build and Deploy

Compile a standalone production-ready optimized binary executable file:

```bash
# Build the application
go build -o ascii-web ./cmd/main.go

# Start the application server
./ascii-web
```

---

## 📝 Authors

**Charles Newman** | cnewman | Cohort 2.3 | Zone01 Athens

**Nikos Antoniou** | nantoniou | Cohort 2.3 | Zone01 Athens

**Pavlos Avgerinos** | pavgerin | Cohort 2.3 | Zone01 Athens

---

## 📄 License

Completed as an advanced project within the official Zone01 Athens curriculum — Cohort 2.3.