package main

import (
	"fmt"

	"github.com/tcz001/pinentry/protocol"
)

func main() {
	client := protocol.NewClient("pinentry-mac")
	p := client.GetPin()
	fmt.Println(string(p))
	client.Close()
}
