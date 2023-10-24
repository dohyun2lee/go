package main

import "fmt"

var arry = [5]int{1:10, 3:30}

type ClassInfo struct {
	Class int
	No int
}

type Student struct {
	Class ClassInfo
	Name string
}

func main() {
	var a Student = Student {
	Class: ClassInfo{Class:1, No:1},
	Name: "John",
	}

	fmt.Println(a.Class.Class)
	fmt.Println(a.Name)

	fmt.Println(arry[1])
}

