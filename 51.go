package main

import "fmt"

func main(){
	var x float64
	var y float64

	fmt.Println("x?")
	fmt.Scanln(&x)
	fmt.Println("y?")
	fmt.Scanln(&y)

	if x>0&&y>0 {
		fmt.Println("Quadrant 1")
	} else if x>0&&y<0 {
		fmt.Println("Quadrant 4")
	} else if x<0&&y>0 {
		fmt.Println("Quadrant 2")
	} else {
		fmt.Println("Quadrant 4")
	}
}