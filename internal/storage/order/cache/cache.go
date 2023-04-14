package cache

import (
	"errors"
	"sync"
	"testWB-app/internal/entities"
)

type Orders struct {
	Orders map[string]entities.Order
	mx     sync.RWMutex
}

func Create() Orders {
	return Orders{
		Orders: map[string]entities.Order{},
	}
}

func (o Orders) Create(value entities.Order) error {
	o.mx.Lock()
	defer o.mx.Unlock()
	key := value.OrderUID
	if _, ok := o.Orders[key]; ok {
		return errors.New("The order with this id already exists")
	}
	o.Orders[key] = value
	return nil
}

func (o Orders) Get(id string) ([]entities.Order, error) {
	o.mx.RLock()
	defer o.mx.RUnlock()

	if len(o.Orders) == 0 {
		return nil, errors.New("Cache is empty")
	}

	var orders []entities.Order
	if id != "" {
		if o, ok := o.Orders[id]; ok {
			orders = append(orders, o)
			return orders, nil
		} else {
			return nil, errors.New("unknown id order")
		}
	}

	for _, o := range o.Orders {
		orders = append(orders, o)
	}
	return orders, nil
}
