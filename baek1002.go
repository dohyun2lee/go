package main

import ( 
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, r1, r2 int
	var T int

	fmt.Scanln(&T)

	for i:=0;i<T;i++ {
		fmt.Scanln(&x1, &y1, &r1, &x2, &y2, &r2)

		distance := math.Pow(float64(((x1-x2)*(x1-x2))+((y1-y2)*(y1-y2))), 1.0/2.0)

		if (x1==x2&&y1==y2&&r1==r2) {
			fmt.Println("-1")
		} else {
			if (r1>r2) {
				if((distance + float64(r2)) < float64(r1)) {
					fmt.Println("0")
				} else if (distance == float64(r2 + r1)) {
					fmt.Println("1")
				} else {
					fmt.Println("2")
				}
			} else if (r2>r1) {
				if ((distance + float64(r1)) < float64(r2)) {
					fmt.Println("0")
				}  else if (distance == float64(r1 + r2)) {
					fmt.Println("1")
				} else {
					fmt.Println("2")
				}
			}
		}
	}
}