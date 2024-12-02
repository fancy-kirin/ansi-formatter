package ansi

type Format int

var enableFormatting = true

const (
	formats_start Format = iota

	// Text colors.
	FormatNone
	FormatReset
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorGray
	ColorWhite

	// Background colors.
	BackgroundBlack
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite

	// Text styles.
	Bold
	Underline
	Blink
	Reverse
	UnReverse

	formats_end
)

var FormatStrings = map[Format]string{
	FormatNone:        "",
	FormatReset:       "\033[0m",
	ColorRed:          "\033[31m",
	ColorGreen:        "\033[32m",
	ColorYellow:       "\033[33m",
	ColorBlue:         "\033[34m",
	ColorMagenta:      "\033[35m",
	ColorCyan:         "\033[36m",
	ColorGray:         "\033[37m",
	ColorWhite:        "\033[97m",
	BackgroundBlack:   "\033[40m",
	BackgroundRed:     "\033[41m",
	BackgroundGreen:   "\033[42m",
	BackgroundYellow:  "\033[43m",
	BackgroundBlue:    "\033[44m",
	BackgroundMagenta: "\033[45m",
	BackgroundCyan:    "\033[46m",
	BackgroundWhite:   "\033[47m",
	Bold:              "\033[1m",
	Underline:         "\033[4m",
	Blink:             "\033[5m",
	Reverse:           "\033[7m",
	UnReverse:         "\033[27m",
}

func (c Format) CheckType() bool {
	return c > formats_start && c < formats_end
}

func (c Format) String() string {
	if !c.CheckType() {
		return ""
	}
	return FormatStrings[c]
}

func ansi(c Format) string {
	if !enableFormatting {
		return ""
	}

	return c.String()
}
