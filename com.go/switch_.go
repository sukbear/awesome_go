package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Print("Linux")
	default:
		fmt.Printf("%s.\n", os)

	}
}

func f2() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good moring")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
func main() {
	f1()
	fmt.Println()
	f2()
}
