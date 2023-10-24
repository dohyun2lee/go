package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, cnt int

	fmt.Scanln(&cnt)

	var divisor = make([]int, cnt)

	for i:=0;i<cnt;i++ {
		fmt.Fscanf(reader, "%d", &divisor[i])
	}

	for i:=0;i<cnt;i++ {
		for j:=i;j<cnt;j++ {
			if divisor [j] < divisor[i] {
				tmp := divisor[i]
				divisor[i] = divisor[j]
				divisor[j] = tmp
			}
		}
	}

	N = divisor[0] * divisor[cnt-1]

	fmt.Println(N)
}