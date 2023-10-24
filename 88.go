package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var room [14][14]int
	var T, k, n int

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		fmt.Println("k?, n?")
		fmt.Scanln(&k)
		fmt.Scanln(&n)

		for a:=0;a<=k;a++ {
			for b:=1;b<=n;b++ {
				if a == 0 {
					room[0][b] = b
				} else {
					room[a][b] = (room[a-1][b] + room[a][b-1])
				}
				
			}
		}

		fmt.Println(room[k][n])
	}
	
}
