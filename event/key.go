//CSAkeybindings

package event

var KeyCode = map[int]string{
	0: "KEY_RESERVED",
	1: "KEY_ESC",

	2:  "1", // CSA: "&" is replaced with "1"
	3:  "2", // CSA: "é" is replaced with "2"
	4:  "3", // CSA: "\"" is replaced with "3"
	5:  "4", // CSA: "'" is replaced with "4"
	6:  "5", // CSA: "(" is replaced with "5"
	7:  "6", // CSA: "-" is replaced with "6"
	8:  "7", // CSA: "è" is replaced with "7"
	9:  "8", // CSA: "_" is replaced with "8"
	10: "9", // CSA: "ç" is replaced with "9"
	11: "0", // CSA: "à" is replaced with "0"
	12: "-", // CSA: ")" is replaced with "-"
	13: "=", // CSA: "=" stays "="

	14: "KEY_BACKSPACE",
	15: "KEY_TAB",

	16: "q", // CSA: "a" is replaced with "q"
	17: "w", // CSA: "z" is replaced with "w"
	18: "e", // CSA: "e" stays "e"
	19: "r", // CSA: "r" stays "r"
	20: "t", // CSA: "t" stays "t"
	21: "y", // CSA: "y" stays "y"
	22: "u", // CSA: "u" stays "u"
	23: "i", // CSA: "i" stays "i"
	24: "o", // CSA: "o" stays "o"
	25: "p", // CSA: "p" stays "p"
	26: "^", // CSA: "^" is replaced with "["
	27: "ç", // CSA: "$" is replaced with "]"
	28: "KEY_ENTER",
	29: "KEY_L_CTRL",

	30: "a", // CSA: "q" is replaced with "a"
	31: "s", // CSA: "s" stays "s"
	32: "d", // CSA: "d" stays "d"
	33: "f", // CSA: "f" stays "f"
	34: "g", // CSA: "g" stays "g"
	35: "h", // CSA: "h" stays "h"
	36: "j", // CSA: "j" stays "j"
	37: "k", // CSA: "k" stays "k"
	38: "l", // CSA: "l" stays "l"
	39: ";", // CSA: "m" is replaced with ";"
	40: "è", // CSA: "ù" is replaced with "'"
	41: "ù", // CSA: "*" is replaced with "`" // seems to be dead on hp keyboard
	
	// 41 seems to be reverse with 43 on my hp keyboard

	42: "KEY_L_SHIFT",
	43: "à", // CSA: "<" is replaced with "\"
	44: "z", // CSA: "w" is replaced with "z"
	45: "x", // CSA: "x" stays "x"
	46: "c", // CSA: "c" stays "c"
	47: "v", // CSA: "v" stays "v"
	48: "b", // CSA: "b" stays "b"
	49: "n", // CSA: "n" stays "n"
	50: "m", // CSA: "," stays ","
	51: ",", // CSA: ";" is replaced with "."
	52: ".", // CSA: ":" is replaced with "/"
	53: "é", // CSA: "!" is replaced with "KEY_R_SHIFT"

	55: "KEY_KPASTERISK",
	56: "KEY_L_ALT",

	57: "KEY_SPACE",
	58: "KEY_CAPSLOCK",
	59: "KEY_F1",
	60: "KEY_F2",
	61: "KEY_F3",
	62: "KEY_F4",
	63: "KEY_F5",
	64: "KEY_F6",
	65: "KEY_F7",
	66: "KEY_F8",
	67: "KEY_F9",
	68: "KEY_F10",

	87: "KEY_F11",
	88: "KEY_F12",

	100: "KEY_ALT_GR",

	103: "KEY_UP",
	105: "KEY_LEFT",
	106: "KEY_RIGHT",
	108: "KEY_DOWN",

	111: "KEY_DEL",

	183: "KEY_F13",
	184: "KEY_F14",
	185: "KEY_F15",
	186: "KEY_F16",
	187: "KEY_F17",
	188: "KEY_F18",
	189: "KEY_F19",
	190: "KEY_F20",
	191: "KEY_F21",
	192: "KEY_F22",
	193: "KEY_F23",
	194: "KEY_F24",
}

var KeyCodeMaj = map[int]string{
	2:  "!",
	3:  "\"",
	4:  "#",
	5:  "$",
	6:  "%",
	7:  "?",
	8:  "&",
	9:  "*",
	10: "(",
	11: ")",
	12: "_",
	13: "+",

	16: "Q",
	17: "W",
	18: "E",
	19: "R",
	20: "T",
	21: "Y",
	22: "U",
	23: "I",
	24: "O",
	25: "P",
	26: "¨",
	27: "Ç",

	30: "A",
	31: "S",
	32: "D",
	33: "F",
	34: "G",
	35: "H",
	36: "J",
	37: "K",
	38: "L",
	39: ":",
	40: "È",
	41: "Ù",

	43: "À",
	44: "Z",
	45: "X",
	46: "C",
	47: "V",
	48: "B",
	49: "N",
	50: "M",
	51: "'",
	52: ".",
	53: "É",
}

var KeyCodeAlt = map[int]string{
	1:  "\\",
	2:  "¹",
	3:  "@",
	4:  "³",
	5:  "¼",
	6:  "½",
	7:  "¾",
	8:  "{",
	9:  "[",
	10: "]",
	11: "}",
	12: "|",
	13: "¸",

	16: "q",
	17: "w",
	18: "e",
	19: "¶",
	20: "t",
	21: "¥",
	22: "u",
	23: "i",
	24: "ø",
	25: "þ",
	26: "°",
	27: "",
	
	30: "æ", // CSA: "q" is replaced with "a"
	31: "ß", // CSA: "s" stays "s"
	32: "ð", // CSA: "d" stays "d"
	33: "ª", // CSA: "f" stays "f"
	34: "g", // CSA: "g" stays "g"
	35: "h", // CSA: "h" stays "h"
	36: "j", // CSA: "j" stays "j"
	37: "k", // CSA: "k" stays "k"
	38: "l", // CSA: "l" stays "l"
	39: "´", // CSA: "m" is replaced with ";"
	40: "{", // CSA: "ù" is replaced with "'"
	41: "¬", // CSA: "*" is replaced with "`"


	43: "`", // CSA: "<" is replaced with "\"
	44: "«", // CSA: "w" is replaced with "z"
	45: "»", // CSA: "x" stays "x"
	46: "¢", // CSA: "c" stays "c"
	47: "v", // CSA: "v" stays "v"
	48: "b", // CSA: "b" stays "b"
	49: "n", // CSA: "n" stays "n"
	50: "µ", // CSA: "," stays ","
	51: "<", // CSA: ";" is replaced with "."
	52: ">", // CSA: ":" is replaced with "/"
	53: "//", // CSA: "!" is replaced with "KEY_R_SHIFT"

}

var KeyCodeAltGr = map[int]string{
	1:  "\\",
	2:  "¹",
	3:  "@",
	4:  "³",
	5:  "¼",
	6:  "½",
	7:  "¾",
	8:  "{",
	9:  "[",
	10: "]",
	11: "}",
	12: "|",
	13: "¸",

	16: "q",
	17: "w",
	18: "e",
	19: "¶",
	20: "t",
	21: "¥",
	22: "u",
	23: "i",
	24: "ø",
	25: "þ",
	26: "°",
	27: "",
	
	30: "æ", // CSA: "q" is replaced with "a"
	31: "ß", // CSA: "s" stays "s"
	32: "ð", // CSA: "d" stays "d"
	33: "ª", // CSA: "f" stays "f"
	34: "g", // CSA: "g" stays "g"
	35: "h", // CSA: "h" stays "h"
	36: "j", // CSA: "j" stays "j"
	37: "k", // CSA: "k" stays "k"
	38: "l", // CSA: "l" stays "l"
	39: "´", // CSA: "m" is replaced with ";"
	40: "{", // CSA: "ù" is replaced with "'"
	41: "¬", // CSA: "*" is replaced with "`"


	43: "`", // CSA: "<" is replaced with "\"
	44: "«", // CSA: "w" is replaced with "z"
	45: "»", // CSA: "x" stays "x"
	46: "¢", // CSA: "c" stays "c"
	47: "v", // CSA: "v" stays "v"
	48: "b", // CSA: "b" stays "b"
	49: "n", // CSA: "n" stays "n"
	50: "µ", // CSA: "," stays ","
	51: "<", // CSA: ";" is replaced with "."
	52: ">", // CSA: ":" is replaced with "/"
	53: "//", // CSA: "!" is replaced with "KEY_R_SHIFT"

}