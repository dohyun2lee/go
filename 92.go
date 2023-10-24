package main 

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var t, N, count int

	fmt.Println("t?")
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)

	for i:=0;i<t;i++ {
		fmt.Fscanf(reader, "%d", &N)

		var t_count int
		
		for j:=0;j<N;j++ {
			if (N%(j+1)==0) {
				t_count++
			}
		}

		if t_count == 2 {
			count++	
		}
	} 
	fmt.Println(count)
}