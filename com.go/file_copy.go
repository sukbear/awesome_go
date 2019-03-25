package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	copyFile("C:\\Users\\70475\\Desktop\\1.txt","C:\\Users\\70475\\Desktop\\2.txt")
}
func copyFile(dstName, srcName string)(write int64,err error){
	src, err := os.Open(srcName)
	if err != nil{
		fmt.Println("err")
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil{
		return
	}
	return io.Copy(dst,src)
}
