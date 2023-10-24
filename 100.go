package main

import "fmt"

func main() {
	var N int
	var a, b int
	var x []int
	var y []int

	fmt.Println("N?")
	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		fmt.Scanln(&a, &b)
		x = append(x, a)
		y = append(y, b)
	}

	for i:=0;i<N;i++ {
		for j:=i;j<N;j++ {
			if x[i]>x[j] {
				tmp := x[i]
				x[i] = x[j]
				x[j] = tmp
			} else {
				continue
			}
		}
	}

	for i:=0;i<N;i++ {
		for j:=i;j<N;j++ {
			if x[i]==x[j] {
				if y[i]>y[j] {
					tmp := x[i]
					x[i] = x[j]
					x[j] = tmp
				} else {
					continue
				}
			} else {
				continue
			}
		}
	}

	fmt.Println("-------")
	
	for i:=0;i<N;i++ {
		fmt.Println(x[i], y[i])
	}
}