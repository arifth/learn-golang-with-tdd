package main

import (
	"fmt"
	"testing"
)

func BenchmarkIteration(b *testing.B) {

	fmt.Println("apa isi b.N", b.N)
	for i := 0; i < b.N; i++ {
		Iterate("test")
	}
}
