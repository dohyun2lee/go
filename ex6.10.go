package main

import "fmt"

func main() {
	var (
		a int = 10
		b int = 23
	)

	fmt.Println(a, b)

	a, b = b, a

	fmt.Println(a, b)
}
