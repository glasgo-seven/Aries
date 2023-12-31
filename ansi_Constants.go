package Aries


const (
	SY_BORDER_VER_SIDE			string	=	"║"
	SY_BORDER_VER_SIDE_LEFT		string	=	"╣"
	SY_BORDER_VER_SIDE_RIGHT	string	=	"╠"
	SY_BORDER_HOR_SIDE			string	=	"═"
	SY_BORDER_HOR_SIDE_UP		string	=	"╩"
	SY_BORDER_HOR_SIDE_DOWN		string	=	"╦"
	SY_BORDER_CROSS				string	=	"╬"
	SY_BORDER_CORNER_RIGHT_UP	string	=	"╗"
	SY_BORDER_CORNER_RIGHT_DOWN	string	=	"╝"
	SY_BORDER_CORNER_LEFT_DOWN	string	=	"╚"
	SY_BORDER_CORNER_LEFT_UP	string	=	"╔"

	ESC		string	=	"\x1b["

	SCREEN_HOME		string	=	"H"
	LINE_HOME		string	=	"0G"

	MOVE_N_ROW_U	string	=	"A"
	MOVE_N_ROW_D	string	=	"B"

	MOVE_N_COL_R	string	=	"C"
	MOVE_N_COL_L	string	=	"D"

	MOVE_NEXT_LINE	string	=	"E"
	MOVE_PREV_LINE	string	=	"F"

	MOVE_TO_COL_N	string	=	"G"

	CLEAR_SCREEN_TO_END		string	=	"1J"
	CLEAR_SCREEN			string	=	"2J"

	CLEAR_LINE_END		string	=	"0K"
	CLEAR_LINE_BEGIN	string	=	"1K"
	CLEAR_LINE			string	=	"2K"

	STYLE				string	=	"m"
	STYLE_RESET			string	=	"0m"

	STYLE_BOLD			string	=	"1m"
	STYLE_BOLD_RESET	string	=	"21m"

	STYLE_DIM			string	=	"2m"
	STYLE_DIM_RESET		string	=	"22m"

	STYLE_ITALIC		string	=	"3m"
	STYLE_ITALIC_RESET	string	=	"23m"

	STYLE_UNDER			string	=	"4m"
	STYLE_UNDER_RESET	string	=	"24m"

	STYLE_BLINK			string	=	"5m"
	STYLE_BLINK_RESET	string	=	"25m"

	STYLE_INVERSE		string	=	"7m"
	STYLE_INVERSE_RESET	string	=	"27m"

	STYLE_HIDDEN		string	=	"8m"
	STYLE_HIDDEN_RESET	string	=	"28m"

	STYLE_STRIKE		string	=	"9m"
	STYLE_STRIKE_RESET	string	=	"29m"

	STYLE_COLOR_RESET	string	=	"0"

	STYLE_COLOR_FG_BLACK	string	=	"30"
	STYLE_COLOR_FG_RED		string	=	"31"
	STYLE_COLOR_FG_GREEN	string	=	"32"
	STYLE_COLOR_FG_YELLOW	string	=	"33"
	STYLE_COLOR_FG_BLUE		string	=	"34"
	STYLE_COLOR_FG_MAGENTA	string	=	"35"
	STYLE_COLOR_FG_CYAN		string	=	"36"
	STYLE_COLOR_FG_WHITE	string	=	"37"
	STYLE_COLOR_FG_DEFAULT	string	=	"39"

	STYLE_COLOR_FG_BRIGHT_BLACK		string	=	"90"
	STYLE_COLOR_FG_BRIGHT_RED		string	=	"91"
	STYLE_COLOR_FG_BRIGHT_GREEN		string	=	"92"
	STYLE_COLOR_FG_BRIGHT_YELLOW	string	=	"93"
	STYLE_COLOR_FG_BRIGHT_BLUE		string	=	"94"
	STYLE_COLOR_FG_BRIGHT_MAGENTA	string	=	"95"
	STYLE_COLOR_FG_BRIGHT_CYAN		string	=	"96"
	STYLE_COLOR_FG_BRIGHT_WHITE		string	=	"97"

	STYLE_COLOR_BG_BLACK	string	=	"40"
	STYLE_COLOR_BG_RED		string	=	"41"
	STYLE_COLOR_BG_GREEN	string	=	"42"
	STYLE_COLOR_BG_YELLOW	string	=	"43"
	STYLE_COLOR_BG_BLUE		string	=	"44"
	STYLE_COLOR_BG_MAGENTA	string	=	"45"
	STYLE_COLOR_BG_CYAN		string	=	"46"
	STYLE_COLOR_BG_WHITE	string	=	"47"

	STYLE_COLOR_BG_BRIGHT_BLACK		string	=	"100"
	STYLE_COLOR_BG_BRIGHT_RED		string	=	"101"
	STYLE_COLOR_BG_BRIGHT_GREEN		string	=	"102"
	STYLE_COLOR_BG_BRIGHT_YELLOW	string	=	"103"
	STYLE_COLOR_BG_BRIGHT_BLUE		string	=	"104"
	STYLE_COLOR_BG_BRIGHT_MAGENTA	string	=	"105"
	STYLE_COLOR_BG_BRIGHT_CYAN		string	=	"106"
	STYLE_COLOR_BG_BRIGHT_WHITE		string	=	"107"

	STYLE_COLOR_256_FG		string	=	"38;5;"
	STYLE_COLOR_256_BG		string	=	"48;5;"

	STYLE_COLOR_RGB_FG		string	=	"38;2;"
	STYLE_COLOR_RGB_BG		string	=	"48;2;"
)
