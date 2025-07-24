package parser

import (
	"errors"
	"strings"

	"github.com/helloamj/keyra/internal/engine"
	"github.com/helloamj/keyra/internal/executer"
)

type TcpParser struct {
	Engine executer.EngineInterface
}

func NewTcpParser(engine *engine.Engine) *TcpParser {
	return &TcpParser{Engine: engine}
}

func (p *TcpParser) Parse(reqStr string) (executer.Executer, error) {
	reqStr = strings.TrimSpace(reqStr)
	parts := strings.Fields(reqStr)

	if len(parts) == 0 {
		return nil, errors.New("empty input")
	}

	switch strings.ToUpper(parts[0]) {
	case "SET":
		if len(parts) != 3 {
			return nil, errors.New("SET command must have exactly 3 parts: SET key value")
		}
		return executer.NewSetOp(parts[1], parts[2], p.Engine), nil

	case "GET":
		if len(parts) != 2 {
			return nil, errors.New("GET command must have exactly 2 parts: GET key")
		}
		return executer.NewGetOp(parts[1], p.Engine), nil

	case "DELETE":
		if len(parts) != 2 {
			return nil, errors.New("DELETE command must have exactly 2 parts: DELETE key")
		}
		return executer.NewDeleteOp(parts[1], p.Engine), nil

	default:
		return nil, errors.New("invalid command")
	}
}
