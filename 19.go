package main

import "test/testlib"

func main() {
	song := testlib.GetMusic("dog")
	println(song)
}