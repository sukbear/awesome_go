package main

import (
	"fmt"
	"time"
)

func main() {
	ch1,ch2:=make(chan int),make(chan int)
	go pum1(ch1)
	go pum2(ch2)
	go suck_(ch1,ch2)
	time.Sleep(1e9)
}
func pum1(ch chan int){
	for i:=0 ;;i++{
		ch <- i*2
	}
}

func pum2(ch chan int){
	for i:=0;;i++{
		ch<- i+5
	}
}

func suck_(ch1, ch2 chan int){
	for{
		select{
		case v:=<- ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v:=<- ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
			
		}
	}
}