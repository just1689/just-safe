package io

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"runtime"
	"strings"
	"syscall"
)

type NextInput struct {
	Prompt       string
	Secure       bool
	Key          string
	NewLineAfter bool
}

var PromptSite = NextInput{
	Prompt:       "Enter site: ",
	Secure:       false,
	Key:          "site",
	NewLineAfter: true,
}

var PromptUsername = NextInput{
	Prompt:       "Enter username: ",
	Secure:       false,
	Key:          "username",
	NewLineAfter: true,
}

var PromptPassword = NextInput{
	Prompt:       "Enter password: ",
	Secure:       true,
	Key:          "password",
	NewLineAfter: true,
}

var PromptWalletPassword = NextInput{
	Prompt:       "Enter wallet password: ",
	Secure:       true,
	Key:          "walletPassword",
	NewLineAfter: true,
}

func PromptForInput(input []NextInput) (result map[string]string, err error) {
	result = make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	for _, i := range input {
		var bytePassword []byte

		if i.Secure {
			fmt.Print(i.Prompt)
			if runtime.GOOS == "windows" {
				bytePassword, err = terminal.ReadPassword(int(syscall.Stdin)) // 0 on *nix
			} else {
				bytePassword, err = terminal.ReadPassword(0)
			}
			if err != nil {
				//Handle error
				logrus.Errorln(err)
				continue
			}
			result[i.Key] = strings.TrimSpace(string(bytePassword))
			if i.NewLineAfter {
				fmt.Println("")
			}
		} else {
			fmt.Print(i.Prompt)
			next, _ := reader.ReadString('\n')
			result[i.Key] = strings.TrimSpace(next)
		}
	}
	return
}
