package main

import (
	"math"
	"fmt"
)

type Circle struct {
	radius float32
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var areaIntf Shaper
	sq := new(Square)
	sq.side = 5
	areaIntf = sq

	if t, ok := areaIntf.(*Square); ok {
		fmt.Println("the type of areaIntf is : %T\n", t)
	} else{
	fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}
