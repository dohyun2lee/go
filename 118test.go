package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var order = make([]int, n)
	
	for i:=0;i<n;i++ {
		fmt.Println(order[i])
		order[i]++
		fmt.Println(order[i])
	}

	fmt.Println(order)
}