package lora

import (
	"encoding/json"
	"log"
	"net"
)

const (
	PUSH_DATA = iota
	PUSH_ACK  = iota
	PULL_DATA = iota
	PULL_ACK  = iota
	PULL_RESP = iota
)

var buf = make([]byte, 2048)

type Message struct {
	ProtocolVersion int
	Token           []byte
	Identifier      int
	Payload         *json.RawMessage
	GatewayEUI      int64
	Address         *net.UDPAddr
}

type Conn struct {
	Raw *net.UDPConn
}

func NewConn(r *net.UDPConn) *Conn {
	return &Conn{r}
}

func (c *Conn) ReadMessage() (*Message, error) {
	n, addr, err := c.Raw.ReadFromUDP(buf)
	if err != nil {
		log.Print("Error: ", err)
		return nil, err
	}
	log.Print("Received ", string(buf[0:n]), " from ", addr)
	msg := &Message{
		Address:    addr,
		Identifier: int(buf[3]),
	}
	return msg, nil
}