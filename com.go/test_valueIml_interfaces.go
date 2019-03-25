package main

import "fmt"

type Stringer interface {
	String() string
}

func pr(args ...int) {
	for _, s := range args {
		fmt.Println(s)
	}
}
func main() {
	pr(1, 2, 5, 6, 12, )
}
