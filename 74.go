package main

import "fmt"

func main() {
	var T, R int
	var S []string

	fmt.Println("T?")
	fmt.Scanln(&T)

	for i:=0;i<T;i++{
		fmt.Println("R?")
		fmt.Scanln(&R)

		for j:=0;j<100;j++ {
			fmt.Println("S?")
			fmt.Scanln(&S[j])
			
			if S[j] == "" {
				break
			}
		}
		for j:=0;j<100;j++ {
			for k:=0;k<R;k++ {
				fmt.Print(S[j])
			}
		}
	}

}