package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func open() {
	file, err := os.Create(filepath.Join(os.TempDir(), "textFile.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "New file content\n")
}

func read() {
	file, err := os.Open(filepath.Join(os.TempDir(), "textFile.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Read File:")
	io.Copy(os.Stdout, file)
}

func append() {
	file, err := os.OpenFile(filepath.Join(os.TempDir(), "textFile.txt"), os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "Appended content\n")
}

func remove() {
	os.Remove(filepath.Join(os.TempDir(), "textFile.txt"))
}

func main() {
	remove()
	open()
	append()
	read()
}
