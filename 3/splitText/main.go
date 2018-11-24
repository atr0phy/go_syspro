package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `line1
line2
line3
`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("%#v\n", line)
	}

	// or

	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
