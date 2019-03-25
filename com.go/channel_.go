package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done

	fmt.Println("=================")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
