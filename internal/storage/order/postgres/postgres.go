package postgres

import (
	"database/sql"
	"log"
	"testWB-app/internal/entities"
	"time"
)

type OrderStorer struct {
	db *sql.DB
}

func New(db *sql.DB) OrderStorer {
	return OrderStorer{db: db}
}

func (o OrderStorer) Get(id string) ([]entities.Order, error) {
	var (
		rows *sql.Rows
		err  error
	)

	if id != "" {
		rows, err = o.db.Query("SELECT \n    "+
			"order_uid, \n    "+
			"track_number,\n    "+
			"entry,\n    "+
			"delivery,\n    "+
			"items,\n "+
			"payment,\n"+
			"locale_chr,\n	"+
			"internal_signature,\n	"+
			"customer_id,\n	"+
			"delivery_service,\n	"+
			"shardkey,\n	"+
			"sm_id integer,\n	"+
			"data_created,\n	"+
			"oof_shard\n	"+
			"FROM Orders\n"+
			"WHERE order_uid = $1", id)
	} else {
		rows, err = o.db.Query("SELECT \n    " +
			"order_uid, \n    " +
			"track_number,\n    " +
			"entry,\n    " +
			"delivery,\n    " +
			"items,\n " +
			"payment,\n" +
			"locale_chr,\n	" +
			"internal_signature,\n	" +
			"customer_id,\n	" +
			"delivery_service,\n	" +
			"shardkey,\n	" +
			"sm_id integer,\n	" +
			"data_created,\n	" +
			"oof_shard\n	" +
			"FROM Orders")
	}
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []entities.Order
	date := new(string)

	for rows.Next() {
		var o entities.Order
		_ = rows.Scan(&o.OrderUID, &o.TrackNumber, &o.Entry,
			&o.Delivery,
			&o.Items,
			&o.Payment,
			&o.Locale, &o.InternalSignature, &o.CustomerID, &o.DeliveryService,
			&o.Shardkey, &o.SmID, date, &o.OofShard,
		)

		layout := "2006-01-02T15:04:05Z"
		o.DateCreated, err = time.Parse(layout, *date)
		if err != nil {
			log.Println(err)
		}
		orders = append(orders, o)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o OrderStorer) Create(order entities.Order) error {
	queryInsertStr := order.InsertSQL()
	_, err := o.db.Exec(queryInsertStr)

	if err != nil {
		return err
	}
	return nil
}
