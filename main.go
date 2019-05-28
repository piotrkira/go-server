package main

import (
	"go-server/pkg/server"
)

func main() {
	s := server.NewServer()
	s.Start()
}
