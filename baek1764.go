package main

import "fmt"

func main() {
	var N, M int
	var neverseen []string
	var neverheard []string
	var all []string

	fmt.Scanln(&N, &M)

	for i:=0;i<N;i++ {
		var n string
		fmt.Scanln(&n)
		neverseen = append(neverseen, n)
	}
	for i:=0;i<M;i++ {
		var m string
		fmt.Scanln(&m)
		neverheard = append(neverheard, m)
	}
	
	if N > M {
		for i:=0;i<N;i++ {
			for j:=0;j<M;j++ {
				if (neverseen[i] == neverheard[j]) {
					all = append(all, neverseen[i])
				}
			}
		}
	}

	
}