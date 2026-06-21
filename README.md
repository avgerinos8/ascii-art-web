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

**Ascii-art** is a Go-based CLI tool that transforms text into stylized ASCII art. It supports multiple banner fonts, selective ANSI coloring (by name, HEX, RGB, or HSL), output to files, terminal alignment, and reverse mode (ASCII art back to text).

---

## 🚀 Usage

```bash
go run ./... [FLAGS] [TEXT] [BANNER]
```

All flags use `--flag=value` syntax. `TEXT` and `BANNER` are positional arguments. The order of flags relative to each other doesn't matter, but positional arguments (substrings and text) are assigned left-to-right.

---

## 📐 Align Flag (`--align`)

Aligns the ASCII art output relative to the terminal width.

```bash
go run ./... --align=center "Hello"
go run ./... --align=right "Hello"
go run ./... --align=left "Hello"
go run ./... --align=justify "Hello"
```

Valid values: `left`, `right`, `center`, `justify`.

Can be combined with colors and fonts:

```bash
go run ./... --align=center --color=cyan "Hello World" shadow
go run ./... --align=right --color1=red --substring1=Hello "Hello World" standard
```
---

## 🎨 Color Flag (`--color`)

Colors specific text using ANSI escape codes. By default, colors the entire output.

```bash
# Color the entire output red
go run ./... --color=red "Hello World"

# Color only a specific substring
go run ./... --color=blue --substring=World "Hello World"

# Shorthand: substring as a positional argument after the flag
go run ./... --color=blue World "Hello World"
```

### Multiple colors

Use numbered suffixes to apply different colors to different parts of the text.

```bash
# Two colors, using --substring flags
go run ./... --color1=red --substring1=Hello --color2=blue --substring2=World "Hello World"

# Two colors, using positional substrings
go run ./... --color1=red --color2=blue Hello World "Hello World"

# Three colors
go run ./... --color1=red --substring1=H --color2=green --substring2=ello --color3=blue --substring3=World "Hello World"
```

> **Note:** When multiple colors match the same character, the one with the **longest substring** wins (greedy match). This means `--substring=Hello` takes priority over `--substring=H` for the letter `H`.

### Supported color formats

| Format | Example | Notes |
|---|---|---|
| Name | `--color=red` | Plain color name |
| HEX (6-digit) | `--color=#ff5733` | Full hex code |
| HEX (3-digit) | `--color=#f53` | Shorthand, expands to 6 digits |
| RGB | `--color=rgb(255,87,51)` | 0–255 per channel |
| HSL | `--color=hsl(14,100,60)` | H: 0–360, S/L: 0–100 |

### All named colors

```
black      red        green      yellow     blue
purple     cyan       white      bold

brightblack    brightred      brightgreen   brightyellow
brightblue     brightpurple   brightcyan    brightwhite

orange   pink     magenta   gray/grey   brown
darkred  gold     silver    navy        teal
olive    lime     indigo    violet      maroon
beige    coral    crimson   khaki

default  reset
```

---

## 💾 Output Flag (`--output`)

Saves the ASCII art to a `.txt` file instead of printing to the terminal. Color and alignment are **not** applied when writing to a file — plain ASCII only.

```bash
# Save to a specific file
go run ./... --output=result.txt "Hello World"

# Save with a specific banner
go run ./... --output=result.txt "Hello World" shadow

# Combine with other flags
go run ./... --output=result.txt "All these words" stylish

# No filename provided — saves to output.txt automatically
go run ./... --output "Hello"
```

> The file **must** end with `.txt`. Any other extension will cause an error.

---

## 🔄 Reverse Flag (`--reverse`)

Reads an existing `.txt` file containing ASCII art and attempts to decode it back to the original text string.

```bash
# Reverse a previously generated file
go run ./... --reverse=output.txt

# Reverse any compatible ASCII art file
go run ./... --reverse=result.txt
```

> When `--reverse` is used, no `TEXT` argument is needed or expected.

---

## 🔤 Banner (Font) Selection

The banner (font) is passed as a plain positional argument at the end. If none is specified, `standard` is used by default.

```bash
go run ./... "Hello" standard
go run ./... "Hello" shadow
go run ./... "Hello" thinkertoy
go run ./... "Hello" extra
go run ./... "Hello" stylish
go run ./... "Hello" bloody
```

### Available banners

| Name | Style |
|---|---|
| `standard` | Clean block letters (default) |
| `shadow` | Letters with drop shadow |
| `thinkertoy` | Rounded, playful style |
| `extra` | Extended custom characters |
| `stylish` | Elegant, condensed look |
| `bloody` | Graffiti, horror-style font |

### Multiple fonts

If you pass more than one font name, one is **selected at random** each run:

```bash
# Randomly picks either shadow or bloody
go run ./... "Hello" shadow bloody
```

---

## ✏️ Text Input

### Multiple words

Words are collected and joined with a space. Quotes are optional if words don't contain spaces.

```bash
go run ./... Hello World
# same as:
go run ./... "Hello World"
```

### Newlines (`\n`)

Use `\n` inside a quoted string to insert line breaks in the output.

```bash
go run ./... "Hello\nWorld"
go run ./... "Line one\nLine two\nLine three"
```

### Tabs (`\t`)

Use `\t` inside a quoted string. Each tab is rendered as 3 spaces.

```bash
go run ./... "Name:\tJohn"
```

### Non-printable characters

Any character outside the printable ASCII range (space to `~`) is **automatically omitted** with a console warning. The rest of the text renders normally.

---

## 🔀 Combining Flags

All flags can be combined freely.

```bash
# Colored, centered, saved to file
go run ./... --color=gold --align=center --output=banner.txt "Zone01" stylish

# Two colors, right-aligned, thinkertoy font
go run ./... --color1=red --substring1=He --color2=cyan --substring2=llo --align=right "Hello" thinkertoy

# Multi-word input, shadow font, blue color on a substring
go run ./... --color=blue --substring=world hello world shadow

# Full example with newlines and alignment
go run ./... --align=center --color=brightgreen "Welcome\nTo\nAscii Art" bloody
```

---

## 📁 Project Structure

```text
ascii/
├── cmd/
│   └── main.go               # Entry point
├── internal/
│   ├── arghandler/           # Flag parsing, Config struct
│   ├── color/                # Color codes parsing (hex, rgb, hsl, normal text)
│   ├── console/              # Read console width
│   ├── font/                 # Banner loading and rendering, reverse and align
│   └── normalize/            # Text sanitization
├── banners/                  # .txt banner files
│   ├── standard.txt
│   ├── shadow.txt
│   ├── thinkertoy.txt
│   ├── extra.txt
│   ├── stylish.txt
│   └── bloody.txt
└── logs.txt                  # Auto-generated runtime logs
```

---

## ⚠️ Common Errors

| Situation | What happens |
|---|---|
| No text provided | Error message + usage hint, exits |
| `--color` without enough substrings | Error message counting missing args, exits |
| `--align` with invalid value | Error message + example, exits |
| `--output` with non-`.txt` filename | Error message, exits |
| `--reverse` with no filename or bad path | Error message + usage hint, exits |
| Unknown characters in input | Warning printed, characters silently omitted |

---

## 🔨 Build

```bash
# Build the executable
go build -o ascii ./cmd/main.go

# Run the compiled binary (same flags apply)
./ascii "Hello World"
./ascii --color=red --align=center "Hello" shadow
./ascii --output=result.txt "Hello World" stylish
```

> The binary can be moved anywhere on your system. Run it from the directory where the `banners/` folder is located, or from the repo root.

---

## 📝 Authors

**Charles Newman** | cnewman | Cohort 2.3 | Zone01 Athens

**Nikos Antoniou** | nantoniou | Cohort 2.3 | Zone01 Athens

**Pavlos Avgerinos** | pavgerin | Cohort 2.3 | Zone01 Athens

---

## 📄 License

Completed as part of the Zone01 Athens curriculum — Cohort 2.3.