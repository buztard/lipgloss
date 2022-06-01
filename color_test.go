package lipgloss

import (
	"testing"

	"github.com/muesli/termenv"
)

func TestSetColorProfile(t *testing.T) {
	t.Parallel()

	style := NewStyle().Foreground(Color("#5A56E0"))
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

	for i, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			SetColorProfile(tc.profile)
			res := style.Render(input)
			if res != tc.expected {
				t.Log(res, "vs", tc.expected)
				t.Errorf("Test %d, expected:\n\n`%s`\n`%s`\n\nActual output:\n\n`%s`\n`%s`\n\n",
					i, tc.expected, formatEscapes(tc.expected),
					res, formatEscapes(res))
			}
		})
	}
}
