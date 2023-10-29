package writerReader

import (
	"bufio"
	"fmt"
	"os"
)

func Async(file *os.File) {
	dataChan := make(chan string)
	go printer(dataChan)

	br := bufio.NewReader(file)

	for {
		b, err := br.ReadByte()
		if err != nil {
			break
		}

		dataChan <- string(b)
	}
	close(dataChan)
}

func printer(dataChan <-chan string) {
	for data := range dataChan {
		fmt.Print(data)
	}
}
