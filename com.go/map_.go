package main

import (
	"fmt"
)

/*映射是方便而强大的内建数据结构，
它可以关联不同类型的值。其键可以是任何相等性操作符支持的类型，
如整数、浮点数、复数、字符串、指针、接口（只要其动态类型支持相等性判断）、
结构以及数组。 切片不能用作映射键，因为它们的相等性还未定义。与切片一样，
映射也是引用类型。 若将映射传入函数中，
并更改了该映射的内容，则此修改对调用者同样可见。*/
func main() {
	timeZone := map[string]int{"UTC":  0*60*60,
		"EST": -5*60*60,
		"CST": -6*60*60,
		"MST": -7*60*60,
		"PST": -8*60*60,}
	for arg :=range timeZone{
		fmt.Println(arg)
	}
	fmt.Println()


	delete(timeZone,"EST")
	for arg :=range timeZone{
		fmt.Println(arg)
	}

}

