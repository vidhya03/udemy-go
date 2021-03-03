package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func main() {
	tri := triangle{base: 10, height: 20}
	printArea(tri)
}

func (t triangle) getArea() float64 {
	return t.base * t.height
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (s square) getArea() float64 {
	return 0.8
}
