package model

import (
	"time"
)

type OrderStatus int8

const (
	Pending OrderStatus = iota
	Success
	Decline
)

type Order struct {
	Id          string      `bson:"id" json:"id"`
	TableNumber int8        `bson:"tableNumber" json:"tableNumber"`
	OrderItems  []OrderItem `bson:"orderItems" json:"orderItems"`
	Status      OrderStatus `bson:"status" json:"status"`
	CreatedAt   time.Time   `bson:"createdAt" json:"createdAt"`
	OrderBy     string      `bson:"orderBy" json:"orderBy"`
}

type OrderItem struct {
	MenuItemId string  `bson:"menuItemId" json:"menuItemId"`
	Name       string  `bson:"name" json:"name"`
	Quantity   int8    `bson:"quantity" json:"quantity"`
	OutOfStock bool    `bson:"outOfStock" json:"outOfStock"`
	Price      float32 `bson:"price" json:"price"`
}
