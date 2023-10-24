package main

import (
	"fmt"
)

func main() {
	var M, N, count int
	sosu := []int{2}

	for i:=3;i<10000;i++ {
		if (i%2!=0) {
			for j:=1;j<i;j++ {
				if (i%(j+1)==0) {
					count++
				}
			}
		}
		if count == 1 {
			sosu = append(sosu, i)
		}
	}

	var min, max int
	var bet []int

	fmt.Println("M?")
	fmt.Scanln(&M)
	fmt.Println("N?")
	fmt.Scanln(&N)

	for i:=0;i<len(sosu);i++ {
		if ((sosu[i] > M)||(sosu[i] < N)) {
			bet = append(bet, sosu[i])
		} 
	}

	min = sosu[0]
	max = sosu[len(sosu)]

	fmt.Println(max)
	fmt.Println(min)
}