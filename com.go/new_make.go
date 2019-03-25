package main

import "fmt"


/*make 只适用于映射、切片和信道且不返回指针。若要获得明确的指针，
请使用 new 分配内存。*/
func main() {
	v := make([]int, 2)
	v[0] = 1
	v[1] = 2
	fmt.Println(v)
}
