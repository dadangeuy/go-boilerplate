package info

import "go-boilerplate/domain/info/delivery"

func NewInfoDelivery() delivery.Delivery {
	return &delivery.DefaultDelivery{}
}
