package main

import "fmt"

func main() {
	var N, n int
	var a []int

	fmt.Println("N?")
	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		fmt.Println("n?")
		fmt.Scanln(&n)

		a = append(a, n)
	}

	for i:=0;i<len(a);i++ {
		for j:=i;j<len(a);j++ {
			if a[i]>a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			} else {
				continue
			}
		}
	}

	fmt.Println(a)
}