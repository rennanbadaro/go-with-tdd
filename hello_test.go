package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Rennan")
	expected := "Hi there, Rennan"

	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}
