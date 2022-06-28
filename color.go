package lipgloss

// TerminalColor is a color intended to be rendered in the terminal.
type TerminalColor interface {
	colorType()
}

// NoColor is used to specify the absence of color styling. When this is active
// foreground colors will be rendered with the terminal's default text color,
// and background colors will not be drawn at all.
//
// Example usage:
//
//     var style = someStyle.Copy().Background(lipgloss.NoColor{})
//
type NoColor struct{}

func (NoColor) colorType() {}

var noColor = NoColor{}

// Color specifies a color by hex or ANSI value. For example:
//
//     ansiColor := lipgloss.Color("21")
//     hexColor := lipgloss.Color("#0000ff")
//
type Color string

func (Color) colorType() {}

// ANSIColor is a color specified by an ANSI color value.
type ANSIColor uint

func (ANSIColor) colorType() {}

// AdaptiveColor provides color options for light and dark backgrounds. The
// appropriate color will be returned at runtime based on the darkness of the
// terminal background color.
//
// Example usage:
//
//     color := lipgloss.AdaptiveColor{Light: "#0000ff", Dark: "#000099"}
//
type AdaptiveColor struct {
	Light string
	Dark  string
}

func (AdaptiveColor) colorType() {}
