package main

const YSize  = 10
const XSize  = 10
func main() {

	type Transform [3][3]float64  // 一个 3x3 的数组，其实是包含多个数组的一个数组。
	type LinesOfTest [][]byte //包含多个字节切片的一个切片。
	text := LinesOfTest{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	// 分配顶层切片。
	picture := make([][]uint8, YSize) // 每 y 个单元一行。
	// 遍历行，为每一行都分配切片
	for i := range picture {
		picture[i] = make([]uint8, XSize)
	}
}