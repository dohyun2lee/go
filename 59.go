package main

import "fmt"

func main() {
	var T int
	var A, B int
	

	fmt.Println("T?")
	fmt.Scanln(&T)

	for i:=0;i<T;i++ {
		fmt.Scanln(&A, &B)
		fmt.Println("case #",i+1,":",A,"+",B,"=",A+B)
	}
}