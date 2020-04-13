package client

import (
	"fmt"
	"net"

	"log"
)

type Client struct {
	serverAddr string
	conn       net.Conn
}

func NewClient(serverAddr string) *Client {
	return &Client{
		serverAddr: serverAddr,
	}
}

// 第一步：连接服务器
func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", c.serverAddr)
	if err != nil {
		return err
	}

	c.conn = conn

	log.Println("connect server success")
	return nil
}

// 第二步：输入消息，发送消息，接收消息
func (c *Client) Run() {
	var message string           // 消息
	var buf = make([]byte, 1024) // 缓冲区

	// 死循环
	for {
		// 值传递，引用传递
		_, inputErr := fmt.Scanln(&message)
		if inputErr != nil {
			log.Printf("input with err: %v", inputErr)
			break
		}

		// 类型转换
		_, writeErr := c.conn.Write([]byte(message))
		if writeErr != nil {
			log.Printf("write with err: %v", writeErr)
			break
		}

		_, readErr := c.conn.Read(buf)
		if readErr != nil {
			log.Printf("read with err: %v", readErr)
			break
		}

		log.Printf("Recv: %s\n", buf)
	}
}
