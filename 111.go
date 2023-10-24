package main

import "fmt"

func main() {
	var n int
	var Pibonacchi []int
	Pibonacchi = []int{0,1}

	fmt.Scanln(&n)

	for i:=2;i<=n;i++ {
		Pibonacchi = append(Pibonacchi, (Pibonacchi[i-2] + Pibonacchi[i-1]))
	}

		fmt.Println(Pibonacchi[n])
}