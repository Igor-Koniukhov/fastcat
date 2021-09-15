package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	err := SetAndRun()
	if err != nil {
		t.Error("failed run")
	}
}

