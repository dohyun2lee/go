package main

import (
	"fmt"
)

func main() {
	var N, M int
	var C []int
	var a, b int

	fmt.Scanln(&N, &M)

	var A = make([]int, N)
	var B = make([]int, M)

	for i:=0;i<N;i++ {
		fmt.Scanln(a)
		A = append(A, a)
	}
	for i:=0;i<M;i++ {
		fmt.Scanln(b)
		B = append(B, b)
	}

	for i:=0;i<N;i++ {
		for j:=0;j<M;j++ {
			if A[i] == B[j] {
				C = append(C, A[i])
			}
		}
	}


	fmt.Println(len(A) + len(B) - (2*len(C)))
}