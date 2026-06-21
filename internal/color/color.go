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

// ColorMap stores the mapping between color names and their respective ANSI escape sequences.
// By keeping this outside the function, we avoid re-allocating the map on every function call.
var ColorMap = map[string]string{
	// Standard colors
	"black":  "\033[0;30m",
	"red":    "\033[0;31m",
	"green":  "\033[0;32m",
	"yellow": "\033[0;33m",
	"blue":   "\033[0;34m",
	"purple": "\033[0;35m",
	"cyan":   "\033[0;36m",
	"white":  "\033[0;37m",

	// Bold/Bright colors
	"bold":         "\033[1m",
	"brightblack":  "\033[1;30m",
	"brightred":    "\033[1;31m",
	"brightgreen":  "\033[1;32m",
	"brightyellow": "\033[1;33m",
	"brightblue":   "\033[1;34m",
	"brightpurple": "\033[1;35m",
	"brightcyan":   "\033[1;36m",
	"brightwhite":  "\033[1;37m",

	// Additional common colors (extended set)
	"orange":  "\033[38;5;208m",
	"pink":    "\033[38;5;205m",
	"magenta": "\033[0;35m",
	"gray":    "\033[38;5;244m",
	"grey":    "\033[38;5;244m",
	"brown":   "\033[38;5;94m",
	"darkred": "\033[38;5;88m",
	"gold":    "\033[38;5;220m",
	"silver":  "\033[38;5;7m",
	"navy":    "\033[38;5;18m",
	"teal":    "\033[38;5;30m",
	"olive":   "\033[38;5;58m",
	"lime":    "\033[38;5;10m",
	"indigo":  "\033[38;5;54m",
	"violet":  "\033[38;5;170m",
	"maroon":  "\033[38;5;1m",
	"beige":   "\033[38;5;230m",
	"coral":   "\033[38;5;203m",
	"crimson": "\033[38;5;160m",
	"khaki":   "\033[38;5;185m",

	// Utility codes
	"default": "\033[0m",
	"reset":   "\033[0m",
}

// ── main entry point for color parsing ───────────────────────────────────── ⊃

// ColorConvert routes the input color string to the appropriate parsing function
// based on its prefix (#, rgb(, hsl() or plain text) and returns an ANSI escape sequence.
// Logs an error and falls back to a terminal reset code if parsing fails.
func ColorConvert(s string) (string, error) {
	if strings.ToLower(s) == "random" {
		R, G, B := rand.Intn(256), rand.Intn(256), rand.Intn(256)
		randomcolor := fmt.Sprintf("%d,%d,%d)", R, G, B)

		for mean := (R + G + B) / 3; R+G+B < 120 || R+G+B > 650 || (math.Abs(float64(R-mean)) < 30 && math.Abs(float64(G-mean)) < 30 && math.Abs(float64(B-mean)) < 30); randomcolor, mean = fmt.Sprintf("%d,%d,%d)", R, G, B), (R+G+B)/3 {
			fmt.Printf("\033[48;2;%d;%d;%dm We skipped this boring random color(%s!%s\n", R, G, B, randomcolor, ColorMap["reset"])
			R, G, B = rand.Intn(256), rand.Intn(256), rand.Intn(256)
		}
		fmt.Printf("\033[48;2;%d;%d;%dm You picked a random color!! We will choose for you the color rgb(%s!%s\n", R, G, B, randomcolor, ColorMap["reset"])
		return ColorRGB(randomcolor)
	}
	var result string
	var err error
	s = strings.ReplaceAll(s, " ", "")
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
		return "\033[0m", err
	}
	return result, err
}

// ── specific format color converters ─────────────────────────────────────── ⊃

// ColorBasic looks up a plain text color name in the pre-defined ColorMap.
// Returns the corresponding ANSI escape sequence, or a terminal reset sequence
// along with an error if the color name is unrecognized.
func ColorBasic(s string) (string, error) {
	s = strings.ToLower(s)
	result, ok := ColorMap[s]
	if !ok {
		error := fmt.Errorf("Bad Format: Wrong Color Code = %s", s)
		return "\033[0m", error
	}

	return result, nil
}

// ColorHEX converts a 3-character or 6-character hexadecimal color string into Truecolor ANSI format.
// Validates the string length and parses the hexadecimal pairs into red, green, and blue integers.
func ColorHEX(s string) (string, error) {
	if len(s) == 3 {
		s = string([]byte{s[0], s[0], s[1], s[1], s[2], s[2]})
	}

	if len(s) != 6 {
		return "\033[0m", fmt.Errorf("Bad Format: Wrong Color Code")
	}

	// Parse and convert (to integer) the color components
	r, errR := strconv.ParseInt(s[0:2], 16, 0)
	g, errG := strconv.ParseInt(s[2:4], 16, 0)
	b, errB := strconv.ParseInt(s[4:6], 16, 0)

	// Check for errors
	if errR != nil || errG != nil || errB != nil {
		return "\033[0m", fmt.Errorf("Bad Format: Wrong Color Code = #%s", s)
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), nil
}

// ColorRGB extracts and validates individual integer components from an RGB functional notation string.
// Returns an error if any component falls outside the valid 8-bit channel range (0-255).
func ColorRGB(s string) (string, error) {
	R, G, B, err := ThreeValuesSplit(s)
	if err != nil || !(R >= 0 && R <= 255) || !(G >= 0 && G <= 255) || !(B >= 0 && B <= 255) {
		error := fmt.Errorf("Bad Format: Wrong Color Code = %s", s)
		return "\033[0m", error
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", R, G, B), nil
}

// ColorHSL extracts HSL components from the input string, validates their operational boundaries,
// converts the HSL values into the RGB color space, and structures the final Truecolor ANSI sequence.
func ColorHSL(s string) (string, error) {
	H, S, L, err := ThreeValuesSplit(strings.ReplaceAll(s, "%", ""))
	if err != nil || !(H >= 0 && H <= 360) || !(S >= 0 && S <= 100) || !(L >= 0 && L <= 100) {
		error := fmt.Errorf("Bad Format: Wrong Color Code = %s", s)
		return "\033[0m", error
	}
	r, g, b := hslToRgb(float64(H), float64(S)/100.0, float64(L)/100.0)
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b), nil
}

// ── text processing and mathematical helper utilities ────────────────────── ⊃

// ThreeValuesSplit strips the closing parenthesis from a comma-separated string, splits it into three sub-strings,
// and parses each slice element into an integer representation, returning an error on structural mismatches.
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

// hslToRgb converts Hue, Saturation, and Lightness floating-point values into 8-bit RGB integers.
// Uses an intermediary hue conversion process and handles achromatic cases where saturation is zero.
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

// hueToRgb is a specialized mathematical utility function that calculates the channel contribution
// for a single color dimension based on the relative position of the hue on the color wheel.
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
