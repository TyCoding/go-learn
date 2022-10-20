package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"go-rpcx/examples/basic/service"
	"time"
)

func main() {
	flag.Parse()

	// 定义服务发现的方式，Peer2PeerDiscovery 点对点 —— 客户端直连服务器获取服务地址
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*service.Addr, "")

	// FailMode：客户端调用失败的重试方式； SelectMode：客户端在多服务端情况下如选择服务端
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON
	xclient := client.NewXClient("MyService", client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	// 请求参数
	args := service.Args{
		A: 10,
		B: 20,
	}
	fmt.Printf("Client started... address: [%v]\n", *service.Addr)

	// 发送请求
	for {
		// 响应参数
		reply := &service.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Server listen msg: args=%v, reply: %v\n", args, *reply)
		time.Sleep(time.Second * 5)
	}
}
