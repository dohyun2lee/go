package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 101
	fmt.Println(<-c)
}