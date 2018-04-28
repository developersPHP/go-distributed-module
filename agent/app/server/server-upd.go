package server

import (
	"agent/app/options"
	"fmt"
	"net"
)

type UdpServer struct {
	BaseServerInfo options.Options
}

func NewUdpServer(p ...options.Option) *UdpServer {

	obj := new(UdpServer)
	res := options.ApplyOptions(p...)
	obj.BaseServerInfo = res
	return obj

}

func (this *UdpServer) UdpSer() {

	ip := net.ParseIP(this.BaseServerInfo.Host)

	receiveBuffLen := this.BaseServerInfo.ReceiveBuffLen

	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: ip, Port: this.BaseServerInfo.Port})
	if err != nil {
		fmt.Printf("udp监听出错")
	}

	fmt.Printf("Local:<%s>\n", listener.LocalAddr().String())
	data := make([]byte, receiveBuffLen)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read : %s", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])

		//这里可以向发送的upd发送数据
		//_, err = listener.WriteToUDP([]byte("i have recived"), remoteAddr)
		//if err != nil {
		//	fmt.Printf("发送数据出错%s\n", err)
		//}
	}
}
