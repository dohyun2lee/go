package main

import (
	"fmt"
	"bufio"
	"os"
)

type person struct {
	Age []int
	Name []string
}

func main() {
	var N int
	var age int
	var name string

	reader := bufio.NewReader(os.Stdin)

	for i:=0;i<N;i++ {
		fmt.Fscanln(reader, &age, &name)
		
		var Person person
		Person.appendPerson(age, name)
	}

	
	
}

func (data *person) appendPerson(age int,name string) {
	data.Age = append(data.Age, age)
	data.Name = append(data.Name, name)
}