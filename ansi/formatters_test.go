package ansi

import (
	"testing"
)

func TestAnsi(t *testing.T) {
	// Test empty cases.
	for _, f := range []Format{formats_start, formats_start - 1, formats_end, formats_end + 1, FormatNone} {
		s := ansi(f)
		if s != "" {
			t.Errorf("expected empty string, got %q", s)
		}
	}

	// Test specific cases.
	if s := ansi(FormatReset); s != "\033[0m" {
		t.Errorf("expected %q, got %q", "\033[0m", s)
	}
	if s := ansi(UnReverse); s != "\033[27m" {
		t.Errorf("expected %q, got %q", "\033[27m", s)
	}

	// Test all valid formats.
	for f, expected := range FormatStrings {
		s := ansi(f)
		if s != expected {
			t.Errorf("for format %v, expected %q, got %q", f, expected, s)
		}
	}
}

func TestAnsiFormatting(t *testing.T) {
	text := "some text for CLI"

	tests := []struct {
		name     string
		formatFn func(...interface{}) string
		expected string
	}{
		{"Red", Red, ansi(ColorRed) + text + ansi(FormatReset)},
		{"Green", Green, ansi(ColorGreen) + text + ansi(FormatReset)},
		{"Yellow", Yellow, ansi(ColorYellow) + text + ansi(FormatReset)},
		{"Blue", Blue, ansi(ColorBlue) + text + ansi(FormatReset)},
		{"Magenta", Magenta, ansi(ColorMagenta) + text + ansi(FormatReset)},
		{"Cyan", Cyan, ansi(ColorCyan) + text + ansi(FormatReset)},
		{"Gray", Gray, ansi(ColorGray) + text + ansi(FormatReset)},
		{"White", White, ansi(ColorWhite) + text + ansi(FormatReset)},
		{"BgBlack", BgBlack, ansi(BackgroundBlack) + text + ansi(FormatReset)},
		{"BgRed", BgRed, ansi(BackgroundRed) + text + ansi(FormatReset)},
		{"BgGreen", BgGreen, ansi(BackgroundGreen) + text + ansi(FormatReset)},
		{"BgYellow", BgYellow, ansi(BackgroundYellow) + text + ansi(FormatReset)},
		{"BgBlue", BgBlue, ansi(BackgroundBlue) + text + ansi(FormatReset)},
		{"BgMagenta", BgMagenta, ansi(BackgroundMagenta) + text + ansi(FormatReset)},
		{"BgCyan", BgCyan, ansi(BackgroundCyan) + text + ansi(FormatReset)},
		{"BgWhite", BgWhite, ansi(BackgroundWhite) + text + ansi(FormatReset)},
		{"FmtBold", FmtBold, ansi(Bold) + text + ansi(FormatReset)},
		{"FmtUnderline", FmtUnderline, ansi(Underline) + text + ansi(FormatReset)},
		{"FmtBlink", FmtBlink, ansi(Blink) + text + ansi(FormatReset)},
		{"FmtReverse", FmtReverse, ansi(Reverse) + text + ansi(FormatReset)},
		{"FmtUnReverse", FmtUnReverse, ansi(UnReverse) + text + ansi(FormatReset)},
		{"Reset", func(...interface{}) string { return Reset() }, ansi(FormatReset)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if s := tt.formatFn(text); s != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, s)
			}
		})
	}
}
