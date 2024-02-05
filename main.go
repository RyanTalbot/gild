package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

		showLineCount(strings.Split(string(data), "\n"))

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

func showLineCount(lines []string) {
	count := 1
	for line := range lines {
		fmt.Printf("|%4s:\t%4s\n", strconv.Itoa(count), lines[line])
		count++
	}
}
