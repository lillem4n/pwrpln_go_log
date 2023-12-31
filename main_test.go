package go_log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	stdout := ""
	stderr := ""
	testLog := GetLog()
	testLog.Stdout = func(msg string) {
		stdout += msg
	}
	testLog.Stderr = func(msg string) {
		stderr += msg
	}
	testLog.Warn("oh noes")
	testLog.Info("einfoo")
	testLog.Verbose("moar moar")
	testLog.Debug("deboug")

	// Only info should be processed due to default log limit being info
	assert.Equal(t, "[\x1b[32mInf\x1b[0m] einfoo\n", stdout[20:])
	assert.Equal(t, "[\x1b[33mWrn\x1b[0m] oh noes\n", stderr[20:])
}

func TestError(t *testing.T) {
	stderr := ""
	testLog := GetLog()
	testLog.Stderr = func(msg string) {
		stderr += msg
	}
	testLog.Error("lureri")
	assert.Equal(t, "[\x1b[31mErr\x1b[0m] lureri\n", stderr[20:])
}

// func TestMetadata(t *testing.T) {
// 	testLog := GetLog()
// 	testLog.Context = []interface{}{
// 		"foo", "bar",
// 		"lur", "pelle",
// 	}
// 	testLog.Info("bosse")
// 	testLog.Error("frasse", "wat", ":O")
// }

// func TestWoo(t *testing.T) {
// 	loc, _ := time.LoadLocation("UTC")
// 	log := Log{
// 		Context:      []Metadata{},  // Will be prepended to metadata on all log entries
// 		MinLogLvl:    Info,          // Minimal log level to output
// 		Fmt:          DefaultFmt,    // Log message formatter
// 		Stderr:       DefaultStderr, // Log message outputter for Debug, Verbose and Info
// 		Stdout:       DefaultStdout, // Log message outputter for Warning and Error
// 		TimeLocation: loc,           // Timestamp location/time zone setting
// 	}
// 	log.Info("wat")
// }
