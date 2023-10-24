package main 

import (
	"fmt"
)

func main() {
	var N, M int
	var Pokemon = make([]string, N)
	var pokemon string
	var Search = make([]interface{}, M)
	var input interface{}
	

	fmt.Scanln(&N, &M)
	

	for i:=0;i<N;i++ {
		fmt.Scanln(&pokemon)
		Pokemon = append(Pokemon, pokemon)
	}

	for i:=0;i<M;i++ {
		fmt.Scanln(&input)
		Search = append(Search, input)
	}

	for i:=0;i<M;i++ {
		if _, ok := Search[i].(string); ok {
			for j:=0;j<N;j++ {
				if Search[i] == Pokemon[j] {
					fmt.Println(j+1)
				}
			}
		}
		if _, ok := Search[i].(int); ok {
			a := Search[i].(int)
			fmt.Println(Pokemon[a])
		}		

	}
}