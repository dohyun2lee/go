package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, M, cnt int
	var n, m string
	var array_N []string
	var array_M []string
	
	fmt.Scanln(&N, &M)

	for i:=0;i<N;i++ {
		fmt.Scanln(&n)
		array_N = append(array_N, n)
	}
	for i:=0;i<M;i++ {
		fmt.Scanln(&m)
		array_M = append(array_M, m)
	}

	for i:=0;i<M;i++ {
		for j:=0;j<N;j++ {
			if strings.Contains(array_M[i], array_N[j]) {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
