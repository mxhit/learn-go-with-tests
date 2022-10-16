package main

import "testing"

func TestHello(t *testing.T) {
	var got, want string

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("Saying hello to people in French", func(t *testing.T) {
		got = Hello("Doodle", "French")
		want = "Bonjour, Doodle!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Empty parameter", func(t *testing.T) {
		got = Hello("", "Spanish")
		want = "Hola, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Empty language", func(t *testing.T) {
		got = Hello("Modo", "")
		want = "Hello, Modo!"

		assertCorrectMessage(t, got, want)
	})
}
