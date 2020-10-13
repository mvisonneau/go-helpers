package logger

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// Config : Type that handles logging config
type Config struct {
	Level        string
	Format       string
	ReportCaller bool
}

// Configure the logger
func Configure(c Config) (err error) {
	parsedLevel, err := log.ParseLevel(c.Level)
	if err != nil {
		return
	}
	log.SetLevel(parsedLevel)

	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	switch c.Format {
	case "text":
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		err = fmt.Errorf("Invalid log format '%s'", c.Format)
		return
	}

	log.SetReportCaller(c.ReportCaller)
	log.SetOutput(os.Stdout)

	return
}
