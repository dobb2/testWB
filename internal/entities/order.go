package entities

import (
	"strconv"
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             ArrayItem `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

func (ord Order) InsertSQL() string {
	deliveryStr := "('" +
		ord.Delivery.Name + "', '" +
		ord.Delivery.Phone + "', '" +
		ord.Delivery.Zip + "', '" +
		ord.Delivery.City + "', '" +
		ord.Delivery.Address + "', '" +
		ord.Delivery.Region + "', '" +
		ord.Delivery.Email + "')::Delivery"

	paymentStr := "('" + ord.Payment.Transaction + "', '" + ord.Payment.RequestID + "', '" +
		ord.Payment.Currency + "', '" + ord.Payment.Provider + "', " +
		strconv.FormatInt(int64(ord.Payment.Amount), 10) + ", " +
		strconv.FormatInt(int64(ord.Payment.PaymentDt), 10) + ", '" + ord.Payment.Bank + "', " +
		strconv.FormatInt(int64(ord.Payment.DeliveryCost), 10) + ", " +
		strconv.FormatInt(int64(ord.Payment.GoodsTotal), 10) + ", " +
		strconv.FormatInt(int64(ord.Payment.CustomFee), 10) + ")::Payment"

	itemsStr := "ARRAY["
	for i, item := range ord.Items {
		itemElemStr := "(" + strconv.FormatInt(int64(item.ChrtID), 10) + ", '" +
			item.TrackNumber + "', " +
			strconv.FormatInt(int64(item.Price), 10) + ", '" +
			item.Rid + "', '" +
			item.Name + "', " +
			strconv.FormatInt(int64(item.Sale), 10) + ", '" +
			item.Size + "', " +
			strconv.FormatInt(int64(item.TotalPrice), 10) + ", " +
			strconv.FormatInt(int64(item.NmID), 10) + ", '" +
			item.Brand + "', " +
			strconv.FormatInt(int64(item.Status), 10) + ")::Items"
		itemsStr += itemElemStr
		if i+1 != len(ord.Items) {
			itemsStr += ","
		}
	}
	itemsStr += "]"

	orderStr := "'" + ord.OrderUID + "', '" +
		ord.TrackNumber + "', '" +
		ord.Entry + "', " +
		deliveryStr + ", " +
		paymentStr + ", " +
		itemsStr + ", '" +
		ord.Locale + "', '" +
		ord.InternalSignature + "', '" +
		ord.CustomerID + "', '" +
		ord.DeliveryService + "', '" +
		ord.Shardkey + "', " +
		strconv.FormatInt(int64(ord.SmID), 10) + ", '" +
		string(ord.DateCreated.Format("2006-01-02T15:04:05Z")) + "', '" +
		ord.OofShard + "'"

	insertOrder := "INSERT INTO Orders VALUES (" + orderStr + ");"

	return insertOrder
}
