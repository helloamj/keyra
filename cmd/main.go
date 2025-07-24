package main

import (
	"flag"
	"fmt"
	"os"

	keyra "github.com/helloamj/keyra/internal"
	"go.uber.org/fx"
)

func main() {
	port := flag.Int("port", 8080, "Port to run TCP server on")
	numShards := flag.Uint("numShards", 5, "Port to run TCP server on")
	flag.Parse()

	if *port < 1024 || *port > 65535 {
		fmt.Println("Invalid port number. Must be between 1024 and 65535.")
		os.Exit(1)
	}
	app := fx.New(
		keyra.Module(*port, *numShards),
	)
	app.Run()
}
