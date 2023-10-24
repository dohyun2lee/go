package main

import "fmt"

func main() {
	var N int

	fmt.Scanln(&N)

	if(N>1) {
		Movie(N)
	} else {
		fmt.Println("666");
	}
}

func Movie(n int) {
	count := 1
	prev := 0
	num := 0

	for {
		if (((prev % 100000) / 10 == 666) && (prev % 10 != 6)) {
			for i:=0;i<1000;i++ {
				if (count == n) {
					fmt.Println(prev * 1000 + num)
					return
				}
				num++
				count++
			}
			prev++
		} else if (prev % 1000 == 666) {
			num = 0
			for i:=0;i<1000;i++ {
				if (count == n) {
					fmt.Println(prev * 1000 + num)
				}
				num++
				count++
			}
			prev++
		} else if (prev % 100 == 66) {
			num = 600
			for i:=0;i<100;i++ {
				if (count == n) {
					fmt.Println(prev * 1000 + num)
				}
				num++
				count++
			}
			prev++
		} else if (prev % 10 == 6) {
			num = 660
			for i:=0;i<10;i++ {
				if (count == n) {
					fmt.Println(prev * 1000 + num)
				}
				num++
				count++
			}
			prev++
		} else {
			num = 666
			if(count == n) {
				fmt.Println(prev * 1000 + num)
				return
			}
			count++
			prev++
		}
	}
}