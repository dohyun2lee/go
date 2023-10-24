package main

func main() {
	var a interface{} = 1
	var b interface{} = "TOME"

	i := a
	j := a.(int)
	k := b.(string)

	println(i)
	println(j)
	println(k)
}