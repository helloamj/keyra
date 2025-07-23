package server

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/helloamj/keyra/internal/parser"
)

type TcpServer struct {
	Port   int
	Parser *parser.TcpParser
}

func NewTcpServer(port int, p *parser.TcpParser) *TcpServer {
	return &TcpServer{
		Port:   port,
		Parser: p,
	}
}

func (s *TcpServer) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		return err
	}
	defer listener.Close()
	fmt.Println("🌐 TCP server running on :", s.Port)

	for {
		conn, err := listener.Accept()
		log.Printf("🔗 Client Connected")
		if err != nil {
			continue
		}
		go s.handle(conn)
	}
}

func (s *TcpServer) handle(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		conn.Write([]byte("$ "))
		rawReq, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		resp, err := s.Parser.Parse(rawReq)
		if err != nil {
			conn.Write([]byte("Error: " + err.Error() + "\n"))
		} else {
			conn.Write([]byte(resp))
		}
	}
}
