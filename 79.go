package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var s string
	
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &s)

	s ,_ = reader.ReadString('\n')

	words := strings.Split(s, " ")

	var count int

	for i := range words {
		if words[i] != "\n" && words[i] != " " {
			count++	
		}
	}

	fmt.Print(count)
}