package main 

import "fmt"

func main() {
	var N int

	fmt.Println("N?")
	fmt.Scanln(&N)
	
	slc := split(N)

	for i:=0;i<len(slc);i++ {
		for j:=i;j<len(slc);j++ {
			if slc[i]<slc[j] {
				tmp := slc[i]
				slc[i] = slc[j]
				slc[j] = tmp
			} else {
				continue
			}
		}
	}

	fmt.Println(slc)
}

func split(n int) []int {
	slt := []int{}
	for n>0 {
		slt = append(slt, n%10)
		n = n/10
	}
	return slt
}