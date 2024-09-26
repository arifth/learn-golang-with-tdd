package main

func Iterate(character string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated = repeated + character + "\n"
	}

	return repeated
}
