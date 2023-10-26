package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Hi, this program is designed for read text files, " +
			"please enter a path of file, you want to read\n")

		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("Please enter a correct path, your path is %s\n", filePath)

		return
	}
	defer file.Close()
	
	result, err := io.ReadAll(file)
	if err != nil {
		panic("IDK")
	}

	fmt.Printf("%s", string(result))
}
