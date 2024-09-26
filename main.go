package main

import (
	"fmt"
	"os"
)

func Hello(s string) (prefix string) {
	if s == "" {
		prefix = "hello World"
	}
	switch s {
	case "japan":
		fmt.Println("masuk sini")
		prefix = "Hello from japan"
		return
	case "french":
		prefix = "Hello form french"
		return
	}
	prefix = "Hello from anywhere "
	return
}

func main() {
	for _, value := range os.Args[1:] {
		if value != "" {
			fmt.Println(Hello(value))

		}
		fmt.Println(Hello(""))
	}
}
