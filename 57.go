package main

import "fmt"

func main() {
	var n int 
	var sum int

	fmt.Println("n?")
	fmt.Scanln(&n)

	for i:=1;i<n+1;i++ {
		sum = sum+i
	}

	fmt.Println(sum)
}