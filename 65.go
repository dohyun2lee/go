package main

import "fmt"

func main() {
	var arr [50]int
	var Max int
	var count int

	for i:=1;i<10;i++ {
		fmt.Println("arr[",i,"]?")
		fmt.Scanln(&arr[i])
	}

	for j:=1;j<10;j++ {
		if arr[j]>Max {
			Max = arr[j]
			count = j
		}
	}

	fmt.Println(Max, count)
}