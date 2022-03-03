package main

import (
	"fmt"

	"github.com/andreykont/ftp-test/pkg/myftp"
)

func main() {
	var err error
	var ftp *myftp.FTP

	// For debug messages: goftp.ConnectDbg("ftp.server.com:21")
	if ftp, err = goftp.Connect("ftp.server.com:21"); err != nil {
		panic(err)
	}

	defer ftp.Close()
	fmt.Println("Successfully connected to", server)
}
