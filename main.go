package main

import (
	"fmt"
)

func Hello(s string) string {
	char := "hello " + s
	return char
}

func main() {
	fmt.Println(Hello("arif"))
}
