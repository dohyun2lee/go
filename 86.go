package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var x, A, B, V int

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &A, &B, &V)

	for i:=1;i<(V/(A-B)+1);i++ {
		x = (A*i)-(B*i)
		if (x+B) >= V {
			fmt.Println(i)
			break
		}
	}
}