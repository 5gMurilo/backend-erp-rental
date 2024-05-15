package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EpiExits struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	EpiName  string             `json:"epiName" bson:"epiName" binding:"required"`
	Employee Employee           `json:"employee,omitempty" bson:"employee,omitempty" binding:"required"`
	GaveBy   string             `json:"gaveBy" bson:"gaveBy"`
	Quantity int                `json:"quantity" bson:"quantity" binding:"required"`
	ExitTime primitive.DateTime `json:"exitTime,omitempty" bson:"exitTime,omitempty"`
}
