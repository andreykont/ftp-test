package myftp

import (
	"bufio"
	"net"
)

type FTP struct {
	conn    net.Conn
	address string

	reader *bufio.Reader
	writer *bufio.Writer
}
