package model

import "time"

type MenuItem struct {
	Id          string    `bson:"id" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	Catagory    string    `bson:"catagory" json:"catagory"`
	Price       float32   `bson:"price" json:"price"`
	OutOfStock  bool      `bson:"outOfStock" json:"outOfStock"`
	ImagePath   string    `bson:"imagePath" json:"imagePath"`
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	CreatedBy   string    `bson:"creaetdBy" json:"creaetdBy"`
}
