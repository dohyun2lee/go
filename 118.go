package main

import ( 
	"fmt"
	"bufio"
	"os"
)

type person struct {
	Height int
	Weight int
}

func main() {
	var N int
	var height, weight int
	var A []person

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &N)

	order := make([]int, N)

	for i:=0;i<N;i++ {
		fmt.Scanln(&height, &weight)
		a := person{}
		a.Height = height
		a.Weight = weight

		A = append(A, a)
	}

	for i:=0;i<N;i++ {
		order[i]++
		for j:=0;j<N;j++ {
			if ((A[i].Height < A[j].Height)&&(A[i].Weight < A[j].Weight)) {
				order[i]++
			}
		}
	}

	fmt.Println(order)
}