package main

import ( 
	"fmt"
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

	for i:=0;i<N;i++ {
		for j:=i;j<N;j++ {
			if p[i].Age < p[j].Age {
				tmp := p[i]
				p[i] = p[j]
				p[j] = tmp
			} else {
				continue
			}
		}
	}

	fmt.Println(p)
}
