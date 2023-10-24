package main

import "fmt"

func main() {
	var arr [10000]int
	var arr2 [10000]int
	count := 1

	for i:=0;i<10;i++ {
		fmt.Println("arr[",i,"]?")
		fmt.Scanln(&arr[i])

		arr2[i] = arr[1]%42
	}

	for j:=0;j<10;j++ {
		for k:=0;k<j;k++ {
			if arr2[j]!=arr2[k] {
				count++
			}
		}
	}

	fmt.Println(count)
}