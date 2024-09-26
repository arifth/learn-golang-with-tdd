package main

import (
	"testing"
)

func Test_Hello_func(t *testing.T) {
	t.Run("saying hello with japan params", func(t *testing.T) {
		got := Hello("japan")
		want := "Hello from japan"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("saying hello with french params", func(t *testing.T) {
		got := Hello("french")
		want := "Hello from french"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
