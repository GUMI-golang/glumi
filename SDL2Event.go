package glumi

import (
	"github.com/GUMI-golang/gumi"
	"github.com/veandco/go-sdl2/sdl"
	"image"
	"unicode/utf8"
)

type Handler struct {
	glumi  *GLUMIFullScreen
	keymap map[gumi.GUMIKey]gumi.EventKind
}

func (s *Handler) Handle(event sdl.Event) error {
	switch t := event.(type) {
	case *sdl.QuitEvent:
		return Stop
	case *sdl.MouseMotionEvent:
		s.Cursor(t)
	case *sdl.MouseWheelEvent:
		s.Scrool(t)
	case *sdl.MouseButtonEvent:
		s.Mouse(t)
	case *sdl.KeyboardEvent:
		s.Key(t)
	case *sdl.TextEditingEvent:
		s.RuneEdit(t)
	case *sdl.TextInputEvent:
		s.RuneComplete(t)
	}
	return nil
}
func (s *Handler) Key(event *sdl.KeyboardEvent) {
	var key = gumi.KEY_UNKNOWN
	switch event.Keysym.Scancode {
	case sdl.SCANCODE_UNKNOWN:
		key = gumi.KEY_UNKNOWN
	case sdl.SCANCODE_A:
		key = gumi.KEY_A
	case sdl.SCANCODE_B:
		key = gumi.KEY_B
	case sdl.SCANCODE_C:
		key = gumi.KEY_C
	case sdl.SCANCODE_D:
		key = gumi.KEY_D
	case sdl.SCANCODE_E:
		key = gumi.KEY_E
	case sdl.SCANCODE_F:
		key = gumi.KEY_F
	case sdl.SCANCODE_G:
		key = gumi.KEY_G
	case sdl.SCANCODE_H:
		key = gumi.KEY_H
	case sdl.SCANCODE_I:
		key = gumi.KEY_I
	case sdl.SCANCODE_J:
		key = gumi.KEY_J
	case sdl.SCANCODE_K:
		key = gumi.KEY_K
	case sdl.SCANCODE_L:
		key = gumi.KEY_L
	case sdl.SCANCODE_M:
		key = gumi.KEY_M
	case sdl.SCANCODE_N:
		key = gumi.KEY_N
	case sdl.SCANCODE_O:
		key = gumi.KEY_O
	case sdl.SCANCODE_P:
		key = gumi.KEY_P
	case sdl.SCANCODE_Q:
		key = gumi.KEY_Q
	case sdl.SCANCODE_R:
		key = gumi.KEY_R
	case sdl.SCANCODE_S:
		key = gumi.KEY_S
	case sdl.SCANCODE_T:
		key = gumi.KEY_T
	case sdl.SCANCODE_U:
		key = gumi.KEY_U
	case sdl.SCANCODE_V:
		key = gumi.KEY_V
	case sdl.SCANCODE_W:
		key = gumi.KEY_W
	case sdl.SCANCODE_X:
		key = gumi.KEY_X
	case sdl.SCANCODE_Y:
		key = gumi.KEY_Y
	case sdl.SCANCODE_Z:
		key = gumi.KEY_Z
	case sdl.SCANCODE_1:
		key = gumi.KEY_1
	case sdl.SCANCODE_2:
		key = gumi.KEY_2
	case sdl.SCANCODE_3:
		key = gumi.KEY_3
	case sdl.SCANCODE_4:
		key = gumi.KEY_4
	case sdl.SCANCODE_5:
		key = gumi.KEY_5
	case sdl.SCANCODE_6:
		key = gumi.KEY_6
	case sdl.SCANCODE_7:
		key = gumi.KEY_7
	case sdl.SCANCODE_8:
		key = gumi.KEY_8
	case sdl.SCANCODE_9:
		key = gumi.KEY_9
	case sdl.SCANCODE_0:
		key = gumi.KEY_0
	case sdl.SCANCODE_RETURN:
		key = gumi.KEY_ENTER
	case sdl.SCANCODE_ESCAPE:
		key = gumi.KEY_ESCAPE
	case sdl.SCANCODE_BACKSPACE:
		key = gumi.KEY_BACKSPACE
	case sdl.SCANCODE_TAB:
		key = gumi.KEY_TAB
	case sdl.SCANCODE_SPACE:
		key = gumi.KEY_SPACE
	case sdl.SCANCODE_MINUS:
		key = gumi.KEY_MINUS
	case sdl.SCANCODE_EQUALS:
		key = gumi.KEY_EQUAL
	case sdl.SCANCODE_LEFTBRACKET:
		key = gumi.KEY_LEFTBRACKET
	case sdl.SCANCODE_RIGHTBRACKET:
		key = gumi.KEY_RIGHTBRACKET
	case sdl.SCANCODE_BACKSLASH:
		key = gumi.KEY_BACKSLASH
	case sdl.SCANCODE_NONUSHASH:
		//
	case sdl.SCANCODE_SEMICOLON:
		key = gumi.KEY_SEMICOLON
	case sdl.SCANCODE_APOSTROPHE:
		key = gumi.KEY_APOSTROPHE
	//
	case sdl.SCANCODE_GRAVE:
		//key = gumi.KEY_GR
	//	key = gumi.KEY_
	case sdl.SCANCODE_COMMA:
		key = gumi.KEY_COMMA
	//
	case sdl.SCANCODE_PERIOD:
		key = gumi.KEY_PERIOD
	case sdl.SCANCODE_SLASH:
		key = gumi.KEY_SLASH
	case sdl.SCANCODE_CAPSLOCK:
		//key = gumi.KEY_
	case sdl.SCANCODE_F1:
		key = gumi.KEY_F1
	case sdl.SCANCODE_F2:
		key = gumi.KEY_F2
	case sdl.SCANCODE_F3:
		key = gumi.KEY_F3
	case sdl.SCANCODE_F4:
		key = gumi.KEY_F4
	case sdl.SCANCODE_F5:
		key = gumi.KEY_F5
	case sdl.SCANCODE_F6:
		key = gumi.KEY_F6
	case sdl.SCANCODE_F7:
		key = gumi.KEY_F7
	case sdl.SCANCODE_F8:
		key = gumi.KEY_F8
	case sdl.SCANCODE_F9:
		key = gumi.KEY_F9
	case sdl.SCANCODE_F10:
		key = gumi.KEY_F10
	case sdl.SCANCODE_F11:
		key = gumi.KEY_F11
	case sdl.SCANCODE_F12:
		key = gumi.KEY_F12
	case sdl.SCANCODE_PRINTSCREEN:
		//key = gumi.KEY_
	case sdl.SCANCODE_SCROLLLOCK:
		//key = gumi.KEY_
	case sdl.SCANCODE_PAUSE:
		//key = gumi.KEY_P
	case sdl.SCANCODE_INSERT:
		//key = gumi.KEY_
	case sdl.SCANCODE_HOME:
		//key = gumi.KEY_
	case sdl.SCANCODE_PAGEUP:
		//key = gumi.KEY_
	case sdl.SCANCODE_DELETE:
		key = gumi.KEY_DELETE
	case sdl.SCANCODE_END:
		//key = gumi.KEY_
	case sdl.SCANCODE_PAGEDOWN:
		//key = gumi.KEY_
	case sdl.SCANCODE_RIGHT:
		key = gumi.KEY_RIGHT
	case sdl.SCANCODE_LEFT:
		key = gumi.KEY_LEFT
	case sdl.SCANCODE_DOWN:
		key = gumi.KEY_DOWN
	case sdl.SCANCODE_UP:
		key = gumi.KEY_UP
	case sdl.SCANCODE_NUMLOCKCLEAR:
		//key = gumi.KEY_
	case sdl.SCANCODE_KP_DIVIDE:
		key = gumi.KEY_SLASH
	case sdl.SCANCODE_KP_MULTIPLY:
		//key = gumi.KEY_
	case sdl.SCANCODE_KP_MINUS:
		key = gumi.KEY_MINUS
	case sdl.SCANCODE_KP_PLUS:
		//key = gumi.KEY_PL
	case sdl.SCANCODE_KP_ENTER:
		key = gumi.KEY_ENTER
	case sdl.SCANCODE_KP_1:
		key = gumi.KEY_1
	case sdl.SCANCODE_KP_2:
		key = gumi.KEY_2
	case sdl.SCANCODE_KP_3:
		key = gumi.KEY_3
	case sdl.SCANCODE_KP_4:
		key = gumi.KEY_4
	case sdl.SCANCODE_KP_5:
		key = gumi.KEY_5
	case sdl.SCANCODE_KP_6:
		key = gumi.KEY_6
	case sdl.SCANCODE_KP_7:
		key = gumi.KEY_7
	case sdl.SCANCODE_KP_8:
		key = gumi.KEY_8
	case sdl.SCANCODE_KP_9:
		key = gumi.KEY_9
	case sdl.SCANCODE_KP_0:
		key = gumi.KEY_0
	case sdl.SCANCODE_KP_PERIOD:
		key = gumi.KEY_PERIOD
	case sdl.SCANCODE_NONUSBACKSLASH:
		//key = gumi.KEY_
	case sdl.SCANCODE_APPLICATION:
		//key = gumi.KEY_
	case sdl.SCANCODE_POWER:
		//key = gumi.KEY_
	case sdl.SCANCODE_KP_EQUALS:
		key = gumi.KEY_EQUAL
	case sdl.SCANCODE_F13:
	case sdl.SCANCODE_F14:
	case sdl.SCANCODE_F15:
	case sdl.SCANCODE_F16:
	case sdl.SCANCODE_F17:
	case sdl.SCANCODE_F18:
	case sdl.SCANCODE_F19:
	case sdl.SCANCODE_F20:
	case sdl.SCANCODE_F21:
	case sdl.SCANCODE_F22:
	case sdl.SCANCODE_F23:
	case sdl.SCANCODE_F24:
	case sdl.SCANCODE_EXECUTE:
	case sdl.SCANCODE_HELP:
	case sdl.SCANCODE_MENU:
	case sdl.SCANCODE_SELECT:
	case sdl.SCANCODE_STOP:
	case sdl.SCANCODE_AGAIN:
	case sdl.SCANCODE_UNDO:
	case sdl.SCANCODE_CUT:
	case sdl.SCANCODE_COPY:
	case sdl.SCANCODE_PASTE:
	case sdl.SCANCODE_FIND:
	case sdl.SCANCODE_MUTE:
	case sdl.SCANCODE_VOLUMEUP:
	case sdl.SCANCODE_VOLUMEDOWN:
	case sdl.SCANCODE_KP_COMMA:
		key = gumi.KEY_COMMA
	case sdl.SCANCODE_KP_EQUALSAS400:
	case sdl.SCANCODE_INTERNATIONAL1:
	case sdl.SCANCODE_INTERNATIONAL2:
	case sdl.SCANCODE_INTERNATIONAL3:
	case sdl.SCANCODE_INTERNATIONAL4:
	case sdl.SCANCODE_INTERNATIONAL5:
	case sdl.SCANCODE_INTERNATIONAL6:
	case sdl.SCANCODE_INTERNATIONAL7:
	case sdl.SCANCODE_INTERNATIONAL8:
	case sdl.SCANCODE_INTERNATIONAL9:
	case sdl.SCANCODE_LANG1:
	case sdl.SCANCODE_LANG2:
	case sdl.SCANCODE_LANG3:
	case sdl.SCANCODE_LANG4:
	case sdl.SCANCODE_LANG5:
	case sdl.SCANCODE_LANG6:
	case sdl.SCANCODE_LANG7:
	case sdl.SCANCODE_LANG8:
	case sdl.SCANCODE_LANG9:
	case sdl.SCANCODE_ALTERASE:
	case sdl.SCANCODE_SYSREQ:
	case sdl.SCANCODE_CANCEL:
	case sdl.SCANCODE_CLEAR:
	case sdl.SCANCODE_PRIOR:
	case sdl.SCANCODE_RETURN2:
	case sdl.SCANCODE_SEPARATOR:
	case sdl.SCANCODE_OUT:
	case sdl.SCANCODE_OPER:
	case sdl.SCANCODE_CLEARAGAIN:
	case sdl.SCANCODE_CRSEL:
	case sdl.SCANCODE_EXSEL:
	case sdl.SCANCODE_KP_00:
	case sdl.SCANCODE_KP_000:
	case sdl.SCANCODE_THOUSANDSSEPARATOR:
	case sdl.SCANCODE_DECIMALSEPARATOR:
	case sdl.SCANCODE_CURRENCYUNIT:
	case sdl.SCANCODE_CURRENCYSUBUNIT:
	case sdl.SCANCODE_KP_LEFTPAREN:
	case sdl.SCANCODE_KP_RIGHTPAREN:
	case sdl.SCANCODE_KP_LEFTBRACE:
	case sdl.SCANCODE_KP_RIGHTBRACE:
	case sdl.SCANCODE_KP_TAB:
		key = gumi.KEY_TAB
	case sdl.SCANCODE_KP_BACKSPACE:
		key = gumi.KEY_BACKSPACE
	case sdl.SCANCODE_KP_A:
		key = gumi.KEY_A
	case sdl.SCANCODE_KP_B:
		key = gumi.KEY_B
	case sdl.SCANCODE_KP_C:
		key = gumi.KEY_C
	case sdl.SCANCODE_KP_D:
		key = gumi.KEY_D
	case sdl.SCANCODE_KP_E:
		key = gumi.KEY_E
	case sdl.SCANCODE_KP_F:
		key = gumi.KEY_F
	case sdl.SCANCODE_KP_XOR:
	case sdl.SCANCODE_KP_POWER:
	case sdl.SCANCODE_KP_PERCENT:
	case sdl.SCANCODE_KP_LESS:
	case sdl.SCANCODE_KP_GREATER:
	case sdl.SCANCODE_KP_AMPERSAND:
	case sdl.SCANCODE_KP_DBLAMPERSAND:
	case sdl.SCANCODE_KP_VERTICALBAR:
	case sdl.SCANCODE_KP_DBLVERTICALBAR:
	case sdl.SCANCODE_KP_COLON:
	case sdl.SCANCODE_KP_HASH:
	case sdl.SCANCODE_KP_SPACE:
		key = gumi.KEY_SPACE
	case sdl.SCANCODE_KP_AT:
	case sdl.SCANCODE_KP_EXCLAM:
	case sdl.SCANCODE_KP_MEMSTORE:
	case sdl.SCANCODE_KP_MEMRECALL:
	case sdl.SCANCODE_KP_MEMCLEAR:
	case sdl.SCANCODE_KP_MEMADD:
	case sdl.SCANCODE_KP_MEMSUBTRACT:
	case sdl.SCANCODE_KP_MEMMULTIPLY:
	case sdl.SCANCODE_KP_MEMDIVIDE:
	case sdl.SCANCODE_KP_PLUSMINUS:
	case sdl.SCANCODE_KP_CLEAR:
	case sdl.SCANCODE_KP_CLEARENTRY:
	case sdl.SCANCODE_KP_BINARY:
	case sdl.SCANCODE_KP_OCTAL:
	case sdl.SCANCODE_KP_DECIMAL:
	case sdl.SCANCODE_KP_HEXADECIMAL:
	case sdl.SCANCODE_LCTRL:
		key = gumi.KEY_CONTROL
	case sdl.SCANCODE_LSHIFT:
		key = gumi.KEY_SHIFT
	case sdl.SCANCODE_LALT:
		key = gumi.KEY_ALT
	case sdl.SCANCODE_LGUI:
	case sdl.SCANCODE_RCTRL:
		key = gumi.KEY_CONTROL
	case sdl.SCANCODE_RSHIFT:
		key = gumi.KEY_SHIFT
	case sdl.SCANCODE_RALT:
		key = gumi.KEY_ALT
	case sdl.SCANCODE_RGUI:
	case sdl.SCANCODE_MODE:
	case sdl.SCANCODE_AUDIONEXT:
	case sdl.SCANCODE_AUDIOPREV:
	case sdl.SCANCODE_AUDIOSTOP:
	case sdl.SCANCODE_AUDIOPLAY:
	case sdl.SCANCODE_AUDIOMUTE:
	case sdl.SCANCODE_MEDIASELECT:
	case sdl.SCANCODE_WWW:
	case sdl.SCANCODE_MAIL:
	case sdl.SCANCODE_CALCULATOR:
	case sdl.SCANCODE_COMPUTER:
	case sdl.SCANCODE_AC_SEARCH:
	case sdl.SCANCODE_AC_HOME:
	case sdl.SCANCODE_AC_BACK:
	case sdl.SCANCODE_AC_FORWARD:
	case sdl.SCANCODE_AC_STOP:
	case sdl.SCANCODE_AC_REFRESH:
	case sdl.SCANCODE_AC_BOOKMARKS:
	case sdl.SCANCODE_BRIGHTNESSDOWN:
	case sdl.SCANCODE_BRIGHTNESSUP:
	case sdl.SCANCODE_DISPLAYSWITCH:
	case sdl.SCANCODE_KBDILLUMTOGGLE:
	case sdl.SCANCODE_KBDILLUMDOWN:
	case sdl.SCANCODE_KBDILLUMUP:
	case sdl.SCANCODE_EJECT:
	case sdl.SCANCODE_SLEEP:
	case sdl.SCANCODE_APP1:
	case sdl.SCANCODE_APP2:
	}

	switch event.Type {
	default:
	case sdl.KEYDOWN:
		if v, ok := s.keymap[key]; ok && v == gumi.EVENT_KEYPRESS {
			return
		}
		s.keymap[key] = gumi.EVENT_KEYPRESS
		s.glumi.screen.Event(gumi.EventKeyPress{
			Key: key,
		})
	case sdl.KEYUP:
		if v, ok := s.keymap[key]; ok && v == gumi.EVENT_KEYRELEASE {
			return
		}
		s.keymap[key] = gumi.EVENT_KEYRELEASE
		s.glumi.screen.Event(gumi.EventKeyRelease{
			Key: key,
		})

	}
}
func (s *Handler) Cursor(event *sdl.MouseMotionEvent) {
	s.glumi.screen.Event(gumi.EventCursor{
		Point: image.Point{
			X: int(event.X), Y: int(event.Y),
		},
	})
}
func (s *Handler) Mouse(event *sdl.MouseButtonEvent) {
	var key gumi.GUMIKey
	switch event.Button {
	case sdl.BUTTON_LEFT:
		key = gumi.KEY_MOUSE1
	case sdl.BUTTON_MIDDLE:
		key = gumi.KEY_MOUSE3
	case sdl.BUTTON_RIGHT:
		key = gumi.KEY_MOUSE2
	//case sdl.BUTTON_X1:
	//case sdl.BUTTON_X2:
	default:
		key = gumi.KEY_UNKNOWN
	}
	switch event.Type {
	default:
	case sdl.MOUSEBUTTONDOWN:
		s.glumi.screen.Event(gumi.EventKeyPress{
			Key: key,
		})
	case sdl.MOUSEBUTTONUP:
		s.glumi.screen.Event(gumi.EventKeyRelease{
			Key: key,
		})
	}
}
func (s *Handler) Scrool(event *sdl.MouseWheelEvent) {
	s.glumi.screen.Event(gumi.EventScroll{
		Point: image.Point{
			X: int(event.X),
			Y: -int(event.Y),
		},
	})
}
func (s *Handler) RuneEdit(event *sdl.TextEditingEvent) {
	r, size := utf8.DecodeRune(event.Text[:])
	if size <= 0 {
		return
	}
	s.glumi.screen.Event(gumi.EventRuneEdit{
		Rune: r,
	})
}
func (s *Handler) RuneComplete(event *sdl.TextInputEvent) {
	r, size := utf8.DecodeRune(event.Text[:])
	if size <= 0 || r == 0 {
		return
	}
	s.glumi.screen.Event(gumi.EventRuneComplete{
		Rune: r,
	})
}
