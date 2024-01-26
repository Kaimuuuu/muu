package model

import "time"

type PromotionMenuItemType int8

const (
	PromotionBuffet = iota
	PromotionALaCarte
)

type Promotion struct {
	Id                 string              `bson:"id" json:"id"`
	Name               string              `bson:"name" json:"name"`
	Description        string              `bson:"description" json:"description"`
	Price              float32             `bson:"price" json:"price"`
	Duration           time.Duration       `bson:"duration" json:"duration"`
	PromotionMenuItems []PromotionMenuItem `bson:"promotionMenuItems" json:"promotionMenuItems"`
	ImagePath          string              `bson:"imagePath" json:"imagePath"`
	CreatedAt          time.Time           `bson:"createdAt" json:"createdAt"`
	CreatedBy          string              `bson:"createdBy" json:"createdBy"`
}

type PromotionMenuItem struct {
	Type       PromotionMenuItemType `bson:"type" json:"type"`
	MenuItemId string                `bson:"menuItemId" json:"menuItemId"`
	Limit      int                   `bson:"limit" json:"limit"`
}
