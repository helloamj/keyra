package parser

type TcpParser struct{}

func NewTcpParser() *TcpParser {
	return &TcpParser{}
}

func (p *TcpParser) Parse(reqStr string) (string, error) {
	return reqStr, nil
}
