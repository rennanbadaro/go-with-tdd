package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, expected string, got string) {
		t.Helper()

		if got != expected {
			t.Errorf("Expected %q but got %q", expected, got)
		}
	}

	t.Run("Should say hello to someone", func(t *testing.T) {
		got := Hello("Rennan", "")
		expected := "Hi there, Rennan"

		assertMessage(t, expected, got)
	})

	t.Run("Should say 'Hi there, stranger' empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		expected := "Hi there, stranger"

		assertMessage(t, expected, got)
	})

	t.Run("Should greet in Spanish", func(t *testing.T) {
		got := Hello("Rennan", "Spanish")
		expected := "Hola, Rennan"

		assertMessage(t, expected, got)
	})

	t.Run("Should greet in French", func(t *testing.T) {
		got := Hello("Rennan", "French")
		expected := "Bonjour, Rennan"

		assertMessage(t, expected, got)
	})
}
