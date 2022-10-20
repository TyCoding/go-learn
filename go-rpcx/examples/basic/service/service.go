package service

import (
	"context"
	"flag"
	"fmt"
	"time"
)

// Addr 定义服务连接地址
var (
	Addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

// MyService 定义一个服务对象
type MyService struct{}

func (t *MyService) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	fmt.Printf("[%v] Client call msg: args=%v, reply: %v\n", time.Now(), *args, *reply)
	return nil
}
