package main

import "fmt"

func main() {
	var a, b int

	for true {
		fmt.Scanln(&a, &b)

		if a != 0 {
			if a % b == 0 {
				fmt.Println("multiple")
			} else if b % a == 0 {
				fmt.Println("factor")
			} else {
				fmt.Println("neither")
			}
		} else {
			break
		}
	}
}