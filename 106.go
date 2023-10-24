package main

import "fmt"

func main() {
	var N, M, a, b int
	var aa, bb, c []int
	var A [][]int

	fmt.Scanln(&N, &M)

	for i:=0;i<N*M;i++ {
			fmt.Scanln(&a)
			aa = append(aa, a)		
	}

	for i:=0;i<N*M;i++ {
			fmt.Scanln(&b)
			bb = append(bb, b)	
	}

	for i:=0;i<N*M;i++ {
			c = append(c, aa[i]+bb[i])		
	}

	for i:=0;i<N;i++ {
		for j:=0;j<M;j++ {
			A[i][j] = c[(M*i)+j]
		}
		
	}
	fmt.Println(A)
}