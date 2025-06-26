package main

import (
	"github.com/codevsk/grpc/client"
	"github.com/codevsk/grpc/server"
)

func main() {
	go server.Run()
	client.Run()
}
