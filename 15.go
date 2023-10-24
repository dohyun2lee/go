package main

import "fmt"

func main() {
	sliceA := make([]int, 0, 3)

	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA)

	sliceB := []int{1,3,9}

	sliceA = append(sliceA, sliceB...)

	fmt.Println(sliceA)
	fmt.Println(len(sliceA), cap(sliceA))
}