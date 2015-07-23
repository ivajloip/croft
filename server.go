package main

import (
	"fmt"
	"github.com/thethingsnetwork/croft/lora"
	"log"
	"net"
)

func StartUDPServer(port int) {
	ServerAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("0.0.0.0:%d", port))
	CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	lc := lora.NewConn(ServerConn)

	for {
		msg, err := lc.ReadMessage()
		if err != nil {
			continue
		}
		log.Printf("Parsed Message: %#v", msg)
	}
}