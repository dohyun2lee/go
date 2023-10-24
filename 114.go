package main

import "fmt"

func main() {
	var person, headcount int
	var floor int
	var Floor []int
	//var FlooR []int
	var cnt bool

	for i:=0;i<15;i++ {
		fmt.Scanln(&floor)
		for j:=0;j<i;j++ {
			if floor == Floor[j] {
				cnt = true
			}
		}
		if cnt != true {
			Floor = append(Floor, floor)
		}
	}

	for i:=0;i<len(Floor);i++ {
		for j:=i;j<len(Floor);j++ {
			if Floor[i] > Floor[j] {
				tmp := Floor[i]
				Floor[i] = Floor[j]
				Floor[j] = tmp
			} else {
				continue
			}
		}
	}

	for i:=0;i<len(Floor);i++ {
		fmt.Scanln(&person)
		headcount = headcount + person 
	}
	
	fmt.Println(headcount)
}