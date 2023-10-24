package main

import ( 
	"fmt"
	"sort"
)

type Person struct {
	Age	int
	Name string
}

func main() {
	var N int	
	var age int
	var name string

	fmt.Scanln(&N)

	var p []Person
	
	for i:=0;i<N;i++ {
		fmt.Scanln(&age, &name)
		a := Person{}
		a.Age = age
		a.Name = name

		p = append(p, a)
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].Age > p[j].Age
	})

	fmt.Println(p)
}
