package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("./*.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)
}
