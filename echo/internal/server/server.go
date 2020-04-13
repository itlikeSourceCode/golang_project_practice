package server

import (
	"log" // 日志
	"net" // 网络库
)

type Server struct {
	addr     string       // 服务器地址
	listener net.Listener // 监听器
}

// 构造器
// server := &Server{addr: "XXXX"}
func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

// 第一步：创建监听器，监听连接
func (s *Server) Listen() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	s.listener = listener

	log.Printf("server listen on %v", s.addr)

	return nil
}

// 第二步：运行
func (s *Server) Run() {
	// 死循环
	for {
		conn, err := s.listener.Accept()

		if err != nil {
			log.Printf("Accept with err: %v", err)
			continue
		}

		log.Printf("new connection from %v", conn.RemoteAddr())

		s.handleEcho(conn)
	}
}

// 第三步：处理一个连接的消息
func (s *Server) handleEcho(conn net.Conn) {
	// 并发，开启一个Goroutine去运行这个函数
	go func() {
		// 定义了一个缓冲区
		var buf = make([]byte, 1024)

		// 死循环
		for {
			// 如果某个返回值用不到，用_屏蔽掉
			_, readErr := conn.Read(buf)
			if readErr != nil {
				log.Printf("Read with err: %v", readErr)
				break
			}

			log.Printf("Recv message: %s", buf)

			_, writeErr := conn.Write(buf)
			if writeErr != nil {
				log.Printf("Write with err: %v", writeErr)
				break
			}
		}
	}()
}
