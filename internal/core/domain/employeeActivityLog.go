package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmployeeActivityLog struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Activity string             `json:"activity" bson:"activity" binding:"required"`
	Employee Employee           `json:"employee" bson:"employee" binding:"required"`
	Actor    string             `json:"actor" bson:"actor" binding:"required"`
	At       primitive.DateTime `json:"didAt" bson:"didAt"`
}
