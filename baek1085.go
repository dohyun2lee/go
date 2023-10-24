package main

import (
	"fmt"
)

func main() {

	var x, y, w, h int

	fmt.Scanln(&x, &y, &w, &h)

	var xsub, ysub, a, b int

	xsub = w - x
	ysub = h - y

	if x > y {
		a = y
	} else {
		a = x
	}

	if xsub > ysub {
		b = ysub
	} else {
		b = xsub
	}

	if a > b {
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}
}