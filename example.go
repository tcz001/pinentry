package main

import (
	"fmt"

	"github.com/tcz001/pinentry/protocol"
)

func main() {
	client, _ := protocol.NewPinentryClient()
	p, _ := client.GetPin()
	fmt.Println(string(p))
	client.Close()
}
