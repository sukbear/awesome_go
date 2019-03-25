package main

import (
	"io/ioutil"
	"fmt"
	"os"
)
/*将整个文件的内容读到一个字符串里：

可以使用 io/ioutil 包里的 ioutil.ReadFile() 方法，
该方法第一个返回值的类型是 []byte，里面存放读取到的内容，
第二个返回值是错误，如果没有错误发生，第二个返回值为 nil。
函数 WriteFile() 可以将 []byte 的值写入文件。*/
func main() {
	inputFile := "C:\\Users\\70475\\Desktop\\1.txt"
	outputFile := "C:\\Users\\70475\\Desktop\\2.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
