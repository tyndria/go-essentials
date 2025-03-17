package exercises

import (
	"bytes"
	"os"
)

type Capper struct {
	stdout *os.File
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func (c *Capper) Write(p []byte) (n int, err error) {
	output := bytes.ToUpper(p)
	return c.stdout.Write(output)
}