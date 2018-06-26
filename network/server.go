package network

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"time"
)

func Server(addr string) {
	laddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		panic(err)
	}

	lsn, err := net.ListenTCP("tcp4", laddr)
	if err != nil {
		fmt.Println("listen:", err)
		return
	}
	fmt.Println("listen up...")
	for {
		conn, err := lsn.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("new client")
		//go handleConn2(conn)
		//go heartbeat(conn)
		go handle(conn)
	}
}

func handleConn(c *net.TCPConn) {
	// 设置读取时间
	c.SetReadDeadline(time.Now().Add(3 * time.Second)) // 可以通过该方法来定义心跳包
	buf := make([]byte, 256)
	n, err := c.Read(buf)
	if err != nil {
		fmt.Println("read:", err)
		return
	}

	fmt.Println(buf[:n])
	return
}

func handle(fd *net.TCPConn) {
	defer fd.Close()
	for {
		data, err := ioutil.ReadAll(fd)
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(data) == 0 {
			fmt.Println("empty")
			return
		}
		fmt.Println(data)
	}
}

func handleConn2(c *net.TCPConn) {
	for {
		c.Write([]byte("hello world"))
		c.Write([]byte("abcdefghjjikll 123 ll"))
		c.Write([]byte("boll"))
		time.Sleep(2 * time.Second)
	}
}

// 心跳包
func heartbeat(c *net.TCPConn) {
	// TODO
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
