package model

import "time"

type Transaction struct {
	TableNumber       int8          `bson:"tableNumber" json:"tableNumber"`
	Token             string        `bson:"token" json:"token"`
	Size              int8          `bson:"size" json:"size"`
	PromotionName     string        `bson:"promotionName" json:"promotionName"`
	StartPrice        float32       `bson:"startPrice" json:"startPrice"`
	OrderPrice        float32       `bson:"orderPrice" json:"orderPrice"`
	TotalPrice        float32       `bson:"totalPrice" json:"totalPrice"`
	RemainingDuration time.Duration `bson:"remainingDuration" json:"remainingDuration"`
	CreatedAt         time.Time     `bson:"createdAt" json:"createdAt"`
	OrderItems        []OrderItem   `bson:"orderItems" json:"orderItems"`
}
