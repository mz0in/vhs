package dolly

import (
	"time"

	"github.com/go-rod/rod/lib/input"
)

// Type types the given string onto the page at the given speed. The delay is
// the time between each key press.
func (d Dolly) Type(str string, delay time.Duration) {
	for _, r := range str {
		k, ok := keymap[r]
		if ok {
			d.Page.Keyboard.Type(k)
		} else {
			d.Page.MustElement("textarea").Input(string(r))
			d.Page.MustWaitIdle()
		}
		time.Sleep(delay)
	}
}

// Enter is a helper function that press the enter key.
func (d Dolly) Enter() { d.Page.Keyboard.Type(input.Enter) }

func shift(k input.Key) input.Key {
	k, _ = k.Shift()
	return k
}

var keymap = map[rune]input.Key{
	' ':    input.Space,
	'!':    shift(input.Digit1),
	'"':    shift(input.Quote),
	'#':    shift(input.Digit3),
	'$':    shift(input.Digit4),
	'%':    shift(input.Digit5),
	'&':    shift(input.Digit7),
	'(':    shift(input.Digit9),
	')':    shift(input.Digit0),
	'*':    shift(input.Digit8),
	'+':    shift(input.Equal),
	',':    input.Comma,
	'-':    input.Minus,
	'.':    input.Period,
	'/':    input.Slash,
	'0':    input.Digit0,
	'1':    input.Digit1,
	'2':    input.Digit2,
	'3':    input.Digit3,
	'4':    input.Digit4,
	'5':    input.Digit5,
	'6':    input.Digit6,
	'7':    input.Digit7,
	'8':    input.Digit8,
	'9':    input.Digit9,
	':':    shift(input.Semicolon),
	';':    input.Semicolon,
	'<':    shift(input.Comma),
	'=':    input.Equal,
	'>':    shift(input.Period),
	'?':    shift(input.Slash),
	'@':    shift(input.Digit2),
	'A':    shift(input.KeyA),
	'B':    shift(input.KeyB),
	'C':    shift(input.KeyC),
	'D':    shift(input.KeyD),
	'E':    shift(input.KeyE),
	'F':    shift(input.KeyF),
	'G':    shift(input.KeyG),
	'H':    shift(input.KeyH),
	'I':    shift(input.KeyI),
	'J':    shift(input.KeyJ),
	'K':    shift(input.KeyK),
	'L':    shift(input.KeyL),
	'M':    shift(input.KeyM),
	'N':    shift(input.KeyN),
	'O':    shift(input.KeyO),
	'P':    shift(input.KeyP),
	'Q':    shift(input.KeyQ),
	'R':    shift(input.KeyR),
	'S':    shift(input.KeyS),
	'T':    shift(input.KeyT),
	'U':    shift(input.KeyU),
	'V':    shift(input.KeyV),
	'W':    shift(input.KeyW),
	'X':    shift(input.KeyX),
	'Y':    shift(input.KeyY),
	'Z':    shift(input.KeyZ),
	'[':    input.BracketLeft,
	'\'':   input.Quote,
	'\\':   input.Backslash,
	'\b':   input.Backspace,
	'\n':   input.Enter,
	'\r':   input.Enter,
	'\t':   input.Tab,
	'\x1b': input.Escape,
	']':    input.BracketRight,
	'^':    shift(input.Digit6),
	'_':    shift(input.Minus),
	'`':    input.Backquote,
	'a':    input.KeyA,
	'b':    input.KeyB,
	'c':    input.KeyC,
	'd':    input.KeyD,
	'e':    input.KeyE,
	'f':    input.KeyF,
	'g':    input.KeyG,
	'h':    input.KeyH,
	'i':    input.KeyI,
	'j':    input.KeyJ,
	'k':    input.KeyK,
	'l':    input.KeyL,
	'm':    input.KeyM,
	'n':    input.KeyN,
	'o':    input.KeyO,
	'p':    input.KeyP,
	'q':    input.KeyQ,
	'r':    input.KeyR,
	's':    input.KeyS,
	't':    input.KeyT,
	'u':    input.KeyU,
	'v':    input.KeyV,
	'w':    input.KeyW,
	'x':    input.KeyX,
	'y':    input.KeyY,
	'z':    input.KeyZ,
	'{':    shift(input.BracketLeft),
	'|':    shift(input.Backslash),
	'}':    shift(input.BracketRight),
	'~':    shift(input.Backquote),
	'←':    input.ArrowLeft,
	'↑':    input.ArrowUp,
	'→':    input.ArrowRight,
	'↓':    input.ArrowDown,
}
