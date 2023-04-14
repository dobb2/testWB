package backup

import (
	"log"
	"testWB-app/internal/storage"
)

func Restore(cache, db storage.OrdersGetter) {
	ordersDb, err := db.Get("")
	if err != nil {
		log.Println(err)
		return
	}

	if len(ordersDb) == 0 {
		log.Println("Database is empty")
		return
	}

	for _, order := range ordersDb {
		err = cache.Create(order)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("backup done")

}
