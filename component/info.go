package component

import (
	"go-boilerplate/domain/info"
	"go-boilerplate/domain/info/delivery"
)

func NewInfoDelivery() info.Delivery {
	return &delivery.Delivery{}
}
