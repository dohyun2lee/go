package main

import "fmt"

func main() {
	var a []int
	a = []int{1,2,3}
	a[2] = 0
	fmt.Print(a)

	println("")

	s := make([]int, 5, 10)
	println(len(s), cap(s))
	println(s[4])
	s[2] = 12345
	println(s[2])
	x := s[0:10]
	println(s)
	println(x)
}