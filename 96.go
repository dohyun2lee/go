package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N, k int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("?")
	fmt.Fscanln(reader, &N, &k)
	
	var X = make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Fscanf(reader, "%d", &X[i])
	}
	
	for i:=0;i<len(X);i++ {
		for j:=i;j<len(X);j++ {
			if X[i]<X[j] {
				tmp := X[i]
				X[i] = X[j]
				X[j] = tmp
			} else {
				continue
			}
		}
	}
	
	fmt.Println(X[k-1])
}