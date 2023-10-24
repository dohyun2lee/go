package main

import "fmt"

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int :
		fmt.Println(int(t), "is int")
	case float64 :
		fmt.Println(float64(t), "is float")
	case string :
		fmt.Println(string(t), "is string")
	default : 
		fmt.Println(t, "is not supported")
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("hello")
	PrintVal(Student{15})
}