package main

import "fmt"

func main() {
	var N int
	
	fmt.Println("N?")
	fmt.Scanln(&N)

	for i:=1;i<10;i++ {
		fmt.Println(N,"x",i,"=",N*i)
	}
}