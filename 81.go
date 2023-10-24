package main

import (
	"fmt"
	"bufio"
	"os"
	//"math"
)

func main() {
	var N, count int
	var s string

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &N)

	defer writer.Flush()

	for i:=0;i<N;i++ {
		fmt.Fscanf(reader, "%s\n", &s)
		for j:=0;j<len(s);j++ {
			for k:=0;k<len(s);k++ {
				if s[j] == s[k] {
					if (k-j) > 1 {
						continue
					} else {
						count += 1
					}
				}
			}
		}
	}

	fmt.Println(count)
}