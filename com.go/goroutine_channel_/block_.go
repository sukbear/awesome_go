package main

import (
	"fmt"

)
/*所有的协程都休眠了 - 死锁！*/

/*
ch := make(chan int) 和 ch := make(chan int, 0)创建的都是无缓冲通道。
ch := make(chan int, x)，其中x>0，建立的是有缓存通道。
两者的区别是，无缓冲通道是同步的，写入数据时，必须有协程把数据读走才返回，
有缓存通道是半异步的，当缓冲还没有用满时，写入数据时，只要缓存还有空余空间，
写入到缓冲就可以返回。
这里是无缓冲通道，向out写数据时，协程f1还没有创建，
也就无协程读out的数据，造成程序死锁。
*/
func main() {
	in := make(chan int,1)
	go f1(in)
	in <- 2
	//go f1(in)
}

func f1(in chan int) {
	fmt.Println(<-in)
}
