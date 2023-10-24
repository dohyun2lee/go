package main 

import "fmt"

func main() {
	var N, count int

	fmt.Println("N?")
	fmt.Scanln(&N)
	
	if (N%5%3 != 0) {
		count = -1
	} else {
		count = (N/5) + (N%5/3)
	}

	fmt.Println(count)
}