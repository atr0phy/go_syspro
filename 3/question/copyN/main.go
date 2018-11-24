package main

import (
	"crypto/rand"
	"io"
	"os"
)

func CopyN(dest io.Writer, src io.Reader, length int64) {
	_, err := io.Copy(dest, io.LimitReader(src, length))
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Create("/tmp/randomFile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	CopyN(file, rand.Reader, 1024)
}
