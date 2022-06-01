package lipgloss

import (
	"testing"

	"github.com/muesli/termenv"
)

func TestSetColorProfile(t *testing.T) {
	input := "hello"

	tt := []struct {
		name     string
		profile  termenv.Profile
		expected string
	}{
		{
			"ascii",
			termenv.Ascii,
			"hello",
		},
		{
			"ansi",
			termenv.ANSI,
			"\x1b[94mhello\x1b[0m",
		},
		{
			"ansi256",
			termenv.ANSI256,
			"\x1b[38;5;62mhello\x1b[0m",
		},
		{
			"truecolor",
			termenv.TrueColor,
			"\x1b[38;2;89;86;224mhello\x1b[0m",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			SetColorProfile(tc.profile)
			style := NewStyle().Foreground(Color("#5A56E0"))
			res := style.Render(input)

			if res != tc.expected {
				t.Errorf("Expected:\n\n`%s`\n`%s`\n\nActual output:\n\n`%s`\n`%s`\n\n",
					tc.expected, formatEscapes(tc.expected),
					res, formatEscapes(res))
			}
		})
	}
}
