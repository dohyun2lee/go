package main

import "fmt"

func main() {
	var X, sum, n, a int

	fmt.Println("X?")
	fmt.Scanln(&X)

	for i:=0;i<4480;i++ {
		sum += i
		if X <= sum {
			n = i
			a = sum - X
			break
		}
	}

	fmt.Printf("%d / %d", (n-a), (a+1))

}