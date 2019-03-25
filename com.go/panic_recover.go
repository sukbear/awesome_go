package main

import "fmt"

/*defer-panic-recover 在某种意义上也是一种像 if，for 这样的控制流机制。

Go 标准库中许多地方都用了这个机制，例如，json 包中的解码
和 regexp 包中的 Complie 函数。Go 库的原则是即使在包的内部使用了panic，
在它的对外接口（API）中也必须用 recover 处理成返回显式的错误。*/
func badCall()  {
	panic("bad end")
}
func test(){
	defer func() {
		if e:=recover(); e!=nil{
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
}
func main() {
	fmt.Print("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}