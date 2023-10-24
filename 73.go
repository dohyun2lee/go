package main

import "fmt"

func main() {
	var alp = [26]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	var S [100]string
	var count []int

	for j:=0;j<100;j++ {
		fmt.Println("S?")
		fmt.Scanln(&S[j])
		
		if S[j]=="" {
			break
		}
	}

	fmt.Println(S)
	
	for i:=0;i<26;i++ {
		for j:=0;j<100;j++ {
			if alp[i] == S[j] {
				count = append(count, j)
				break
			} 
		}
	}

	fmt.Println(count)
}