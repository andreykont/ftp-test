package main

import (
	"fmt"

	"github.com/andreykont/ftp-test/pkg/myftp"
)

const FtpServer string = "localhost:21"

func main() {
	var err error
	var ftp *myftp.FTP

	if ftp, err = myftp.ConnectFtp(FtpServer); err != nil {
		fmt.Printf("Can't connect remote FTP server")
		return
	}
	defer ftp.CloseFtp()
	fmt.Println("Successfully connected to", FtpServer)
	// authentication here
	// change dir
	// get some file
	// quit session
}
