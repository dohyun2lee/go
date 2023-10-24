package main

import (
	"fmt"
	"math"
)



func main() {
	var N, C, sum int

	fmt.Println("N?")
	fmt.Scanln(&N)

	fmt.Println("C?")
	fmt.Scanln(&C)

	for i:=0;i<N;i++ {
		A := math.Pow(10, float64(N-i))
		B := math.Pow(10, float64(N-1-i))
		A1 := int(A)
		B1 := int(B)
		sum += (C%A1)/B1
		//fmt.Println(sum)
	}



	fmt.Println(sum)
}