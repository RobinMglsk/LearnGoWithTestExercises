package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("It should say hello to people", func(t *testing.T) {
		got := Hello("Robin", "")
		want := "Hello, Robin"

		assertCorrectMessage(t, got, want)
	})

	t.Run("It should say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("It should support Spanish", func(t *testing.T) {
		got := Hello("Jolien", "es")
		want := "Hola, Jolien"

		assertCorrectMessage(t, got, want)
	})

	t.Run("It should support French", func(t *testing.T) {
		got := Hello("Ebe", "fr")
		want := "Bonjour, Ebe"

		assertCorrectMessage(t, got, want)
	})

	t.Run("It should support Dutch", func(t *testing.T) {
		got := Hello("Robin", "nl")
		want := "Hallo, Robin"

		assertCorrectMessage(t, got, want)
	})
}
