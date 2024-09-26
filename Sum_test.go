package main

import "testing"

func Test_sum(t *testing.T) {
	got := Add(1, 2)
	want := 5

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
