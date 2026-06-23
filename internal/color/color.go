package color

import (
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// ── color dictionary and global configurations ───────────────────────────── ⊃

// ColorMap stores the mapping between color names and their respective HEX values.
// This allows uniform HTML span generation across all input formats.
var ColorMap = map[string]string{
	// Standard colors
	"black":  "#000000",
	"red":    "#ff0000",
	"green":  "#008000",
	"yellow": "#ffff00",
	"blue":   "#0000ff",
	"purple": "#800080",
	"cyan":   "#00ffff",
	"white":  "#ffffff",

	// Bold/Bright colors mapped to clear hex variations
	"bold":         "", // Handled as reset/ignore or custom fallback if needed
	"brightblack":  "#555555",
	"brightred":    "#ff5555",
	"brightgreen":  "#55ff55",
	"brightyellow": "#ffff55",
	"brightblue":   "#5555ff",
	"brightpurple": "#ff55ff",
	"brightcyan":   "#55ffff",
	"brightwhite":  "#ffffff",

	// Additional common colors (extended set converted to standard web hex codes)
	"orange":  "#ff8700",
	"pink":    "#ff5faf",
	"magenta": "#ff00ff",
	"gray":    "#808080",
	"grey":    "#808080",
	"brown":   "#875f00",
	"darkred": "#870000",
	"gold":    "#ffd700",
	"silver":  "#c0c0c0",
	"navy":    "#000080",
	"teal":    "#008080",
	"olive":   "#808000",
	"lime":    "#00ff00",
	"indigo":  "#4b0082",
	"violet":  "#ee82ee",
	"maroon":  "#800000",
	"beige":   "#f5f5dc",
	"coral":   "#ff7f50",
	"crimson": "#dc143c",
	"khaki":   "#f0e68c",

	// Utility codes mapped to closing tags or standard text colors
	"default": "</span>",
	"reset":   "</span>",
}

// ── main entry point for color parsing ───────────────────────────────────── ⊃

// ColorConvert routes the input color string to the appropriate parsing function
// based on its prefix (#, rgb(, hsl() or plain text) and returns an HTML opening span tag.
// Logs an error and falls back to an empty string if parsing fails.
func ColorConvert(s string) (string, error) {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	if s == "random" {
		R, G, B := rand.Intn(256), rand.Intn(256), rand.Intn(256)
		randomcolor := fmt.Sprintf("%d,%d,%d)", R, G, B)

		for mean := (R + G + B) / 3; R+G+B < 120 || R+G+B > 650 || (math.Abs(float64(R-mean)) < 30 && math.Abs(float64(G-mean)) < 30 && math.Abs(float64(B-mean)) < 30); randomcolor, mean = fmt.Sprintf("%d,%d,%d)", R, G, B), (R+G+B)/3 {
			R, G, B = rand.Intn(256), rand.Intn(256), rand.Intn(256)
		}
		return ColorRGB(randomcolor)
	}

	// Handle direct utility closure maps immediately
	if s == "default" || s == "reset" {
		return "</span>", nil
	}

	var result string
	var err error
	if temp, ok := strings.CutPrefix(s, "#"); ok {
		result, err = ColorHEX(temp)
	} else if temp, ok := strings.CutPrefix(s, "rgb("); ok {
		result, err = ColorRGB(temp)
	} else if temp, ok := strings.CutPrefix(s, "hsl("); ok {
		result, err = ColorHSL(temp)
	} else {
		result, err = ColorBasic(s)
	}

	if err != nil {
		slog.Error(err.Error(), "Skipping color", s)
		return "", err
	}
	return result, err
}

// ── specific format color converters ─────────────────────────────────────── ⊃

// ColorBasic looks up a plain text color name in the pre-defined ColorMap.
// Returns the corresponding HTML opening span tag.
func ColorBasic(s string) (string, error) {
	hexVal, ok := ColorMap[s]
	if !ok {
		return "", fmt.Errorf("Bad Format: Wrong Color Name = %s", s)
	}

	// If it maps to a closing tag token directly
	if hexVal == "</span>" {
		return hexVal, nil
	}

	return fmt.Sprintf("<span style=\"color:%s\">", hexVal), nil
}

// ColorHEX validates and converts hexadecimal color strings into HTML style span components.
func ColorHEX(s string) (string, error) {
	if len(s) == 3 {
		s = string([]byte{s[0], s[0], s[1], s[1], s[2], s[2]})
	}

	if len(s) != 6 {
		return "", fmt.Errorf("Bad Format: Wrong Color Code length")
	}

	// Validate hex characters by attempting to parse them
	_, errR := strconv.ParseInt(s[0:2], 16, 0)
	_, errG := strconv.ParseInt(s[2:4], 16, 0)
	_, errB := strconv.ParseInt(s[4:6], 16, 0)

	if errR != nil || errG != nil || errB != nil {
		return "", fmt.Errorf("Bad Format: Wrong Color Code = #%s", s)
	}

	return fmt.Sprintf("<span style=\"color:#%s\">", s), nil
}

// ColorRGB extracts integer components and generates inline color span markup.
func ColorRGB(s string) (string, error) {
	R, G, B, err := ThreeValuesSplit(s)
	if err != nil || !(R >= 0 && R <= 255) || !(G >= 0 && G <= 255) || !(B >= 0 && B <= 255) {
		return "", fmt.Errorf("Bad Format: Wrong Color Code = %s", s)
	}

	return fmt.Sprintf("<span style=\"color:rgb(%d,%d,%d)\">", R, G, B), nil
}

// ColorHSL extracts HSL numbers and generates standard web CSS functional span values.
func ColorHSL(s string) (string, error) {
	H, S, L, err := ThreeValuesSplit(strings.ReplaceAll(s, "%", ""))
	if err != nil || !(H >= 0 && H <= 360) || !(S >= 0 && S <= 100) || !(L >= 0 && L <= 100) {
		return "", fmt.Errorf("Bad Format: Wrong Color Code = %s", s)
	}

	return fmt.Sprintf("<span style=\"color:hsl(%d,%d%%,%d%%)\">", H, S, L), nil
}

// ── text processing and mathematical helper utilities ────────────────────── ⊃

// ThreeValuesSplit parses the inner functional notation values.
func ThreeValuesSplit(s string) (int, int, int, error) {
	cleaned := strings.TrimSuffix(s, ")")
	output := strings.SplitN(cleaned, ",", 3)

	if len(output) < 3 {
		return 0, 0, 0, fmt.Errorf("Bad Format: Color Code is missing components")
	}

	var nums []int
	for _, numbers := range output {
		num, err := strconv.Atoi(numbers)
		nums = append(nums, num)
		if err != nil {
			return 0, 0, 0, err
		}
	}
	return nums[0], nums[1], nums[2], nil
}

// hslToRgb remains fully functional in background if you ever need to fall back to raw RGB calculations
func hslToRgb(h, s, l float64) (int, int, int) {
	var r, g, b float64

	if s == 0 {
		r = l
		g = l
		b = l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q

		r = hueToRgb(p, q, h/360.0+1.0/3.0)
		g = hueToRgb(p, q, h/360.0)
		b = hueToRgb(p, q, h/360.0-1.0/3.0)
	}

	return int(math.Round(r * 255)), int(math.Round(g * 255)), int(math.Round(b * 255))
}

func hueToRgb(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6.0*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6.0
	}
	return p
}
