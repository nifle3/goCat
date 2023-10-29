package writerReader

import (
	"fmt"
	"io"
	"os"
)

func NotAsync(file *os.File) {
	result := make([]byte, 32)

	for {
		n, err := file.Read(result)
		if err == io.EOF {
			break
		}

		fmt.Print(string(result[:n]))
	}

}
