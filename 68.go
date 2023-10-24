package main

import "fmt"

func main() {
	var arr []string
	var count []int
	var o_count int
	var N,sum int

	fmt.Println("N?")
	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		var tmp string

		fmt.Println("o?x?")
		fmt.Scanln(&tmp)

		arr = append(arr, tmp)
	}

	for j:=1;j<N;j++ {
		if arr[0]=="o" {
			count = append(count, 1)
		} else {
			continue
		}
		if arr[j]=="o" {
			o_count++
			if arr[j]==arr[j-1] {
				o_count++
			} else {
				continue
			}
		} else {
			o_count = 0
		}
		count = append(count, o_count)
	}
	for k:=0;k<N;k++ {
		sum+=count[k]
	}

	fmt.Println("sum?", sum-1)
}