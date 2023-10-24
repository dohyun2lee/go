package main 

import "fmt"

func main() {
	var N, sum int
	s := []int{0}
  
	fmt.Println("N?")
	fmt.Scanln(&N)

	for i:=1;i<166666667;i++ {
		s = append(s, (6*i))
	}
	for j:=0;j<166666667;j++ {
		sum += s[j]

		if N < s[j] {
			fmt.Print(j)
			break
		}
	}
}