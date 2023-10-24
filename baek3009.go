package main 

import "fmt"

func main() {
	var X []int
	var Y []int

	var x, y, cord_x, cord_y int 

	for i:=0;i<3;i++ {
		fmt.Scanln(&x, &y)
		X = append(X, x)
		Y = append(Y, y)
	}

	if X[0] == X[1] {
		cord_x = X[2]
	} else {
		if X[0] == X[2] {
			cord_x = X[1]
		} else {
			cord_x = X[0]
		}
	}

	if Y[0] == Y[1] {
		cord_y = Y[2]
	} else {
		if Y[0] == Y[2] {
			cord_y = Y[1]
		} else {
			cord_y = Y[0]
		}
	}

	fmt.Println(cord_x, cord_y)
}