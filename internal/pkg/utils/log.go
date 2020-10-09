package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

// init is configuring the project logrus.Logger
//
// Configuration changes if we have set GOPYPI_ENV="PRODUCTION"
// which means we have deployed gopypi for production use.
func init() {
	f := &logrus.TextFormatter{
		DisableColors:   false,
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	if runtime.GOOS == "windows" {
		f.ForceColors = true
	}

	if os.Getenv("GOPYPI_ENV") == "PRODUCTION" {
		// Log as JSON instead of the default ASCII formatter.
		logrus.SetFormatter(&logrus.JSONFormatter{})
		// Disable colors when using JSON
		f.DisableColors = true
		// Only log the warning severity or above.
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetFormatter(f)
}
