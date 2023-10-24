package main

import "fmt"

func main() {
    var s []int
 
    if s == nil {
        println("Nil Slice")
    }
    println(len(s), cap(s))

	a := []int{0,1,2,3,4,5,6,7}
	//a = a[2:5]
	fmt.Print(a[2:4])
	fmt.Print(a[1:])
	fmt.Print(a[:4])
	a = append(a, 10, 8, 9)
	fmt.Print(a)
}