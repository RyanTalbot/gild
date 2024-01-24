package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	var source string

	if len(args) > 1 {
		source = args[1]
	}
	fmt.Println(source)

	file, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("File opened")

	data, err := os.ReadFile(source)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Contents:\n", string(data))

	file.Close()
}
