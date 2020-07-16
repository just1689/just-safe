package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/micmonay/keybd_event"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func main() {
	endpoint := "http://localhost:8080/"

	if os.Args[1] == "add" {
		site, username, password, walletPassword := credentials()
		url := endpoint + "api/addPassword/v1"
		body := map[string]string{
			"site":           site,
			"username":       username,
			"walletPassword": walletPassword,
			"password":       password,
		}
		b, err := json.Marshal(body)
		if err != nil {
			//handle
			logrus.Errorln(err)
			return
		}
		//TODO: loading bar
		http.Post(url, "application/json", bytes.NewReader(b))

	} else if os.Args[1] == "get" {
		site, username, walletPassword := getter()
		url := endpoint + "api/getPassword/v1"
		body := map[string]string{
			"site":           site,
			"username":       username,
			"walletPassword": walletPassword,
		}
		b, err := json.Marshal(body)
		if err != nil {
			//handle
			logrus.Errorln(err)
			return
		}
		//TODO: loading bar
		resp, err := http.Post(url, "application/json", bytes.NewReader(b))
		if err != nil {
			logrus.Error(err)
			return
		}
		fmt.Println(resp.StatusCode)
		b, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
			return
		}

		r := make(map[string]string)
		err = json.Unmarshal(b, &r)
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Println("")
		fmt.Println("")

		pass := r["password"]
		if os.Args[2] == "clipboard" || os.Args[2] == "c" {
			if err := clipboard.WriteAll(pass); err != nil {
				panic(err)
			}
			fmt.Println("COPIED TO CLIPBOARD!")
		}

		if os.Args[2] == "keyboard" || os.Args[2] == "k" {
			kb, err := keybd_event.NewKeyBonding()
			if err != nil {
				panic(err)
			}
			keys := make([]BasicFunc, 0)
			for i := 0; i < len(pass); i++ {
				keys = append(keys, getNext(string(pass[i]), kb))
			}

			for i := 1; i <= 3; i++ {
				fmt.Println("Writing in", i)
				time.Sleep(500 * time.Millisecond)
			}
			for _, k := range keys {
				k()
			}
		}

	}

}

type BasicFunc func()

func credentials() (string, string, string, string) {
	reader := bufio.NewReader(os.Stdin)
	var bytePassword []byte
	var err error

	fmt.Print("Enter site: ")
	site, _ := reader.ReadString('\n')

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	if runtime.GOOS == "windows" {
		bytePassword, err = terminal.ReadPassword(int(syscall.Stdin)) // 0 on unix
	} else {
		bytePassword, err = terminal.ReadPassword(0)
	}
	if err != nil {
		//Handle error
		return "", "", "", ""
	}
	password := string(bytePassword)
	fmt.Println("")

	fmt.Print("Enter Wallet password: ")
	if runtime.GOOS == "windows" {
		bytePassword, err = terminal.ReadPassword(int(syscall.Stdin)) // 0 on unix
	} else {
		bytePassword, err = terminal.ReadPassword(0)
	}
	if err != nil {
		//Handle error
		return "", "", "", ""
	}
	walletPassword := string(bytePassword)

	return strings.TrimSpace(site), strings.TrimSpace(username), strings.TrimSpace(password), strings.TrimSpace(walletPassword)
}

func getter() (string, string, string) {
	reader := bufio.NewReader(os.Stdin)
	var bytePassword []byte
	var err error

	fmt.Print("Enter site: ")
	site, _ := reader.ReadString('\n')

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter Wallet password: ")
	if runtime.GOOS == "windows" {
		bytePassword, err = terminal.ReadPassword(int(syscall.Stdin)) // 0 on unix
	} else {
		bytePassword, err = terminal.ReadPassword(0)
	}
	if err != nil {
		//Handle error
		return "", "", ""
	}
	walletPassword := string(bytePassword)

	return strings.TrimSpace(site), strings.TrimSpace(username), strings.TrimSpace(walletPassword)
}

func pressKeys() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important to wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	// Select keys to be pressed
	kb.SetKeys(keybd_event.VK_A, keybd_event.VK_B)

	// Set shift to be pressed
	kb.HasSHIFT(true)

	// Press the selected keys
	err = kb.Launching()
	if err != nil {
		panic(err)
	}

	// Or you can use Press and Release
	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()

	// Here, the program will generate "ABAB" as if they were pressed on the keyboard.
}

func getNext(char string, kb keybd_event.KeyBonding) func() {
	return func() {
		kb.Clear()
		if char >= "A" && char <= "Z" {
			kb.HasSHIFT(true)
		}
		key, found := KeyMapping[char]
		if found {
			kb.AddKey(key)
		} else {
			fmt.Println("could not find key", char)
		}
		kb.Press()
		time.Sleep(10 * time.Millisecond)
		kb.Release()
	}
}

var KeyMapping = map[string]int{
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
