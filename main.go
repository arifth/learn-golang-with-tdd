package main

import (
	"fmt"
	"os"
)

func main() {
	for _, value := range os.Args[1:] {
		if value != "" {
			fmt.Println(Iterate(value))

		}
		fmt.Println(Iterate(""))
	}
}
