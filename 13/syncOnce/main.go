package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("initialize")
}

var once sync.Once

func main() {
	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
