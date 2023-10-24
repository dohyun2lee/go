package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var A, B, sum, ten int
	var a []int
	var a1 []int
	var b []int
	var b1 []int
	var s []int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("A?B?")
	fmt.Fscanln(reader, &A, &B)

	for A>0 {
		a = append(a, (A%10))
		A = A/10
	}
	for B>0 {
		b = append(b, (B%10))
		B = B/10
	}
	for i:=len(a);i>=0;i--{
		a1 = append(a1, a[i])
	}
		
	for i:=len(b);i>=0;i-- {
		b1 = append(b1, b[i])
	}

	if A>B {
		for i:=0;i<len(a);i++ {
			s[i] = a1[i]+b1[i]
		}
	} else {
		for i:=0;i<len(b);i++ {
			s[i] = a1[i]+b1[i]
		}
	}

	ten = 1

	for i:=len(s);i>=0;i-- {
		sum += s[i] * ten
		ten = ten * 10
	}

	fmt.Println(sum)
}