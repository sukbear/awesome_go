package main


import (
	"net"
	"fmt"
)
/*在main()创建了一个net.Listener的变量，他是一个服务器的基本函数：
用来监听和接收来自客户端的请求
（来自localhost即IP地址为127.0.0.1端口为50000基于TCP协议）
。这个Listen()函数可以返回一个error类型的错误变量。
用一个无限for循环的listener.Accept()来等待客户端的请求。
客户端的请求将产生一个net.Conn类型的连接变量。
然后一个独立的协程使用这个连接执行doServerStuff()，
开始使用一个512字节的缓冲data来读取客户端发送来的数据并且把它们打印到服务器的终端，
len获取客户端发送的数据字节数；
当客户端发送的所有数据都被读取完成时，
协程就结束了。这段程序会为每一个客户端连接创建一个独立的协程。
必须先运行服务器代码，再运行客户端代码。*/
func main() {
	fmt.Println("start the server")
	listener, err := net.Listen("tcp", "localhost:5050")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error listening", err.Error())
			return //终止程序
		}
		go doServerStuff(conn)
	}
}
func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}
