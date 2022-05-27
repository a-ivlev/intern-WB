package main

import (
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

var (
	clusterID   = "test-cluster"
	clientID    = "client-id"
	channelName = "orders"
)

func main() {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalln(err)
	}

	for _, msg := range message {
		time.Sleep(10 * time.Second)
		err = sc.Publish(channelName, msg)
		if err != nil {
			log.Println(err)
		}
	}
}
