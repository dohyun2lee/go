package main

import "fmt"

func main() {
	var r [][]int
	r = append(r, []int{})
	r = append(r, []int{})
	r = append(r, []int{})

	r[0] = append(r[0], 4, 5, 6)
	fmt.Println(r)

}
