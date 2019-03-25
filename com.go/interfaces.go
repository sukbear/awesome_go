package main

import "fmt"

type Shaper interface {
	Area() float32
	len() float32
}

type Square struct {
	side float32
}

type Rectangle struct {
	length, width float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) len() float32 {
	return sq.side
}

func (re Rectangle) Area() float32 {
	return re.length * re.width
}

func (re Rectangle) len() float32 {
	return re.length + re.width
}
func main() {

	sq := new(Square)
	sq.side = 5
	var areaIntf Shaper
	areaIntf = sq
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	fmt.Println()
	re := new(Rectangle)
	re.length, re.width = 3, 5
	var shaper Shaper
	shaper = re
	fmt.Println(shaper.len())
	fmt.Println(shaper.Area())
}
