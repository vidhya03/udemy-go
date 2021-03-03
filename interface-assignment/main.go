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
	sqe := square{side: 5}
	printArea(tri)
	printArea(sqe)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (s square) getArea() float64 {
	return s.side * s.side
}
