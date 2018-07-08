// 实现简单的聊天功能
package chat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

type ChatServer struct {
	self     string
	laddr    *net.TCPAddr
	connpool map[string]*conn
	mux      sync.Mutex
	//readConsole chan string // 从命令行读取数据
	exit      chan struct{}
	closeConn chan string // 关闭的链接
}

type conn struct {
	fd *net.TCPConn
	id string
	//read  <-chan message
	write   chan message
	close   chan struct{}
	closing chan string //
}

func New(addr, id string) *ChatServer {
	laddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		panic(err)
	}

	chat := new(ChatServer)
	chat.self = id
	chat.laddr = laddr
	chat.connpool = make(map[string]*conn)
	chat.exit = make(chan struct{})
	chat.closeConn = make(chan string, 5)

	go chat.listenUp()
	return chat
}

func (chat *ChatServer) listenUp() {
	lsn, err := net.ListenTCP("tcp", chat.laddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("listen up...")
	for {
		fd, err := lsn.AcceptTCP()
		if err != nil {
			continue
		}

		go chat.newConnection(fd, false)
	}
}

var (
	conFlag bool
)

func (chat *ChatServer) Start() {
	console := NewConsole()
	go console.Start()

	for {
		select {
		case <-chat.exit:
			return
		case <-console.exit:
			return
		case closeID := <-chat.closeConn:
			// 删除链接
			fmt.Println("closing conn:", closeID)
			chat.delete(closeID)
		case read := <-console.Read():
			if console.command.connect {
				chat.Dial(read)
				console.command.connect = false
				break
			}

			if read == "peers" {
				fmt.Println(chat.clients())
				break
			}
			chat.writeMsg(read)
		}
	}
}

func (chat *ChatServer) stop() {
	close(chat.exit)
}

func (chat *ChatServer) newConnection(fd *net.TCPConn, dial bool) {
	conn := new(conn)
	conn.fd = fd
	msg := make(chan message)
	//conn.read = msg
	conn.write = msg
	conn.close = make(chan struct{})
	conn.closing = chat.closeConn

	var err error

	if dial {
		err = conn.protoHandshake(chat.self)
	} else {
		err = conn.ackHandshake(chat.self)
	}
	if err != nil {
		fmt.Println("handshale fail:", err)
		return
	}

	chat.add(conn)
	go conn.readMsg()
	go conn.writeMsg()
}

func (chat *ChatServer) add(c *conn) {
	chat.mux.Lock()
	chat.connpool[c.id] = c
	chat.mux.Unlock()
}

func (chat *ChatServer) delete(id string) {
	chat.mux.Lock()
	c := chat.connpool[id]
	if c == nil {
		return
	}
	c.fd.Close()
	close(c.close)
	close(c.write)
	delete(chat.connpool, id)
	chat.mux.Unlock()
}

func (chat *ChatServer) writeMsg(read string) {
	msg := message{From: chat.self, Timestamp: time.Now(), Data: []byte(read)}
	for _, conn := range chat.connpool {
		fmt.Println("write to >>", conn.id)
		conn.write <- msg
	}
}

func (chat *ChatServer) clients() []string {
	var clis []string
	for _, c := range chat.connpool {
		clis = append(clis, fmt.Sprintf("ID:%v, Addr: %v", c.id, c.fd.RemoteAddr()))
	}
	return clis
}

const (
	ReqHandShake = iota + 1
	RspHandShake
)

// 发送握手协议，确保链接正常
func (c *conn) protoHandshake(id string) error {
	buf := new(bytes.Buffer)
	buf.WriteByte(ReqHandShake)
	buf.WriteString(id)
	c.fd.Write(buf.Bytes())

	c.fd.SetReadDeadline(time.Now().Add(10 * time.Second))
	data := make([]byte, 256)
	n, err := c.fd.Read(data)
	if err != nil {
		return err
	}

	if data[0] != RspHandShake {
		fmt.Println("not match proto procotol:", data[0], RspHandShake)
		return errors.New("not match proto procotol")
	}

	c.id = string(data[1:n])
	return nil
}

// 确认握手协议，返回自己的节点ID
func (c *conn) ackHandshake(id string) error {
	c.fd.SetReadDeadline(time.Now().Add(10 * time.Second))
	data := make([]byte, 256)
	n, err := c.fd.Read(data)
	if err != nil {
		return err
	}

	if data[0] != ReqHandShake {
		fmt.Println("not match proto procotol:", data[0], ReqHandShake)
		return errors.New("not match proto procotol")
	}

	c.id = string(data[1:n])

	buf := new(bytes.Buffer)
	buf.WriteByte(RspHandShake)
	buf.WriteString(id)
	c.fd.Write(buf.Bytes())
	return nil
}

func (c *conn) readMsg() {
	buf := make([]byte, 1028)
	for {
		c.fd.SetReadDeadline(time.Now().Add(5 * time.Second))
		n, err := c.fd.Read(buf)
		// 屏蔽超时读取
		if err != nil && strings.Contains(err.Error(), "i/o timeout") {
			continue
		}

		if err != nil && err != io.EOF {
			fmt.Println("readerr:", err)
			continue
		}
		if err == io.EOF {
			fmt.Println("readerr:", err)
			c.close <- struct{}{}
			return
		}
		msg, err := decodeMsg(buf[:n])
		if err != nil {
			fmt.Println("error decode:", err)
			continue
		}
		fmt.Println(msg)
	}
}

func (c *conn) writeMsg() {
	for {
		select {
		case msg := <-c.write:
			// fmt.Println("write:", msg)
			data, err := encodeMsg(msg)
			if err != nil {
				fmt.Println("ecode err:", err)
				break
			}
			_, err = c.fd.Write(data)
			if err != nil {
				fmt.Println("write error:", err)
			}
		case <-c.close:
			c.closing <- c.id
			// fmt.Println("emit signal to close conn")
			return
		}
	}
}

type message struct {
	From      string
	Timestamp time.Time
	Data      []byte
}

func (m message) String() string {
	return fmt.Sprintf("<%v> from <%s>: %v", m.Timestamp, m.From, string(m.Data))
}

func decodeMsg(buf []byte) (message, error) {
	var msg message
	err := json.Unmarshal(buf, &msg)
	return msg, err
}

func encodeMsg(msg message) ([]byte, error) {
	return json.Marshal(&msg)
}
