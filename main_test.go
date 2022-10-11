package main

import (
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	testLog := Log{}
	testLog.SetDefaultValues()
	testLog.Context = []Metadata{
		{
			Name:  "foo",
			Value: "bar",
		},
		{
			Name:  "lur",
			Value: "pelle",
		},
	}
	testLog.Info("bosse")
	testLog.ErrorM("frasse", []Metadata{{Name: "wat", Value: ":O"}})
}

func TestWoo(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	log := Log{
		Context:      []Metadata{},  // Will be prepended to metadata on all log entries
		MinLogLvl:    Info,          // Minimal log level to output
		Fmt:          DefaultFmt,    // Log message formatter
		Stderr:       DefaultStderr, // Log message outputter for Debug, Verbose and Info
		Stdout:       DefaultStdout, // Log message outputter for Warning and Error
		TimeLocation: loc,           // Timestamp location/time zone setting
	}
	log.Info("wat")
}
