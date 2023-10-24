package main 

import "fmt"

func main() {
	var A, B [3]int
	var a, ba, b, bb int
	

	for i:=0;i<3;i++ {
		fmt.Println("A?")
		fmt.Scanln(&A[i])
	}

	for i:=0;i<3;i++ {
		fmt.Println("B?")
		fmt.Scanln(&B[i])
	}

	a = (100 * A[0]) + (10 * A[1]) + A[2]
	b = (100 * B[0]) + (10 * B[1]) + B[2]

	fmt.Println("A = %d, B = %d",a,b)
	
	ba = (100 * A[2]) + (10 * A[1]) + A[0]
	bb = (100 * B[2]) + (10 * B[1]) + B[0]

	fmt.Println("new A = %d, new B = %d",ba,bb)

	if ba > bb {
		print(ba)
	} else {
		print(bb)
	}
}