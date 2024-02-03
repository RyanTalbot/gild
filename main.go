package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	args_count := len(args)

	if args_count <= 1 {
		readStandardIn()
	}

	if args_count > 1 {
		if args[1] == "-" {
			readStandardIn()
		} else {
			readFileFromUser(args[1:])
		}
	}
}

func readFileFromUser(paths []string) {
	for path := range paths {
		file, err := os.Open(paths[path])
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err := os.ReadFile(paths[path])
		if err != nil {
			return
		}
		fmt.Print(string(data))

		file.Close()
	}
}

func readStandardIn() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	str := string(input)

	fmt.Print(str)
}
