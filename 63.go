package main

import "fmt"

func main() {
	var N, cycle int
	var o int
	//var t int
	var arr [1000]int
	var arr2 [1000]int
	
	fmt.Println("N?")
	fmt.Scanln(&N)
	
	o = N % 10 
	//t = N / 10

	for i:=1;i<1000;i++ {
		arr[0] = o
		arr2[0] = N
		arr[i] = ((arr2[i-1]/10) + (arr2[i-1]%10)) % 10
		arr2[i] = (10 * arr[i-1]) + (arr[i] % 10)
		if arr2[i] != N {
			fmt.Println(i, arr[i], arr2[i])
			continue
		} else {
			cycle = i
			break
		}
	}

	fmt.Println(cycle)
}