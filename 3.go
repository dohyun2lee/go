package main


func main() {
	
	sum := 0

	for i:=1; i <20; i++ {
		sum += i
	}
	println(sum)


	hello(4)
}

func hello(x int) {
	

	for x <100 {
		x *= 2
	}

	println(x)
}

