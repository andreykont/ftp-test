package myftp

import (
	"errors"
	"fmt"
	"strings"
)

func (ftp *FTP) receiveOneLine() (string, error) {
	line, err := ftp.reader.ReadString('\n')
	return line, err
}

// receive bytes from terminal
// get one line and multiline string too
func (ftp *FTP) receive() (string, error) {
	line, err := ftp.receiveOneLine()

	if err != nil {
		return line, err
	}

	if (len(line) >= 4) && (line[3] == '-') {
		// get multiline response
		// not implemented yet
		// read many lines by loop
	}
	return line, err
}

// send bytes in terminal
func (ftp *FTP) send(command string, arguments ...interface{}) error {
	command = fmt.Sprintf(command, arguments...)
	command += "\r\n"
	if _, err := ftp.writer.WriteString(command); err != nil {
		return err
	}
	if err := ftp.writer.Flush(); err != nil {
		return err
	}
	return nil
}

func (ftp *FTP) makeCommand(result string, command string, args ...interface{}) (line string, err error) {
	if err = ftp.send(command, args...); err != nil {
		return
	}
	if line, err = ftp.receive(); err != nil {
		return
	}
	if !strings.HasPrefix(line, result) {
		err = errors.New(line)
		return
	}
	return
}
