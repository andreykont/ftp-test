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

// Connect to remote ftp server.
func Connect(address string) (*FTP, error) {
	var err error
	var conn net.Conn

	if conn, err = net.Dial("tcp", address); err != nil {
		return nil, err
	}
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	ftp := &FTP{conn: conn, address: address, reader: reader, writer: writer}
	ftp.receive()

	return ftp, nil
}

// Close ftp connection
func (ftp *FTP) Close() error {
	return ftp.conn.Close()
}
