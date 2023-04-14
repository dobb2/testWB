package storage

import "testWB-app/internal/entities"

type OrdersGetter interface {
	Get(id string) ([]entities.Order, error)
	Create(entities.Order) error
}
