package main

import "fmt"

func main() {
	m := make(map[int]string)

	m[1] = "coffee"
	m[2] = "tea"

	println(m[1])
	println(m[2])

	m[1] = "beverage"

	println(m[1])

	for key, val := range m {
		fmt.Println(key, val)
	}

	for val := range m {
		fmt.Println(val)
	}

	delete(m, 2)

	println(m[2])

	if (m[2] == "") {
		println("null")
	}
}