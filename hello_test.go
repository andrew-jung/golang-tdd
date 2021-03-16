package main

import "testing"

/*
File must be xyz_test.go
Test function must start with Test*
Takes only one argument, from "testing"
*/
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Saying hello to specific person", func(t *testing.T) {
		got := Hello("Andrew", "")
		want := "Hello Andrew"

		assertCorrectMessage(t, got, want)

	})
	t.Run("Saying hello world, when given empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello in Spanish", func(t *testing.T) {
		got := Hello("Bob", "Spanish")
		want := "Hola Bob"

		assertCorrectMessage(t, got, want)
	})
}
