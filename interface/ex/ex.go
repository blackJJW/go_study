package main

import "go_study/interface/fedex"

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var sender Sender = &fedex.FedexSender{}
	SendBook("The little price", sender)
	SendBook("greece ", sender)
}
