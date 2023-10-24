package main

import "fmt"

func main() {
	var N, sum int

	fmt.Scanln(&N)

	for i:=1;i<=N;i++ {
		sum = i + (i/1000000) + ((1%1000000)/100000) + ((i%100000)/10000) + ((i%10000)/1000) + ((i%1000)/100) + ((i%100)/10) + (i%10)
		if sum == N {
			println(i)
			break; 
		}
		if i == N {
			println("0")
		}
	}
}