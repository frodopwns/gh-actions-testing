package main

import (
	"testing"
)

func TestSay(t *testing.T) {
	s := say("Hello, World!!")
	if s != "Hello, World!" {
		t.Error("Expected 'Hello, World!', got ", s)
	}
}
