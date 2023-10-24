package main 

import "fmt"

func main() {
	var year int
	fmt.Println("연도를 입력하시오")
	fmt.Scanln(&year)

	if year%4==0&&year%100!=0 {
		fmt.Println("1")
	} else if year%400==0 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}