package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	colors := [3]string{"🟢 - green", "🟡 - yellow", "🔴 - red"}
	var light int = 0

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println("Signal received:", sig)
		fmt.Println(os.Environ())
		fmt.Println("AWS_ACCESS_KEY_ID=SOMETHINGSOMETHING")
		fmt.Println("AWS_SECRET_ACCESS_KEY=FLAG(DodgeViper)")
		fmt.Println("AWS_DEFAULT_REGION=🗽 - ⬆️  ⬇️  ➡️  ⬅️  - 🍔")
		done <- true
	}()

	fmt.Println("awaiting signal")

	for {
		select {
		case <-done:
			fmt.Println("exiting")
			return
		default:
			if light > 2 {
				light = 0
			}
			fmt.Println(colors[light])
			light++
			time.Sleep(2 * time.Second)
		}
	}
}
