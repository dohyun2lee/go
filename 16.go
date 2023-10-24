package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3}
	sliceB := make([]int, len(sliceA), cap(sliceA)*2)

	copy(sliceB, sliceA)
	fmt.Println(sliceB)
	
}