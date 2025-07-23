package keyra

import (
	"github.com/helloamj/keyra/internal/parser"
	"github.com/helloamj/keyra/internal/server"
	"go.uber.org/fx"
)

func Module(port int) fx.Option {
	return fx.Options(
		fx.Provide(
			func() int { return port },
			parser.NewTcpParser,
			server.NewTcpServer,
		),
		fx.Invoke(func(s *server.TcpServer) {
			go s.Start()
		}),
	)
}
