package main

import "fmt"

func main() {
	var A int
	var B int

	fmt.Println("A를 입력하시오")
	fmt.Scanln(&A)
	fmt.Println("B를 입력하시오")
	fmt.Scanln(&B)
	if A>B {
		fmt.Println(">")
	} else if A<B {
		fmt.Println("<")
	} else {
		fmt.Println("==")
	}
}