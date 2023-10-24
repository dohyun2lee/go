package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var s string
	var time int

	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanln(reader, &s)

	for i:=0;i<len(s);i++ {
		if string(s[i]) == "A" || string(s[i]) == "B" || string(s[i]) == "C" {
			time += 3
		} else if string(s[i]) == "D" || string(s[i]) == "E" || string(s[i]) == "F" {
			time += 4
		} else if string(s[i]) == "G" || string(s[i]) == "H" || string(s[i]) == "I" {
			time += 5
		} else if string(s[i]) == "J" || string(s[i]) == "K" || string(s[i]) == "L" {
			time += 6
		} else if string(s[i]) == "M" || string(s[i]) == "N" || string(s[i]) == "O" {
			time += 7
		} else if string(s[i]) == "P" || string(s[i]) == "Q" || string(s[i]) == "R" || string(s[i]) == "S" {
			time += 8
		} else if string(s[i]) == "T" || string(s[i]) == "U" || string(s[i]) == "V" {
			time += 9
		} else {
			time += 10
		}
	}

	fmt.Print(time)

}