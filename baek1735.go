package main

import "fmt"

func main() {
	var ax, ay, bx, by int

	fmt.Scanln(&ax, &ay) 
	fmt.Scanln(&bx, &by)

	a := mother_gcd(ay, by) / ay
	b := mother_gcd(ay, by) / by

	fmt.Println(mother_gcd(ay, by))

	fmt.Println((ax*a)+(bx*b), mother_gcd(ay, by))
}

func mother_gcd(a, b int) int {
	var x int

	if a < b {
		for i:=1;i<=a;i++ {
			if a%i==0&&b%i==0 {
				x = i
			}
		}
	} else {
		for i:=1;i<=b;i++ {
			if a%i==0&&b%i==0 {
				x = i
			}
		}
	}

	return x
}