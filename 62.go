package main

import "fmt"

func main() {
	var N, X int

	fmt.Println("N? X?")
	fmt.Scanln(&N, &X)

	var arr [10000]int

	fmt.Println("수열 A?", N, "개 입력 가능")
	for i:=0;i<N;i++ {
			fmt.Scanln(&arr[i])
	}

	fmt.Println("X보다 작은 수")

	for i:=0;i<N;i++ {
		if arr[i] < X {
			fmt.Println(arr[i])
		}
	}
}