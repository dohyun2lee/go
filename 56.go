package main

import "fmt"

func main() {
	var T int
	var A, B int
	

	fmt.Println("T?")
	fmt.Scanln(&T)

	var arr [20]int

	for i:=0;i<T;i++ {
		fmt.Scanln(&A, &B)
		arr[i] = A + B
	}

	for i:=0;i<T;i++ {
		fmt.Println(arr)
	}
}