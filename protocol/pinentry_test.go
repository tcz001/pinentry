package protocol

import "testing"

func TestPinentry(t *testing.T) {
	c := NewClient("pinentry-mac")
	c.SetDesc("Type your passphrase:")
	c.SetPrompt("PIN:")
	c.GetPin()
	c.Close()
}
