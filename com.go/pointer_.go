package main

import "fmt"
/*与 C 不同，Go 没有指针运算。*/
func main() {
	i,j := 42,2701
	p :=&i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p /= 37

	fmt.Println(j)
}