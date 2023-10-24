package main

import "fmt"

type Product struct {
	Name string
	Price int
}

func main() {
	m := make(map[int]Product) 

	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}

	for k, v := range m {
		fmt.Println(k, v)
	}
}