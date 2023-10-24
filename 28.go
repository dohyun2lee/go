package main

import "fmt"

func main() {
	var x interface{}
	x = 1 
	x = "Tom"

	println(x)
}

func println(v interface{}) {
	fmt.Println(v)
}