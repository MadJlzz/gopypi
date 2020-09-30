package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = &logrus.Logger{
	Out:   os.Stderr,
	Level: logrus.DebugLevel,
	Formatter: &logrus.TextFormatter{
		DisableColors:   true,
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	},
}
