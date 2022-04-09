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
		time.Sleep(10*time.Second)
		err = sc.Publish(channelName, message[0])
		if err != nil {
			log.Println(err)
		}
		time.Sleep(10*time.Second)

		err = sc.Publish(channelName, message[1])
		if err != nil {
			log.Println(err)
		}
		time.Sleep(15*time.Second)

		err = sc.Publish(channelName, message[2])
		if err != nil {
			log.Println(err)
		}
		time.Sleep(20*time.Second)

		err = sc.Publish(channelName, message[3])
		if err != nil {
			log.Println(err)
	 	}
		time.Sleep(25*time.Second)
		err = sc.Publish(channelName, message[4])
		if err != nil {
			log.Println(err)
	 	} 	
}