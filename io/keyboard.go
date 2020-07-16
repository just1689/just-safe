package io

import (
	"github.com/micmonay/keybd_event"
	"github.com/sirupsen/logrus"
	"time"
)

type BasicFunc func()

func GenerateKeyboardScript(in string) (f func(), err error) {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		logrus.Errorln(err)
		return
	}

	keys := make([]BasicFunc, 0)
	for i := 0; i < len(in); i++ {
		keys = append(keys, getNext(string(in[i]), kb))
	}

	f = func() {
		for _, k := range keys {
			k()
		}
	}
	return
}

func getNext(char string, kb keybd_event.KeyBonding) func() {
	return func() {
		kb.Clear()
		if char >= "A" && char <= "Z" {
			kb.HasSHIFT(true)
		}
		key, found := keyMapping[char]
		if found {
			kb.AddKey(key)
		} else {
			logrus.Errorln("could not find key", char)
		}
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
	}
}

var keyMapping = map[string]int{
	"A": keybd_event.VK_A,
	"a": keybd_event.VK_A,
	"B": keybd_event.VK_B,
	"b": keybd_event.VK_B,
	"C": keybd_event.VK_C,
	"c": keybd_event.VK_C,
	"d": keybd_event.VK_D,
	"D": keybd_event.VK_D,
	"e": keybd_event.VK_E,
	"E": keybd_event.VK_E,
	"f": keybd_event.VK_F,
	"F": keybd_event.VK_F,
	"G": keybd_event.VK_G,
	"g": keybd_event.VK_G,
	"H": keybd_event.VK_H,
	"h": keybd_event.VK_H,
	"I": keybd_event.VK_I,
	"i": keybd_event.VK_I,
	"J": keybd_event.VK_J,
	"j": keybd_event.VK_J,
	"K": keybd_event.VK_K,
	"k": keybd_event.VK_K,
	"L": keybd_event.VK_L,
	"l": keybd_event.VK_L,
	"M": keybd_event.VK_M,
	"m": keybd_event.VK_M,
	"N": keybd_event.VK_N,
	"n": keybd_event.VK_N,
	"O": keybd_event.VK_O,
	"o": keybd_event.VK_O,
	"P": keybd_event.VK_P,
	"p": keybd_event.VK_P,
	"Q": keybd_event.VK_Q,
	"q": keybd_event.VK_Q,
	"R": keybd_event.VK_R,
	"r": keybd_event.VK_R,
	"S": keybd_event.VK_S,
	"s": keybd_event.VK_S,
	"T": keybd_event.VK_T,
	"t": keybd_event.VK_T,
	"U": keybd_event.VK_U,
	"u": keybd_event.VK_U,
	"V": keybd_event.VK_V,
	"v": keybd_event.VK_V,
	"W": keybd_event.VK_W,
	"w": keybd_event.VK_W,
	"X": keybd_event.VK_X,
	"x": keybd_event.VK_X,
	"Y": keybd_event.VK_Y,
	"y": keybd_event.VK_Y,
	"Z": keybd_event.VK_Z,
	"z": keybd_event.VK_Z,
	" ": keybd_event.VK_SPACE,
	"":  keybd_event.VK_SPACE,
	// Need more than one
	"-":  keybd_event.VK_MINUS,
	"=":  keybd_event.VK_EQUAL,
	"[":  keybd_event.VK_LEFTBRACE,
	"]":  keybd_event.VK_RIGHTBRACE,
	";":  keybd_event.VK_SEMICOLON,
	"'":  keybd_event.VK_APOSTROPHE,
	"\\": keybd_event.VK_BACKSLASH,
	",":  keybd_event.VK_COMMA,
	".":  keybd_event.VK_DOT,
	"/":  keybd_event.VK_SLASH,
	"`":  keybd_event.VK_KPASTERISK,
	// Other half
	"_":  keybd_event.VK_MINUS,
	"+":  keybd_event.VK_EQUAL,
	"{":  keybd_event.VK_LEFTBRACE,
	"}":  keybd_event.VK_RIGHTBRACE,
	":":  keybd_event.VK_SEMICOLON,
	"\"": keybd_event.VK_APOSTROPHE,
	"|":  keybd_event.VK_BACKSLASH,
	"<":  keybd_event.VK_COMMA,
	">":  keybd_event.VK_DOT,
	"?":  keybd_event.VK_SLASH,
	"~":  keybd_event.VK_KPASTERISK,
}
