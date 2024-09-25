package main

import (
	"testing"
)

func Test(t *testing.T) {
	// fmt.Println("apa ini ? ", t)
	got := Hello("chris")
	want := "hallo arif"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
