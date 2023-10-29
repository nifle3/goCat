package main

import (
	"fmt"
	"goCustomCat/internal/writerReader"
	"os"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) <= 2 {
		fmt.Printf("Hi, this program is designed for read text files, " +
			"please enter a path of file, you want to read.\n " +
			"You can enter the '--notAsync' flag, then the program will run sequentially\n")

		os.Exit(0)
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Please enter a correct path, your path is %s\n", filePath)

		os.Exit(0)
	}
	defer file.Close()

	if len(os.Args) == 3 && os.Args[2] == "--notAsync" {
		writerReader.NotAsync(file)
	} else {
		writerReader.Async(file)
	}

	duration := time.Since(start)
	fmt.Println(duration)
}
