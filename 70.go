package main

import "fmt"

func main() {
	var self []int
	var nonself []int

	for i:=1;i<10001;i++ {
		if i<10 {
			nonself = append(nonself,i+i)
		} else if i<100 {
			nonself = append(nonself,i+(i/10)+(i%10))
		} else if i<1000 {
			nonself = append(nonself,i+(i/100)+(i%100/10)+(i%100%10))
		} else if i<10000{
			nonself = append(nonself,i+(i/1000)+(i%1000/100)+(i%1000%100/10)+(i&1000%100%10))
		} else {
			nonself = append(nonself,i+1)
		}
	}


	for j:=1;j<10001;j++ {
		for k:=1;k<len(nonself);k++ {
			if nonself[k]<=10000 {
				if j!=nonself[k] {
				self = append(self,j)
				}
			} 
		}
		fmt.Println(self[j])
	}

}
