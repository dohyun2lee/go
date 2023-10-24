package main

import "fmt"

func main() {
	var N int
	factorial := 1
	
	fmt.Scanln(&N)

	if N==0 {
		factorial = 1
	} else {
		for i:=1;i<=N;i++ {
			factorial = factorial * i
		}
	}

	fmt.Println(factorial)
}