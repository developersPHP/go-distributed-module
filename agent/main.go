package main

import (
	cli2 "agent/app/cli"
	"agent/app/lib"
	"agent/app/options"
	server2 "agent/app/server"
	"fmt"
	"time"
)

var (
	cli       *cli2.UdpClient
	server    *server2.UdpServer
	transport *lib.UdpTransportSender
)

func init() {
	cli = cli2.NewUdpClient(options.Host("127.0.0.1"), options.Port(9981))
	server = server2.NewUdpServer(options.Host("127.0.0.1"), options.Port(9981),
		options.ReceiveBuffLen(1024))
	transport = lib.NewUdpTransportSender(cli, options.SendBufferLen(120), options.TransportPipe(make(chan []byte, 50)),
		options.MaxPoolBufferLen(250), options.BufferFlushInterval(5))
}
func main() {
	//启动udp服务器监听端口
	fmt.Println("开始")
	time.Sleep(time.Second * 3)
	InitUdpServer()
	InitTransportController()
	transport.ReporterToAgent("hello world i am coming")
	time.Sleep(time.Second * 3)
}

func InitUdpServer() {
	go server.UdpSer()
}

func InitTransportController() {
	go transport.TransportController()
}
