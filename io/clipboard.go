package io

import (
	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"
)

func WriteToClipboard(in string) {
	if err := clipboard.WriteAll(in); err != nil {
		logrus.Errorln(err)
		return
	}
	logrus.Infoln("> Copied to clipboard")
}
