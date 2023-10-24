package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var N, M int
	var have []int
	var check []int
	var Bool []int
	var n, m, bool int

	reader := bufio.NewReader(os.Stdin)

	fmt.Scanln(reader, &N)

	for i:=0;i<N;i++ {
		fmt.Scanln(&n)
		have = append(have, n)
	}

	fmt.Scanln(&M)
	
	for i:=0;i<M;i++ {
		fmt.Scanln(&m)
		check = append(check, m)
	}

	for i:=0;i<len(check);i++ {
		for j:=0;j<len(have);j++ {
			if check[i] == have[j] {
				bool = 1
				break
			} else {
				bool = 0
			}
		}
		Bool = append(Bool, bool)
	}

	fmt.Println(Bool)
}	