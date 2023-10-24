package main

import "fmt"

func main() {
	var time int
	var hour int
	var minute int
	var cook int

	fmt.Println("hour and minute?")
	fmt.Scanln(&hour, &minute)
	fmt.Println("cook?")
	fmt.Scanln(&cook)

	time = hour*60 + minute + cook

	if time<1440 {
		target_hour := time / 60
		target_minute := time % 60

		fmt.Println(target_hour, target_minute)
	} else {
		target_hour := (time - 1440) / 60
		target_minute := (time - 1440) % 60

		fmt.Println(target_hour, target_minute)
	}
}