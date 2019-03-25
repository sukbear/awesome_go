package main

import "fmt"

type valuable interface {
	getValue() float32
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

func (s stockPosition) getValue() float32 {
	return s.count * s.sharePrice
}

func showvalue(arg valuable) {
	fmt.Println(arg)
}
func main() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showvalue(o)

	o = car{"BMW", "M3", 20}
	showvalue(o)
}
