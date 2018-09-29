package buffer

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func BenchmarkUnbufferedFileWrite(b *testing.B) {
	file, err := os.Create("unbuffered.test")
	if err != nil {
		b.Fatalf("Unable to create file: %v", err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	for i := 0; i < b.N; i++ {
		fmt.Fprintln(file, "Hello world")
	}
}

func BenchmarkBufferedFileWrite(b *testing.B) {
	file, err := os.Create("buffered.test")
	if err != nil {
		b.Fatalf("Unable to create file: %v", err)
	}
	defer func() {
		file.Close()
		os.Remove(file.Name())
	}()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < b.N; i++ {
		fmt.Fprintln(writer, "Hello world")
	}
}
