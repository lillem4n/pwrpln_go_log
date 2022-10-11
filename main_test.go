package main

import "testing"

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
