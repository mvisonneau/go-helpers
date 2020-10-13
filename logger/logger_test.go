package logger

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLoggerConfigureFatalText(t *testing.T) {
	err := Configure(Config{
		Level:  "fatal",
		Format: "text",
	})

	assert.NoError(t, err)
	assert.Equal(t, log.FatalLevel, log.GetLevel())
}

func TestLoggerConfigureInvalidFormat(t *testing.T) {
	err := Configure(Config{
		Level:  "fatal",
		Format: "foo",
	})

	assert.Error(t, err)
}

func TestLoggerConfigureFormatJson(t *testing.T) {
	err := Configure(Config{
		Level:  "debug",
		Format: "json",
	})

	assert.NoError(t, err)
}

func TestLoggerConfigureInvalidLevel(t *testing.T) {
	err := Configure(Config{
		Level:  "foo",
		Format: "text",
	})
	assert.Error(t, err)
}
