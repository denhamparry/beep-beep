package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig, "some creds and a flag dumped")
		done <- true
	}()

	fmt.Println("awaiting signal")

	for {
		select {
		case <-done:
			fmt.Println("exiting")
			return
		default:
			fmt.Println("!!beep beep!!!")
			time.Sleep(2 * time.Second)
		}
	}
}
