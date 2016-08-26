package protocol

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
)

type pinentry interface {
	SetDesc(desc string)
	SetPrompt(prompt string)
	SetTitle(title string)
	SetOK(ok string)
	SetCancel(cancel string)
	SetError(errorMsg string)
	SetQualityBar()
	SetQualityBarTT(tt string)
	GetPin() []byte
	Confirm() bool
}

type client struct {
	in   io.WriteCloser
	pipe *bufio.Reader
}

func (c *client) SetDesc(desc string) {
	c.in.Write([]byte("SETDESC " + desc + "\n"))
	// ok
	ok, _, _ := c.pipe.ReadLine()
	if bytes.Compare(ok, []byte("OK")) != 0 {
		panic(string(ok))
	}
}

func (c *client) SetPrompt(prompt string) {
	c.in.Write([]byte("SETPROMPT " + prompt + "\n"))
	// ok
	ok, _, _ := c.pipe.ReadLine()
	if bytes.Compare(ok, []byte("OK")) != 0 {
		panic(string(ok))
	}
}

func (c *client) SetTitle(title string) {
	c.in.Write([]byte("SETTITLE " + title + "\n"))
	// ok
	ok, _, _ := c.pipe.ReadLine()
	if bytes.Compare(ok, []byte("OK")) != 0 {
		panic(string(ok))
	}
}

func (c *client) SetOK(okLabel string) {
	c.in.Write([]byte("SETOK " + okLabel + "\n"))
	// ok
	ok, _, _ := c.pipe.ReadLine()
	if bytes.Compare(ok, []byte("OK")) != 0 {
		panic(string(ok))
	}
}

func (c *client) SetCancel(cancelLabel string) {
	c.in.Write([]byte("SETCANCEL " + cancelLabel + "\n"))
	// ok
	ok, _, _ := c.pipe.ReadLine()
	if bytes.Compare(ok, []byte("OK")) != 0 {
		panic(string(ok))
	}
}

func (c *client) SetError(errorMsg string) {
	c.in.Write([]byte("SETERROR " + errorMsg + "\n"))
	// ok
	ok, _, _ := c.pipe.ReadLine()
	if bytes.Compare(ok, []byte("OK")) != 0 {
		panic(string(ok))
	}
}

func (c *client) SetQualityBar() {
	c.in.Write([]byte("SETQUALITYBAR\n"))
	// ok
	ok, _, _ := c.pipe.ReadLine()
	if bytes.Compare(ok, []byte("OK")) != 0 {
		panic(string(ok))
	}
}

func (c *client) SetQualityBarTT(tt string) {
	c.in.Write([]byte("SETQUALITYBAR_TT" + tt + "\n"))
	// ok
	ok, _, _ := c.pipe.ReadLine()
	if bytes.Compare(ok, []byte("OK")) != 0 {
		panic(string(ok))
	}
}

func (c *client) Confirm() bool {
	confirmed := false
	c.in.Write([]byte("CONFIRM\n"))
	// ok
	ok, _, _ := c.pipe.ReadLine()
	if bytes.Compare(ok, []byte("OK")) == 0 {
		confirmed = true
	}
	return confirmed
}

func (c *client) GetPin() []byte {
	c.in.Write([]byte("GETPIN\n"))
	// D
	c.pipe.ReadBytes('D')
	c.pipe.ReadBytes(' ')
	// pin
	pin, _, _ := c.pipe.ReadLine()
	return pin
}

func (c *client) Close() {
	c.in.Close()
	return
}

func NewClient(path string) pinentry {
	cmd := exec.Command(path)
	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	bufout := bufio.NewReader(out)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	// welcome
	welcome, _, _ := bufout.ReadLine()
	if bytes.Compare(welcome, []byte("OK Your orders please")) != 0 {
		panic(welcome)
	}
	return &client{in, bufout}
}
