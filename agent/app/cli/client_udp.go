package cli

import (
	"agent/app/options"
	"fmt"
	"net"
)

type UdpClient struct {
	BaseClientInfo options.Options
}

func NewUdpClient(p ...options.Option) *UdpClient {
	obj := new(UdpClient)
	res := options.ApplyOptions(p...)
	obj.BaseClientInfo = res
	return obj
}

func (this *UdpClient) UdpCli(data []byte) {

	ip := net.ParseIP(this.BaseClientInfo.Host)

	srcAdd := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAdd := &net.UDPAddr{IP: ip, Port: this.BaseClientInfo.Port}

	conn, err := net.DialUDP("udp", srcAdd, dstAdd)
	if err != nil {
		fmt.Printf("udp client fail")
	}
	defer conn.Close()
	conn.Write(data)
}
