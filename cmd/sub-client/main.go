package main

import (
	"encoding/json"
	"github.com/caarlos0/env/v7"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"testWB-app/internal/backup"
	"testWB-app/internal/entities"
	"testWB-app/internal/storage/order/postgres"
	"time"

	stan "github.com/nats-io/stan.go"

	"log"
	"net/http"
	"testWB-app/internal/config"
	"testWB-app/internal/driver"
	"testWB-app/internal/handlers"
	"testWB-app/internal/storage/order/cache"
)

func main() {
	var conf config.Config
	err := env.Parse(&conf)
	if err != nil {
		log.Println(err)
	}

	// connect to postgres
	db, err := driver.ConnectToPostgres(conf)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	datastore := cache.Create()
	handler := handlers.New(datastore)
	postgrestore := postgres.New(db)

	// Data retrieval in the repository
	go backup.Restore(datastore, postgrestore)

	// connect to nats-streaming
	sc, err := stan.Connect(conf.NatsClusterID, conf.NatsClientID)
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	aw, _ := time.ParseDuration("60s")
	sc.Subscribe("Order", func(msg *stan.Msg) {
		msg.Ack() // message confirmation
		order := entities.Order{}
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Println("the data does not match the expected data type in json format")
			return
		}

		err = postgrestore.Create(order)
		if err != nil {
			log.Println(err)
		}

		err = datastore.Create(order)
		if err != nil {
			log.Println(err)
		}

		log.Printf("Subscribed message from clientID - %s for OrderID: %+v\n", conf.NatsClientID, order.OrderUID)
	}, stan.DeliverAllAvailable(),
		stan.DurableName(conf.NatsDurableID),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Route("/getOrder", func(r chi.Router) {
		r.Get("/{id}", handler.GetOrderByID)
	})

	log.Fatal(http.ListenAndServe(conf.Address, r))
}
