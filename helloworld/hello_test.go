package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func (t *testing.T) {
	got := Hello("Issam", "")

	want := "Salam, Issam!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Salam, 3alam' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Salam, 3alam!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In tamazight", func( t * testing.T) {
		got := Hello("Tayurt", "tamazight" ) 
		want := "Azul, Tayurt!"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("In french", func( t *testing.T) {
		got := Hello("Adrien", "french")
		want := "Bonjour, Adrien!"
		assertCorrectMessage(t, got, want)
	})
}


func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
