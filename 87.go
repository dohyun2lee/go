package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var T, H, W, N int

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &T)

	for i:=0;i<T;i++ {
		fmt.Fscanln(reader, &H, &W, &N)
		h := (N % H)
		w := (N / H) + 1
		
		if w < 10 {
			fmt.Println(h,"0",w)
		} else {
			fmt.Println(h,w)
		}
	}
}