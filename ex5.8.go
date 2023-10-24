package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a, b int
	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(n, err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}