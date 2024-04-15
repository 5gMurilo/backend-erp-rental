package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Epi struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Ca           string             `json:"ca" bson:"ca, omitempty'`
	EpiType      string             `json:"type" bson:"type, omitempty"`
	Stock        int32              `json:"stock" bson:"stock"`
	MinimumStock int32              `json:"minimumStock" bson:"minimumStock"`
	IsReturnable bool               `json:"isReturnable" bson:"isReturnable"`
}
