package lib

import (
	"agent/app/cli"
	"agent/app/options"
	"fmt"
	"time"
)

type UdpTransportSender struct {
	BaseInfo options.UdpTransportSenders
	UdpCli   *cli.UdpClient
}

func NewUdpTransportSender(udpCli *cli.UdpClient, option ...options.UdpTransportSender) *UdpTransportSender {

	//TODO YOU CAN JUDGE PARAM AND ASSIGNMENT DEFAULT VALUE
	obj := new(UdpTransportSender)
	opt := options.ApplyUdpTransportSenders(option...)
	obj.BaseInfo = opt
	obj.UdpCli = udpCli
	return obj
}

func (this *UdpTransportSender) Append() {

}

func (this *UdpTransportSender) Flush() {

}

//TODO YOU CAN DO SOMETHING COMBINE WITH THRIFT
func (this *UdpTransportSender) CalcSizeOfSerializedThrift() {

}

//向bufferPool发送数据
func (this *UdpTransportSender) ReporterToAgent(data string) {
	for {
		select {
		case this.BaseInfo.TransportPipe <- []byte(data):
			//len1 := <-reprot.quent
			//fmt.Println("数据长度为:", len([]byte(len1)))
		default:
		}
	}

}

//
func (this *UdpTransportSender) TransportBufferPoolHandle() {

}

//发送控制器
func (this *UdpTransportSender) TransportController() {
	//TODO 可以抽出来到flush方法里面，利用thrift协议转换再传
	flush := func() {
		if len(this.BaseInfo.BufferPool) >= this.BaseInfo.MaxPoolBufferLen {
			this.UdpCli.UdpCli(this.BaseInfo.BufferPool)
			fmt.Println("输出的bufferlen", len(this.BaseInfo.BufferPool))
			//清空buffer
			this.BaseInfo.BufferPool = []byte{}
		}
	}
	timer := time.NewTicker(this.BaseInfo.BufferFlushInterval)
	for {
		select {
		case <-timer.C:
			fmt.Println("时间到了")
			flush()
		case item := <-this.BaseInfo.TransportPipe:
			//TODO 可以抽出来到Append方法里面，通过THRIFT协议转换再追加
			this.BaseInfo.BufferPool = append(this.BaseInfo.BufferPool, item...)
			flush()
		}
	}
}

func (this *UdpTransportSender) Close() {

}

func (this *UdpTransportSender) ResetBuffer() {

}
