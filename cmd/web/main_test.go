package main

import "testing"

func TestR(t *testing.T) {
	err := run()
	if err != nil {
		t.Error("failed run()")
	}
}