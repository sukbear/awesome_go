package main
/*

使用锁的情景：
访问共享数据结构中的缓存信息
保存应用程序上下文和状态信息数据

使用通道的情景：
与异步操作的结果进行交互
分发任务
传递数据所有权

*/
type Task struct {
}

func process(t chan *Task) {

}
func Worker(in, out chan *Task) {
	for {
		t := <-in
		//process(t)
		out <- t
	}
}
