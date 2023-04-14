package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"testWB-app/internal/entities"

	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster-wb"
	clientID  = "event-store"
)

func Restore(StoreFile string) []entities.Order {
	jsonFile, err := os.OpenFile(StoreFile, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
	}
	orders := make([]entities.Order, 0)

	json.Unmarshal(byteValue, &orders)

	return orders
}

func main() {
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL("nats://0.0.0.0:4222"),
	)
	if err != nil {
		log.Print(err)
		return
	}
	defer sc.Close()

	Orders := Restore("models.json")
	for _, order := range Orders {
		data, err := json.Marshal(order)
		if err != nil {
			log.Println(err)
			continue
		}
		sc.Publish("Order", data)
		log.Println("published message on channel: Order")
	}

	for i := 0; i < 10; i++ {
		sc.Publish("Order", []byte("bad data"))
		log.Println("published message on channel: Order")
	}
}
