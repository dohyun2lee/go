package main

import "fmt"

func main() {
	var N int

	fmt.Println("N?")
	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		for j:=0;j<i+1;j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}