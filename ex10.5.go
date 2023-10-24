package main

import "fmt"

func main() {
	var temp int

	fmt.Scanln(&temp)

	if temp < 10 || temp > 30 {
		fmt.Println("1")
	} else if temp >= 10 && temp < 20 {
		fmt.Println("2")
	} else if temp >= 15 && temp < 25 {
		fmt.Println("3")
	} else {
		fmt.Println("4")
	}
}