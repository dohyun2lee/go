package main

import "fmt"

func main() {
	var C, N, sum int
	var avg, percent float64
	var score [10000]int
	var count int

	fmt.Println("C?")
	fmt.Scanln(&C)

	fmt.Println("N?")
	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		fmt.Println("score?")
		fmt.Scanln(&score[i])

		sum += score[i]
	}

	avg = float64(sum / N)

	for j:=0;j<N;j++ {
		if avg<float64(score[j]){
			count++
		} 
	}
	fmt.Println(count)
	percent = float64(count)/float64(N)*100
	fmt.Println(percent)
}