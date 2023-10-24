package main

import "fmt"

func main() {
	var a [][]int
	var aa []int
	var bb []int
	var b [][]int
	var c [][]int
	var cc []int
	var N, M, x, y, z int

	fmt.Scanln(&N, &M)

	println("1st")
	for i:=0;i<N;i++ {
		for j:=0;j<M;j++ {
			fmt.Scanln(&x)
			aa = append(aa, x)
			//println(i,j)
		}
		a = append(a, aa)
	}
	fmt.Println(a)
	println("2nd")
	for i:=0;i<N;i++ {
		for j:=0;j<M;j++ {
			fmt.Scanln(&y)
			bb = append(bb, y)
			//println(i,j)
		}
		b = append(b, bb)
	}
	fmt.Println(b)
	println("+")
	for i:=0;i<N;i++ {
		for j:=0;j<M;j++ {
			z = a[i][j] + b[i][j]
			cc = append(cc, z)
		}
		c = append(c, cc)
	}

	fmt.Println(c)
}