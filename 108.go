package main

import "fmt"

func main() {
	var a, sum int
	var A []int

	for i:=0;i<5;i++ {
		fmt.Scanln(&a)
		A = append(A, a)
	}

	for i:=0;i<5;i++ {
		for j:=i;j<5;j++ {
			if A[i] > A[j] {
				tmp := A[i]
				A[i] = A[j]
				A[j] = tmp
			}
		}
	}

	for i:=0;i<5;i++ {
		sum += A[i]
	}

	fmt.Println(sum/5, A[2])
}