package main

import "fmt"

func main() {
	m := make(map[int]int)
	
	m[1] = 0
	m[2] = 1
	m[3] = 2

	delete(m, 3)
	delete(m, 4)

	fmt.Println(m[3])
}