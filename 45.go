package main 

import "fmt"

func main() {
	var name string
	fmt.Print("이름: ")
	fmt.Scanln(&name)
	fmt.Print(name)
}