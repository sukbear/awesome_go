package main

import "fmt"

/*结构体字段使用点号来访问。*/
/*结构体字段可以通过结构体指针来访问。*/

/*如果我们有一个指向结构体的指针 p，
那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，
所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。*/
type Vertex struct {
	X ,Y int
}
var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)
func main() {
	fmt.Println(v1,v2,v3,p)
}