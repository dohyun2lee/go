package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Println("?")
	fmt.Scanln(&N)
	
	var X = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Scanf("%d", &X[i])
	}
	
	for i:=0;i<len(X);i++ {
		for j:=i;j<len(X);j++ {
			if X[i]>X[j] {
				tmp := X[i]
				X[i] = X[j]
				X[j] = tmp
			} else {
				continue
			}
		}
	}
	
	fmt.Println(X)
}