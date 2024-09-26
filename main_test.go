package main

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("chris")
		want := "hallo arif"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("say: 'Hello World', when empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "hello World"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
