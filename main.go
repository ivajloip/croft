package main

import (
	"flag"
	"log"
)

func main() {
	log.Print("Croft is ALIVE")

	port := flag.Int("port", 1790, "Port on which to listen for UDP packets")
	zmq := flag.Bool("useRabbitMQ", false, "Use zmq or rabbitmq")
	flag.Parse()

	var err error
	var publisher Publisher

	log.Print(*zmq)
	if !*zmq {
		publisher, err = connectZmqPublisher()
	} else {
		publisher, err = connectPublisher()
	}

	if err != nil {
		log.Fatalf("Failed to connect publisher: %s", err.Error())
	}

	messages := make(chan interface{})

	go readUDPMessages(*port, messages)
	log.Printf("Started server on port %d", port)
	for msg := range messages {
		err = publisher.Publish(msg)
		if err != nil {
			log.Printf("Failed to publish message %#v: %s", msg, err.Error())
		}
	}
}

func connectZmqPublisher() (Publisher, error) {
	publisher, err := ConnectZmqPublisher()
	if err != nil {
		return nil, err
	}

	err = publisher.Configure()
	if err != nil {
		return nil, err
	}

	return publisher, nil
}

func connectPublisher() (Publisher, error) {
	publisher, err := ConnectRabbitPublisher()
	if err != nil {
		return nil, err
	}

	err = publisher.Configure()
	if err != nil {
		return nil, err
	}

	return publisher, nil
}
