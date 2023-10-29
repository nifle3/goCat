package benchmarks

import (
	"fmt"
	"goCustomCat/internal/writerReader"
	"os"
	"testing"
)

var fileNames = [...]string{"test.txt", "test2.txt"}

func BenchmarkAsync(b *testing.B) {
	for _, fileName := range fileNames {
		b.Run(fmt.Sprintf("%s async", fileName), func(b *testing.B) {
			file, _ := os.Open(fileName)
			writerReader.Async(file)

			file.Close()
		})
	}
}

func BenchmarkNotAsync(b *testing.B) {
	for _, fileName := range fileNames {
		b.Run(fmt.Sprintf("%s not async", fileName), func(b *testing.B) {
			file, _ := os.Open(fileName)
			writerReader.NotAsync(file)

			file.Close()
		})
	}
}
