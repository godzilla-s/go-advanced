// 半双工
package network

import (
	"fmt"
	"io"
	"net"
	"time"
)

type connect struct {
	fd     *net.TCPConn
	read   chan string
	write  chan string
	closed chan struct{}
}

func full_duplex_server(addr string) {
	laddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		panic(err)
	}

	lsn, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		fmt.Println("listen:", err)
	}

	fmt.Println("server run...")
	for {
		conn, err := lsn.AcceptTCP()
		if err != nil {
			continue
		}

		c := newConnection(conn)
		go c.run()
		go c.input("server")
	}
}

func newConnection(c *net.TCPConn) *connect {
	return &connect{
		fd:     c,
		read:   make(chan string, 1),
		write:  make(chan string, 1),
		closed: make(chan struct{}),
	}
}

func (c *connect) run() {
	go c.loopRead()
	for {
		select {
		case <-c.closed:
			c.close()
			return
		case s := <-c.read:
			fmt.Println("read <:", s)
		case w := <-c.write:
			c.fd.Write([]byte(w))
		}
	}
}

func (c *connect) loopRead() {
	buf := make([]byte, 512)
	for {
		n, err := c.fd.Read(buf)
		if err != nil && err != io.EOF {
			continue
		}

		if err == io.EOF {
			c.closed <- struct{}{}
			return
		}
		c.read <- string(buf[:n])
		//fmt.Println("read:", buf[:n])
	}
}

func (c *connect) close() {
	c.fd.Close()
}

func (c *connect) input(tag string) {
	for {
		c.write <- fmt.Sprintf("%s: %d", tag, time.Now().Unix())
		time.Sleep(1 * time.Second)
	}
}

func full_duplex_client(addr string) {
	laddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, laddr)
	if err != nil {
		fmt.Println("dial:", err)
		return
	}

	fmt.Println("client run ...")
	c := newConnection(conn)
	go c.input("client")
	c.run()
}
