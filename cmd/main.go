package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Hi, this program is designed for read text files, " +
			"please enter a path of file, you want to read\n")

		os.Exit(0)
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Please enter a correct path, your path is %s\n", filePath)

		os.Exit(0)
	}
	defer file.Close()

	br := bufio.NewReader(file)

	for {
		b, err := br.ReadByte()
		if err != nil {
			break
		}

		fmt.Print(string(b))
	}
}
