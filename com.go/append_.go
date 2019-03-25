package main

import "fmt"
/*func append(slice []T, 元素 ...T) []T*/
func main() {
	x :=[]int{1,2,3}
	x = append(x,4,5,6,6,6,6,7)
	fmt.Println(x)
}