package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randValue := rand.Intn(100)
	var guessValue int
	cnt := 0

	for {
		fmt.Println("숫자값을 입력하세요.")
		fmt.Scanln(&guessValue)

		if guessValue < randValue {
			fmt.Println("입력하신 값이 더 작습니다.")
		} else if guessValue > randValue {
			fmt.Println("입력하신 값이 더 큽니다.")
		} else {
			fmt.Println("축하합니다. 숫자를 맞추셨습니다. 시도횟수: ", cnt)
			break
		}

		cnt++
	}

}