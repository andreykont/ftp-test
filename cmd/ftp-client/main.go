package main

import (
	"fmt"

	"github.com/andreykont/ftp-test/pkg/myftp"
)

const FtpServer = "localhost:21"

func main() {
	var err error
	var ftp *myftp.FTP

	if ftp, err = myftp.Connect(FtpServer); err != nil {
		panic(err)
	}
	defer ftp.Close()
	fmt.Println("Successfully connected to", FtpServer)
	// authentication
	// change dir
	// get some file
	// quit session
}
