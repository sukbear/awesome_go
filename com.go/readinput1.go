package main

import (
	"fmt"
	"bufio"
	"os"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	inputReader            *bufio.Reader
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)

	fmt.Println("======")
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("enter the input: ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(input)
	}
}
