package main

import (
	"time"
	"fmt"
)

func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(2000)
		fmt.Println(message)
	}()
}

func main() {
	Announce("fuck", 100)

	messages := make(chan string, 2)

	go func() { messages <- "ping" }()
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
