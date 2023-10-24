package main 

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	
	wait.Add(1)

	go func() {
		for i:=0;i<10;i++{
			fmt.Println(i)
			
		}
		wait.Done()
	} ()

	defer wait.Wait()
}