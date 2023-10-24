package main

import "fmt"

func main() {
	var a, b, c int

	for true {
		fmt.Scanln(&a, &b, &c)

		if a == 0 {
			break
		} else {
			if (a*a + b*b == c*c) {
				fmt.Println("right")
			} else {
				fmt.Println("wrong")
			}
		}
	} 
}