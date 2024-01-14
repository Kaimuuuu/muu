package model

import "time"

type EmployeeRole int8

const (
	Admin EmployeeRole = iota
	Chef
	Waiter
)

type Employee struct {
	Id        string       `bson:"id" json:"id"`
	Name      string       `bson:"name" json:"name"`
	Password  string       `bson:"password" json:"password"`
	Age       int8         `bson:"age" json:"age"`
	Role      EmployeeRole `bson:"role" json:"role"`
	Email     string       `bson:"email" json:"email"`
	ImagePath string       `bson:"imagePath" json:"imagePath"`
	CreatedAt time.Time    `bson:"createdAt" json:"createdAt"`
	CreatedBy string       `bson:"creaetdBy" json:"creaetdBy"`
}
