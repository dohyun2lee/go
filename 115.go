package main

import (
	"fmt"
)

func main() {
	var sign string
	var A int
	var formular []int
	var formular2 []string
	var sum float32
	
	for sign != "=" {
		fmt.Scanln(&A, &sign)
		formular = append(formular, A)
		formular2 = append(formular2, sign)
	} 
	
	sum = float32(formular[0])

	for i:=0;i<len(formular);i++ {
		if formular2[i] == "+" {
			sum = sum + float32(formular[i+1])
		} 
		if formular2[i] == "-" {
			sum = sum - float32(formular[i+1])
		} 
		if formular2[i] == "*" {
			sum = sum * float32(formular[i+1])
		} 
		if formular2[i] == "/" {
			sum = sum / float32(formular[i+1])
		} 
		if formular2[i] == "=" {
			fmt.Println(sum)
		}
	}
	
}