package main

import "fmt"

func main() {
	var N, count int

	fmt.Println("N?")
	fmt.Scanln(&N)

	if N<100 {
		count = N 
		fmt.Println(count)
	} else if N<1000 {
		for i:=100;i<N;i++ {
			if (i%10)==(i%100/10)+((i%100/10)-(i/100)) {
				count++
			}
		}
		fmt.Println(count+99)
	} else {
		for i:=100;i<N;i++ {
			if (i%10)==(i%100/10)+((i%100/10)-(i/100)) {
				count++
			}
		}
		fmt.Println(count+99)
	}

}