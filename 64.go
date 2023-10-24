package main

import "fmt"

func main() {
	var N int
	var arr []int
	var Max, Min int
	fmt.Println("N?")
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		fmt.Println("arr[", i, "]?")
		var tempInt int

		fmt.Scanln(&tempInt)
		arr = append(arr, tempInt)

	}
	Max = arr[0]
	Min = arr[0]
	
	for j := 1; j < N; j++ {
		if arr[j] >= Max {
			Max = arr[j]
		}
		if arr[j] <= Min {
			Min = arr[j]
		}
	}

	fmt.Println(Max, Min)
}
