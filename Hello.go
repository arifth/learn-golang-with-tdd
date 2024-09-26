package main

func Hello(s string) (prefix string) {
	if s == "" {
		prefix = "hello World"
	}
	switch s {
	case "japan":
		prefix = "Hello from japan"
		return
	case "french":
		prefix = "Hello from french"
		return
	}
	prefix = "Hello from anywhere "
	return
}
