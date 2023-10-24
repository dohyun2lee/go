package main

import "fmt"

func main() {
	var A, B, C int
	var prize int

	fmt.Println("주사위")
	fmt.Scanln(&A, &B, &C)

	if A==B&&B==C {
		prize = 10000 + (A * 1000) 
	} else if A==B||A==C||B==C {
		if A==B {
			prize = 1000 + (A * 100)
		} else if A==C {
			prize = 1000 + (A * 100)
		} else {
			prize = 1000 + (B * 100)
		}
	} else {
		if A >= B && A >= C {
			prize = A * 100
		} else if B >=A && B >= C {
			prize = B * 100
		} else {
			prize = C * 100
		}
	}

	fmt.Println(prize)
}