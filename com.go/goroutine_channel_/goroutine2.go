package main

import (
	"fmt"
	"time"
)

/*main() 函数中启动了两个协程：sendData() 通过通道 ch 发送了 5 个字符串，
getData() 按顺序接收它们并打印出来。
如果 2 个协程需要通信，你必须给他们同一个通道作为参数才行。
尝试一下如果注释掉 time.Sleep(1e9) 会如何。*/

/*waitGroup*/
func main() {
	ch := make(chan string)
	go senddata(ch)
	go getData(ch)
	time.Sleep(1e9)
}
func senddata(ch chan string)  {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string)  {
	var input string
	for{
		input = <-ch
		fmt.Println(input)
	}

}