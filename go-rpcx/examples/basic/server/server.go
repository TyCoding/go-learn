package main

import (
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"go-rpcx/examples/basic/service"
)

func main() {
	flag.Parse()
	fmt.Printf("Server started... address: [%v]\n", *service.Addr)

	s := server.NewServer()
	s.Register(new(service.MyService), "")
	err := s.Serve("tcp", *service.Addr)
	if err != nil {
		panic(err)
	}
}
