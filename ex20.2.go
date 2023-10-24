package main

import "github.com/tuckersGo/musthaveGo/ch20/fedex"

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &fedex.FedexSender{}
	Sendbook("어린왕자", sender)
	Sendbook("그리스인 조르바", sender)
}