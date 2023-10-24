package main

import "fmt"

type coordinate struct {
	x int
	y int
}

func main() {
	var N int
	var X, Y int
	var c []coordinate

	fmt.Scanln(&N)

	for i:=0;i<N;i++ {
		fmt.Scanln(&X, &Y)
		a := coordinate{}
		a.x = X
		a.y = Y

		c = append(c, a)
	}

	for i:=0;i<N;i++ {
		for j:=i;j<N;j++ {
			if (c[i].x > c[j].y) {
				tmp := c[i]
				c[i] = c[j]
				c[j] = tmp
			} else if (c[i].x == c[j].y) {
				if (c[i].y > c[j].y) {
					tmp := c[i]
					c[i] = c[j]
					c[j] = tmp
				}
			}
		}
	}

	fmt.Println(c)
}