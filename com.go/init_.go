package main

import (
	"os"
	"fmt"
)

var(
	home = os.Getenv("HOME")
	user = os.Getenv("USER")
	path = os.Getenv("GOPATH")
)

func main() {
	fmt.Println(home,user,path)
}
