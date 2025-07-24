package server

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/helloamj/keyra/internal/parser"
)

type TcpServer struct {
	port   int
	parser *parser.TcpParser
}

func NewTcpServer(port int, parser *parser.TcpParser) *TcpServer {
	return &TcpServer{
		port:   port,
		parser: parser,
	}
}

func (s *TcpServer) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	defer listener.Close()
	fmt.Println("🌐 TCP server running on :", s.port)

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
	write := func(msg string) { conn.Write([]byte(msg)) }
	write("[+] Successfully Connected to Keyra\r\n")
	for {
		req, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		exec, err := s.parser.Parse(req)
		if err != nil {
			write("Bad Request Error: " + err.Error() + "\r\n")
			continue
		}
		resp, err := exec.Execute()
		if err != nil {
			write("Unexpected Error: " + err.Error() + "\r\n")
			continue
		}
		write(*resp + "\r\n")
	}
}
