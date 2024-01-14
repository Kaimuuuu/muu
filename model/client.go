package model

import "time"

type Client struct {
	TableNumber   int8      `bson:"tableNumber" json:"tableNumber"`
	Size          int8      `bson:"size" json:"size"`
	Token         string    `bson:"token" json:"token"`
	PromotionId   string    `bson:"promotionId" json:"promotionId"`
	PromotionName string    `bson:"promotionName" json:"promotionName"`
	Expire        time.Time `bson:"expire" json:"expire"`
	CreatedAt     time.Time `bson:"createdAt" json:"createdAt"`
	CreatedBy     string    `bson:"creaetdBy" json:"creaetdBy"`
}
