package entities

import (
	"errors"
	"strings"
)

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

func (del *Delivery) Scan(value interface{}) error {
	var valueStr string
	if value == nil {
		*del = Delivery{}
		return nil
	}

	switch value.(type) {
	case string:
		valueStr = value.(string)
	default:
		return errors.New("Incompatible type for Payment")
	}

	valueStr = strings.Trim(valueStr, "()")
	deliveryElemStr := strings.Split(valueStr, ",")
	for i := range deliveryElemStr {
		deliveryElemStr[i] = strings.Trim(deliveryElemStr[i], "\"")
	}

	*del = Delivery{
		Name:    deliveryElemStr[0],
		Phone:   deliveryElemStr[1],
		Zip:     deliveryElemStr[2],
		City:    deliveryElemStr[3],
		Address: deliveryElemStr[4],
		Region:  deliveryElemStr[5],
		Email:   deliveryElemStr[6],
	}
	return nil
}
