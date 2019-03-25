package main

import "fmt"
/*切片*/
func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "word"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	p := [6]int{1, 2, 3, 4, 5, 6}
	/*类型 []T 表示一个元素类型为 T 的切片。
	前闭后开
	*/
	s := p[1:4]
	fmt.Println(p)
	fmt.Println(s)
	/*	切片并不存储任何数据，它只是描述了底层数组中的一段。

		更改切片的元素会修改其底层数组中对应的元素。

		与它共享底层数组的切片都会观测到这些修改。*/

	names := [4]string{
		"John", "Tom", "Jack", "Ringo",
	}
	fmt.Println(names)

	m := names[:]
	m[0]="Fuck"
	fmt.Println(m)


/*	切片的零值是 nil。

	nil 切片的长度和容量为 0 且没有底层数组。*/
	var f []int
	fmt.Println(f, len(f), cap(f))
	if f == nil {
		fmt.Println("nil!")
	}
}
