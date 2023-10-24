package main

import "fmt"

func main() {
	var time int
	var hour int
	var minute int

	fmt.Println("hour?")
	fmt.Scanln(&hour)
	fmt.Println("minute?")
	fmt.Scanln(&minute)

	time = hour*60 + minute

	if time > 45 {
		target_hour := (time - 45) / 60
		target_minute := (time - 45) % 60

		fmt.Println(target_hour, target_minute)
	} else {
		time := time + 1440
		target_hour := (time - 45) / 60
		target_minute := (time - 45) % 60

		fmt.Println(target_hour, target_minute)
	}
}