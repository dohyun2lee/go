package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var count int = 0 

	for i:=0;i<n;i++ {
		var str string
		fmt.Fscanln(reader, &str)

		a := make([]bool, 26)

		firstchar := int(str[0]) - 97
		a[firstchar] = true

		for j:=1;j<len(str);j++ {
			charnum := int(str[j]) - 97
			if str[j] != str[j-1] && a[charnum] == true {
				count += 1
				break
			} else {
				a[charnum] = true
			}
		}
	}

	fmt.Fprintln(writer, n-count)
}