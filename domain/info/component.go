package info

import (
	"go-boilerplate/domain/info/delivery"
)

func NewInfoDelivery() Delivery {
	return &delivery.Delivery{}
}
