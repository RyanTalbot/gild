package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		readStandardIn("")
	}

	if len(args) > 1 {
		if args[1] == "-" {
			readStandardIn(args[1])
		} else {
			readFileFromUser(args[1])
		}
	}
}

func readFileFromUser(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	fmt.Println(string(data))

	file.Close()
}

func readStandardIn(source string) {
	fmt.Println("Let's pipe it")
}
