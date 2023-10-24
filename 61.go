package main 

import "fmt"

func main() {
	var N int

	fmt.Println("N?")
	fmt.Scanln(&N)

	for i:=1;i<N+1;i++ {
		for j:=0;j<N-i;j++ {
			fmt.Print(" ")
		}
		for j:=0;j<i;j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}