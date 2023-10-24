package main

import "fmt"

func main() {
	var a []int
	var b, sum int

	for i:=0;i<5;i++ {
		fmt.Println("?")
		fmt.Scanln(&b)

		a = append(a, b)
	}

	for i:=0;i<5;i++ {
		for j:=i;j<5;j++ {
			if a[i]>a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			} else {
				continue
			}
		}
	}

	for i:=0;i<5;i++ {
		sum += a[i]
	}

	fmt.Println(sum/5)
	fmt.Println(a[2])
}