package main 

import "fmt"

type Stringer interface { //1
	String() string
}

type Student struct {
	Name string
	Age int
}

func (s Student) String() string { //2
	return fmt.Sprintf("Hi Im %s, and Im %d years old", s.Name, s.Age) //3
}

func main() {
	student := Student{"cheolsu", 12}
	var stringer Stringer
	
	stringer = student //4

	fmt.Printf("%s\n", stringer.String()) 
}