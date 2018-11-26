package main

import (
	"fmt"
	"time"
)

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is finished")
}

func main() {
	fmt.Println("start sub()")
	go sub()
	//closure
	go func() {
		fmt.Println("sub2() is running")
		time.Sleep(time.Second)
		fmt.Println("sub2() is finished")
	}()
	time.Sleep(2 * time.Second)
}
