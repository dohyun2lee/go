package main

import "fmt"

func main() {
	var A int
	var B int

	fmt.Println("A와 B를 입력하시오")
	fmt.Print("A : ")
	fmt.Scanln(&A)
	fmt.Print("B : ")
	fmt.Scanln(&B)
	fmt.Println("A+B=", A+B)
}