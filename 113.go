package main

import "fmt"

func main() {
	var N int
	var good map[string]int
	var selection string

	good= make(map[string]int)

	good["핫식스"] = 1260
	good["레쓰비"] = 680
	good["펩시"] = 1190
	good["밀키스"] = 1230

	fmt.Scanln(&N)
	fmt.Println("무엇을 고르시겠습니까?")
	
	for N>0 {
		if N >= 1260 {
			fmt.Println("핫식스\n레쓰비\n펩시\n밀키스\n취소")
			fmt.Scanln(&selection)
			if selection == "취소" {
				fmt.Println("취소")
				break
			} else {
				fmt.Println("남은 돈 : ",N - good[selection])
				N = N - good[selection]
			}	
		} else if N >= 1230 {
			fmt.Println("레쓰비\n펩시\n밀키스\n취소")
			fmt.Scanln(&selection)
			if selection == "취소" {
				fmt.Println("취소")
				break
			} else {
				fmt.Println("남은 돈 : ",N - good[selection])
				N = N - good[selection]
			}
		} else if N >= 1190 {
			fmt.Println("레쓰비\n펩시\n취소")
			fmt.Scanln(&selection)
			if selection == "취소" {
				fmt.Println("취소")
				break
			} else {
				fmt.Println("남은 돈 : ",N - good[selection])
				N = N - good[selection]
			}
		} else if N >= 680 {
			fmt.Println("레쓰비\n취소")
			fmt.Scanln(&selection)
			if selection == "취소" {
				fmt.Println("취소")
				break
			} else {
				fmt.Println("남은 돈 : ",N - good[selection])
				N = N - good[selection]
			}
		} else {
			fmt.Println("금액이 부족합니다.")
			break
		}
		
	}
	
}