package main

import (
	"encoding/json"
	"log"

	zmq "github.com/pebbe/zmq4"
)

type ZmqPublisher struct {
	socket *zmq.Socket
}

func ConnectZmqPublisher() (Publisher, error) {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.DEALER)
	//defer context.Term()
	//defer socket.Close()

	publisher := &ZmqPublisher{socket}

	return publisher, nil
}

func (p *ZmqPublisher) Configure() error {
	uri := "tcp://127.0.0.1:5556"
	err := p.socket.Connect(uri)

	if err != nil {
		log.Printf("Failed to connect to %s, because %s", uri, err.Error())
		return err
	}
	log.Printf("Connected to %s", uri)

	return nil
}

func (p *ZmqPublisher) Publish(data interface{}) error {
	msg, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal data: %s", err.Error())
		return err
	}

	_, err = p.socket.SendBytes(msg, 0)
	if err != nil {
		log.Printf("Failed to send data: %s", err.Error())
		return err
	}

	return nil
}
