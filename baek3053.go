package main

import (
	"fmt"
	"math"
)

func main() {
	var R int
	pi := math.Pi

	fmt.Scanln(&R)

	r := float64(R)

	fmt.Println(pi*r*r)
	fmt.Println(float64(2)*r*r)
}