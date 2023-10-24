package main

import "fmt"

func main() {
	var N, a int
	var A []int
	
	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		fmt.Scanln(&a)
		A = append(A, a)
	}

	for i:=0;i<N;i++ {
		for j:=i;j<N;j++ {
			if A[i] > A[j] {
				tmp := A[i]
				A[i] = A[j]
				A[j] = tmp
			}
		}
	}

	fmt.Println(A)
}