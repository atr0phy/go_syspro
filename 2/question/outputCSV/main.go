package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("/tmp/test.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	writer.Write([]string{"fruit", "price", "number"})
	writer.Write([]string{"apple", "100", "2"})
	writer.Write([]string{"orange", "60", "1"})
	writer.Write([]string{"lemon", "120", "3"})
	writer.Flush()
}