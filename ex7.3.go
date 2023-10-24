package main

func PrintAvgScore(name string, math int, eng int, history int) int {
	total := math + eng + history
	score := total / 3
	return score
}

func main() {
	print(PrintAvgScore("1번", 80, 74, 95))
}