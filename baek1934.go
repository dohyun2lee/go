package main

import "fmt"

func main() {
	var T, a, b int

	fmt.Scanln(&T)

	for i:=0;i<T;i++ {
		fmt.Scanln(&a, &b)
		fmt.Println(a*b/gcd(a,b))
	}
}

func gcd(a, b int) int {
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