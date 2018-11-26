package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	// receive ctrl + c
	signal.Notify(signals, syscall.SIGINT)

	fmt.Println("Waiting SIGINT (CTRL + C)")
	<-signals
	fmt.Println("SIGINT arrived")
}
