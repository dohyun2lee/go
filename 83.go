package main

import (
	"fmt"
	"bufio"
	"os"
) 

func main () {
	var A, B, C, count int

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A, &B, &C)

	if B>=C {
		fmt.Print("-1")
	} else {
		count = 1+(A/(C-B))
	}

	fmt.Print(count)
} 