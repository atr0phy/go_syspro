package main

import (
	"bytes"
	"fmt"
)

func main() {
	// empty buffer
	var buffer1 bytes.Buffer
	// initialize bytes
	buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
	// initialize string
	buffer3 := bytes.NewBufferString("初期文字列")
	fmt.Println(buffer1.String())
	fmt.Println(buffer2.String())
	fmt.Println(buffer3.String())
}
