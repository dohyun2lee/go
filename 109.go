package main

import "fmt"

func main() {
	var N, k, x int
	var X []int
	fmt.Scanln(&N, &k)

	for i:=0;i<N;i++ {
		fmt.Scanln(&x)
		X = append(X, x)
	}

	for i:=0;i<5;i++ {
		for j:=i;j<5;j++ {
			if X[i] < X[j] {
				tmp := X[i]
				X[i] = X[j]
				X[j] = tmp
			}
		}
	}

	fmt.Println("커트라인:", X[k-1])
}