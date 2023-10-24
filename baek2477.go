package main

import "fmt"

func main() {
	var (
		//field map[int]int
		melon, a, b                int
		direction                  []int
		distance                   []int
		target                     []int
		target_num                 []int
		target_area, height, width int
	)

	fmt.Scanln(&melon)

	for i := 0; i < 6; i++ {
		fmt.Scanln(&a, &b)
		direction = append(direction, a)
		distance = append(distance, b)
	}

	for i := 0; i < 6; i++ {
		for j := i + 1; j < 6; j++ {
			if direction[i] == direction[j] {
				target = append(target, direction[i])
			}
		}
	}

	for i:=0;i<6;i++ {
		if direction[i] == target[0] {
			target_num = append(target_num, i)
		}
	}

	for i:=0;i<6;i++ {
		if direction[i] == target[1] {
			target_num = append(target_num, i)
		}
	}

	for i := 0; i < 6; i++ {
		if direction[i] == 2 {
			width += distance[i]
		}
		if direction[i] == 4 {
			height += distance[i]
		}
	}

	target_area = (width * height) - (distance[target_num[1]] * distance[target_num[2]])

	fmt.Println(melon * target_area)
}
