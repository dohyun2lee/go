package main

import "fmt"

func main() {
	var X int
	var N int
	var arr1 [100]int
	var arr2 [100]int
	var price int

	fmt.Println("X, N?")
	fmt.Scanln(&X, &N)

	for i:=0;i<N;i++ {
		fmt.Println("가격, 개수?")
		fmt.Scanln(&arr1[i], &arr2[i])

		price = price + (arr1[i] * arr2[i])
	}

	if price == X {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}