package ansi

import "fmt"

// SetFormattingEnabled sets whether or not to enable formatting.
// If disabled, all formatting functions will return the same string as their input.
func SetFormattingEnabled(enabled bool) {
	enableFormatting = enabled
}

func formatText(c Format, args ...interface{}) string {
	return ansi(c) + fmt.Sprint(args...) + ansi(FormatReset)
}

func Red(args ...interface{}) string          { return formatText(ColorRed, args...) }
func Green(args ...interface{}) string        { return formatText(ColorGreen, args...) }
func Yellow(args ...interface{}) string       { return formatText(ColorYellow, args...) }
func Blue(args ...interface{}) string         { return formatText(ColorBlue, args...) }
func Magenta(args ...interface{}) string      { return formatText(ColorMagenta, args...) }
func Cyan(args ...interface{}) string         { return formatText(ColorCyan, args...) }
func Gray(args ...interface{}) string         { return formatText(ColorGray, args...) }
func White(args ...interface{}) string        { return formatText(ColorWhite, args...) }
func BgBlack(args ...interface{}) string      { return formatText(BackgroundBlack, args...) }
func BgRed(args ...interface{}) string        { return formatText(BackgroundRed, args...) }
func BgGreen(args ...interface{}) string      { return formatText(BackgroundGreen, args...) }
func BgYellow(args ...interface{}) string     { return formatText(BackgroundYellow, args...) }
func BgBlue(args ...interface{}) string       { return formatText(BackgroundBlue, args...) }
func BgMagenta(args ...interface{}) string    { return formatText(BackgroundMagenta, args...) }
func BgCyan(args ...interface{}) string       { return formatText(BackgroundCyan, args...) }
func BgWhite(args ...interface{}) string      { return formatText(BackgroundWhite, args...) }
func FmtBold(args ...interface{}) string      { return formatText(Bold, args...) }
func FmtUnderline(args ...interface{}) string { return formatText(Underline, args...) }
func FmtBlink(args ...interface{}) string     { return formatText(Blink, args...) }
func FmtReverse(args ...interface{}) string   { return formatText(Reverse, args...) }
func FmtUnReverse(args ...interface{}) string { return formatText(UnReverse, args...) }
func Reset() string                           { return ansi(FormatReset) }

func CombineFormats(formats ...Format) string {
	var result string
	for _, format := range formats {
		result += ansi(format)
	}
	return result
}

func FormatTextWithStyles(text string, color Format, styles ...Format) string {
	combinedStyles := CombineFormats(styles...)
	return combinedStyles + ansi(color) + text + ansi(FormatReset)
}

func CustomFormat(format string, args ...interface{}) string {
	return format + fmt.Sprint(args...) + ansi(FormatReset)
}
