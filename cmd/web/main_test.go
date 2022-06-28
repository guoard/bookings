package main

import "testing"

func TestR(t *testing.T) {
	_, err := run()
	if err != nil {
		t.Error("failed run()")
	}
}