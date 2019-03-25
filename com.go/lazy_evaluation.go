package main

import "fmt"

var resume chan int
func integers() chan int{
	yield:= make(chan int)
	count:=0
	go func() {
		for{
			yield <- count
			count++
		}
	}()
	return yield
}

func genIntegers() int{
	return <- resume
}
func main() {
	resume = integers()
	fmt.Println(genIntegers())
	fmt.Println(genIntegers())
	fmt.Println(genIntegers())
	fmt.Println(genIntegers())
}