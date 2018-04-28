package options

import "time"

type UdpTransportSender func(*UdpTransportSenders)

type UdpTransportSenders struct {
	SendBufferLen       int
	TransportPipe       chan []byte
	MaxPoolBufferLen    int
	BufferPool          []byte
	BufferFlushInterval time.Duration
}

func SendBufferLen(c int) UdpTransportSender {
	return func(senders *UdpTransportSenders) {
		senders.SendBufferLen = c
	}
}

func TransportPipe(c chan []byte) UdpTransportSender {
	return func(senders *UdpTransportSenders) {
		senders.TransportPipe = c
	}
}

func MaxPoolBufferLen(c int) UdpTransportSender {
	return func(senders *UdpTransportSenders) {
		senders.MaxPoolBufferLen = c
	}
}
func BufferFlushInterval(c time.Duration) UdpTransportSender {
	return func(senders *UdpTransportSenders) {
		senders.BufferFlushInterval = c
	}
}

//申请结构体值
func ApplyUdpTransportSenders(options ...UdpTransportSender) UdpTransportSenders {
	opts := UdpTransportSenders{}
	for _, option := range options {
		option(&opts)
	}
	return opts
}
