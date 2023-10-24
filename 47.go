package main 

import "fmt"

func main() {
	var year int
	fmt.Println("불기 연도를 입력하시오 (1000이상 3000이하)")
	fmt.Scanln(&year)
	if year>=1000 && year<=3000 {
		fmt.Println("서기 연도는 :",year-543)
	} else { fmt.Println("잘못된 연도입니다.") }
	
}