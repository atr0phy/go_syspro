package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("start timer: %s\n", time.Now())
	<-time.After(3 * time.Second)
	fmt.Printf("end timer: %s\n", time.Now())
}
