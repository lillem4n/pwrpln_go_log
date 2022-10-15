package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	stdout := ""
	stderr := ""
	testLog := Log{}
	testLog.SetDefaultValues()
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
	testLog := Log{}
	testLog.SetDefaultValues()
	testLog.Stderr = func(msg string) {
		stderr += msg
	}
	testLog.Error("lureri")
	assert.Equal(t, "[\x1b[31mErr\x1b[0m] lureri\n", stderr[20:])
}

// func TestMetadata(t *testing.T) {
// 	testLog := Log{}
// 	testLog.SetDefaultValues()
// 	testLog.Context = []Metadata{
// 		{
// 			Name:  "foo",
// 			Value: "bar",
// 		},
// 		{
// 			Name:  "lur",
// 			Value: "pelle",
// 		},
// 	}
// 	testLog.Info("bosse")
// 	testLog.ErrorM("frasse", []Metadata{{Name: "wat", Value: ":O"}})
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
