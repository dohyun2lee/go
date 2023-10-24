package main

import (
	"fmt"
)

func main() {
	var N int
	var Max int
	var arr [10000]int
	var arr2 [10000]float32
	var sum2, avg, N2 float32
	var s [10000]float32

	fmt.Println("N?")
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		fmt.Println(i+1, "번 째 과목 점수?")
		fmt.Scanln(&arr[i])
	}

	for j := 0; j < N; j++ {
		if arr[j] > Max {
			Max = arr[j]
		}
	}

	for k := 0; k < N; k++ {
		arr2[k] = float32(arr[k])
		Max2 := float32(Max)
		s[k] = (arr2[k] / Max2) * 100
		fmt.Println(s[k])
		sum2 += s[k]
	}
	N2 = float32(N)
	avg = sum2 / N2

	fmt.Println(avg)
}
