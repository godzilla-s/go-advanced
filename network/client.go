package network

import (
	"fmt"
	"io"
	"net"
	"time"
)

func Client(addr string) {
	laddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		panic(err)
	}

	c, err := net.DialTCP("tcp", nil, laddr)
	if err != nil {
		fmt.Println("dial:", err)
		return
	}

	//handleWrite(c)
	//handleRead(c)
	clientHeartbeat(c)
}

func handleWrite(c *net.TCPConn) {
	for {
		time.Sleep(5 * time.Second)

		n, err := c.Write([]byte("abcdefghijk"))
		if err != nil {
			fmt.Println("write:", err)
			continue
		}
		fmt.Println("write:", n)
		c.CloseWrite() // 关闭写端口
	}
}

func handleRead(c *net.TCPConn) {
	//r := new(bytes.Reader)
	buf := make([]byte, 64)
	for {
		//n, err := c.ReadFrom(r)
		c.SetReadBuffer(6) // 设置每次读取的大小
		n, err := c.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read from:", err)
			continue
		}

		if err == io.EOF {
			return
		}
		fmt.Println("read:", string(buf[:n]))
	}
}

func clientHeartbeat(c *net.TCPConn) {
	buf := make([]byte, 256)
	for {
		// SetWriteDeadline 这个有效吗
		//c.SetWriteDeadline(time.Now().Add(10 * time.Second))
		//c.Write()
		c.SetReadDeadline(time.Now().Add(10 * time.Second))
		n, err := c.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("err:", err)
			// TODO heartbeat
			c.Write([]byte("^-^"))
			continue
		}
		if err == io.EOF {
			return
		}
		fmt.Println("read:", buf[:n])
	}

}
