package entities

import (
	"errors"
	"strconv"
	"strings"
)

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

func (pay *Payment) Scan(value interface{}) error {
	var valueStr string
	if value == nil {
		*pay = Payment{}
		return nil
	}

	switch value.(type) {
	case string:
		valueStr = value.(string)
	default:
		return errors.New("Incompatible type for Payment")
	}

	valueStr = strings.Trim(valueStr, "()")
	paymentElemStr := strings.Split(valueStr, ",")
	for i := range paymentElemStr {
		paymentElemStr[i] = strings.Trim(paymentElemStr[i], "\"")
	}
	Amount, _ := strconv.Atoi(paymentElemStr[4])
	PaymentDt, _ := strconv.Atoi(paymentElemStr[5])
	DeliveryCost, _ := strconv.Atoi(paymentElemStr[7])
	GoodsTotal, _ := strconv.Atoi(paymentElemStr[8])
	CustomFee, _ := strconv.Atoi(paymentElemStr[9])

	*pay = Payment{
		Transaction:  paymentElemStr[0],
		RequestID:    paymentElemStr[1],
		Currency:     paymentElemStr[2],
		Provider:     paymentElemStr[3],
		Amount:       Amount,
		PaymentDt:    PaymentDt,
		Bank:         paymentElemStr[6],
		DeliveryCost: DeliveryCost,
		GoodsTotal:   GoodsTotal,
		CustomFee:    CustomFee,
	}

	return nil
}

/*
func (pay Payment) Value() (driver.Value, error) {
	payStr := "('" + pay.Transaction + "', '" + pay.RequestID + "', '" +
		pay.Currency + "', '" + pay.Provider + "', " +
		strconv.FormatInt(int64(pay.Amount), 10) + ", " +
		strconv.FormatInt(int64(pay.PaymentDt), 10) + ", '" + pay.Bank + "', " +
		strconv.FormatInt(int64(pay.DeliveryCost), 10) + ", " +
		strconv.FormatInt(int64(pay.GoodsTotal), 10) + ", " +
		strconv.FormatInt(int64(pay.CustomFee), 10) + ")::Payment"

	return payStr, nil
}

*/
