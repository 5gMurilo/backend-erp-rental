package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type EpiExits struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	EpiName  string             `json:"name" bson:"epiName"`
	Employee *Employee          `json:"employee" bson:"employee"`
	GaveBy   string             `json:"gaveBy" bson:"gaveBy"`
	Quantity int                `json:"quantity" bson:"quantity"`
	ExitTime primitive.DateTime `json:"exitTime" bson:"exitTime"`
}
