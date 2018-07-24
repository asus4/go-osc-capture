package main

import (
	"log"

	"github.com/hypebeast/go-osc/osc"
)

func player(addr string, port uint, file string) error {
	log.Println("player start")

	client := osc.NewClient(addr, int(port))
	msg := osc.NewMessage("/osc/address")
	msg.Append(int32(-10))
	msg.Append("hello")
	msg.Append(true)

	client.Send(msg)
	return nil
}
