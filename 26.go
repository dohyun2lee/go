package main

import "fmt"

type Point struct {
	x, y int
}

func (p *Point) add(a int) {
	p.x += a
	p.y += a
	fmt.Printf("%p\n", p)
}

func (p Point) mul(a int) {
	p.x *= a
	p.y *= a
	fmt.Printf("%p\n", &p)
}

func main() {
	p := Point{3,4}
	fmt.Printf("%p\n", &p)

	p.add(10)
	p.mul(100)
}