package main 

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	println(area)
}