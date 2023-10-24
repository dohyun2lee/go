package main 

import "fmt"

func main() {
	done := make(chan bool, 1)

	go func() {
		for i:=0;i<10;i++{
			fmt.Println(i)
			
		}
		done <- true
	} ()
	fmt.Println(<-done)
	//<- done
}