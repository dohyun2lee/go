package main

import "fmt"

func main() {
	var N int
	var cnt int
	shape := "***\n* *\n***"
	blank := " "

	fmt.Scanln(&N)

	if N==3 {
		println(shape)
	} else{
		for N>1 {
			N = N/3
			cnt++
		}
	}

	fmt.Println(cnt)

	// for i:=0;i<cnt;i++ {
		
	// }

	for i:=0;i<cnt;i++ {
		shape = shape + shape + shape + "\n" + shape + blank + shape + "\n" + shape + shape + shape
		blank = blank + blank + blank + "\n" + blank + blank + blank + "\n" + blank + blank + blank
	}

	fmt.Println(shape)
}