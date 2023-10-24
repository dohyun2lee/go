package main 

import (
	"fmt"
	"strconv"
)

func main() {
	var N, M, check int
	var Pokemon = make([]string, N)
	var pokemon, input string
	var err error	

	fmt.Scanln(&N, &M)
	

	for i:=0;i<N;i++ {
		fmt.Scanln(&pokemon)
		Pokemon = append(Pokemon, pokemon)
	}

	for i:=0;i<M;i++ {
		fmt.Scanln(&input)
		check, err = strconv.Atoi(input)
		if err == nil {
			fmt.Println(Pokemon[check-1])
		} else {
			for j:=0;j<N;j++ {
				if Pokemon[j] == input {
					fmt.Println(j+1)
				}
			}	
		}
	}
	
}