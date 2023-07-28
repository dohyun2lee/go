package main

import "fmt"

func main() {
	var age int
	var name string

	fmt.Print("나이는 ? ")
	fmt.Scanln(&age)
	fmt.Print("이름은 ? ")
	fmt.Scanln(&name)
	//fmt.Printf("나이 : %d \n이름 : %s \n",age,name)
	println("나이 : ", age, "이름 : ", name)
	//println("이름 : ",name)

	check(2)

}
func check(val int) {
	switch val.(type) {
	case int:
		println("int")
	default:
		println("not int")
	}
}
