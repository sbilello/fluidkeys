package colour

import (
	"strings"
)

func Custom(fgColor string, bgColor string, message string) string {

	var FgColorCode string
	var BgColorCode string

	fgColor = strings.ToLower(fgColor)
	bgColor = strings.ToLower(bgColor)

	switch fgColor {
	case "red":
		FgColorCode = FgRed
	case "black":
		FgColorCode = FgBlack
	case "blue":
		FgColorCode = FgBlue
	case "cyan":
		FgColorCode = FgCyan
	case "magenta":
		FgColorCode = FgMagenta
	case "green":
		FgColorCode = FgGreen
	case "white":
		FgColorCode = FgWhite
	case "yellow":
		FgColorCode = FgYellow
	default:
		FgColorCode = FgWhite
	}

	switch bgColor {
	case "red":
		BgColorCode = BgRed
	case "black":
		BgColorCode = BgBlack
	case "blue":
		BgColorCode = BgBlue
	case "cyan":
		BgColorCode = BgCyan
	case "magenta":
		BgColorCode = BgMagenta
	case "green":
		BgColorCode = BgGreen
	case "white":
		BgColorCode = BgWhite
	case "yellow":
		BgColorCode = BgYellow
	default:
		BgColorCode = BgBlack
	}

	return BgColorCode + FgColorCode + message + Reset
}

func LightBlue(message string) string {
	return Bright + FgBlue + message + Reset
}

func Underline(message string) string {
	return Underscore + message + Reset
}

func Bold(message string) string {
	return Bright + message + Reset
}

func Flash(message string) string {
	return Blink + message + Reset
}

func Inverse(message string) string {
	return Reverse + message + Reset
}

func Highlight(message string) string {
	return BgYellow + FgRed + message + Reset
}

func Important(message string) string {
	return BgRed + FgWhite + message + Reset
}

func Success(message string) string {
	return FgGreen + message + Reset
}

func Info(message string) string {
	return FgBlue + message + Reset
}

func Warn(message string) string {
	return FgYellow + message + Reset
}

func Error(message string) string {
	return FgMagenta + Bright + message + Reset
}

func Green(message string) string {
	return FgGreen + message + Reset
}

func Blue(message string) string {
	return FgBlue + message + Reset
}

func Magenta(message string) string {
	return FgMagenta + message + Reset
}

func Red(message string) string {
	return FgRed + message + Reset
}

func Black(message string) string {
	return BgWhite + FgBlack + message + Reset
}

func Cyan(message string) string {
	return FgCyan + message + Reset
}

func White(message string) string {
	return BgBlack + FgWhite + message + Reset
}

func Yellow(message string) string {
	return FgYellow + message + Reset
}

func Grey(message string) string {
	return Bright + FgBlack + message + Reset
}

// StripAllColourCodes strips all the ANSI colour codes from a string
func StripAllColourCodes(message string) string {
	message = strings.Replace(message, Reset, "", -1)
	message = strings.Replace(message, Bright, "", -1)
	message = strings.Replace(message, Dim, "", -1)
	message = strings.Replace(message, Underscore, "", -1)
	message = strings.Replace(message, Blink, "", -1)
	message = strings.Replace(message, Reverse, "", -1)
	message = strings.Replace(message, Hidden, "", -1)

	message = strings.Replace(message, FgBlack, "", -1)
	message = strings.Replace(message, FgRed, "", -1)
	message = strings.Replace(message, FgGreen, "", -1)
	message = strings.Replace(message, FgYellow, "", -1)
	message = strings.Replace(message, FgBlue, "", -1)
	message = strings.Replace(message, FgMagenta, "", -1)
	message = strings.Replace(message, FgCyan, "", -1)
	message = strings.Replace(message, FgWhite, "", -1)

	message = strings.Replace(message, BgBlack, "", -1)
	message = strings.Replace(message, BgRed, "", -1)
	message = strings.Replace(message, BgGreen, "", -1)
	message = strings.Replace(message, BgYellow, "", -1)
	message = strings.Replace(message, BgBlue, "", -1)
	message = strings.Replace(message, BgMagenta, "", -1)
	message = strings.Replace(message, BgCyan, "", -1)
	message = strings.Replace(message, BgWhite, "", -1)

	return message
}
