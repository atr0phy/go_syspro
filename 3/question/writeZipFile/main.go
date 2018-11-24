package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("/tmp/output.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	a, err := zipWriter.Create("a")
	if err != nil {
		panic(err)
	}
	io.Copy(a, strings.NewReader("test zip 1"))

	b, err := zipWriter.Create("b")
	if err != nil {
		panic(err)
	}
	io.Copy(b, strings.NewReader("test zip 2"))
}
