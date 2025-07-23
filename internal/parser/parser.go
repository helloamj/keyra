package parser

import "net"

type Parser interface {
	Parse(conn net.Conn, reqStr string) ([]byte, error)
}
