package main

func main() {
	var a = [...]int{1,2,4,6,9}

	println(a[2])

	var b = [2][4]int {
		{1,2,3,4},
		{5,6,7,8},
	}

	println(b[0][2])
}