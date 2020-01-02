package colors

type Color []byte

type Colors struct {
	NONE          Color
	BLACK         Color
	RED           Color
	GREEN         Color
	YELLOW        Color
	BLUE          Color
	MAGENTA       Color
	CYAN          Color
	WHITE         Color
	BLACK_LIGHT   Color
	RED_LIGHT     Color
	GREEN_LIGHT   Color
	YELLOW_LIGHT  Color
	BLUE_LIGHT    Color
	MAGENTA_LIGHT Color
	CYAN_LIGHT    Color
	WHITE_LIGHT   Color
}

/*
 * Foreground Color list
 */
var Fg = Colors{
	NONE:          Color(""),
	BLACK:         Color("\033[0;30m"),
	RED:           Color("\033[0;31m"),
	GREEN:         Color("\033[0;32m"),
	YELLOW:        Color("\033[0;33m"),
	BLUE:          Color("\033[0;34m"),
	MAGENTA:       Color("\033[0;35m"),
	CYAN:          Color("\033[0;36m"),
	WHITE:         Color("\033[0;37m"),
	BLACK_LIGHT:   Color("\033[0;90m"),
	RED_LIGHT:     Color("\033[0;91m"),
	GREEN_LIGHT:   Color("\033[0;92m"),
	YELLOW_LIGHT:  Color("\033[0;93m"),
	MAGENTA_LIGHT: Color("\033[0;94m"),
	BLUE_LIGHT:    Color("\033[0;95m"),
	CYAN_LIGHT:    Color("\033[0;96m"),
	WHITE_LIGHT:   Color("\033[0;97m"),
}

/*
 * Text color the text with the color
 */
func (c Color) Text(data []byte) []byte {
	return append(c, data...)
}
