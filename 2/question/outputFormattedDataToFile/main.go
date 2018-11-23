package main

import (
	"fmt"
	"os"
)

func main(){
	file, err := os.Create("/tmp/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "%d, %s, %f", 1, "text", 1.111)
}