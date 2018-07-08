package chat

import (
	"fmt"
	"net"
)

func (chat *ChatServer) Dial(addr string) error {
	laddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Println("resolve address:", err)
		return err
	}

	fd, err := net.DialTCP("tcp", nil, laddr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	chat.newConnection(fd, true)
	return nil
}

func inputParse(input string) {
	for _, c := range input {
		switch c {
		case '(':
		default:

		}
	}
}
