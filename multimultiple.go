package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanln(&a, &b)

	fmt.Println(multimultiple(a, b))
}

func multimultiple (N, M int) int {
	result := 1

	for i:=0;i<M;i++ {
		result *= N
	}

	return result
}