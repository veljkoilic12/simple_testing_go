package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Mark", "French")
		want := "Bonjour, Mark"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Serbian", func(t *testing.T) {
		got := Hello("Veljko", "Serbian")
		want := "Pozdrav, Veljko"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// ova linija koda govori da je ova fja helper i da se prilikom error-a, prikaze linija koda koja je
	// pozvala ovu fju, umesto linije koda na kojoj se greska pojavila
	t.Helper()
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}
