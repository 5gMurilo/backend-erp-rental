package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserType string

const (
	Admin      UserType = "admin"
	Supervisor UserType = "supervisor"
)

type User struct {
	Id            primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Name          string             `json:"name" bson:"name, omitempty"`
	Email         string             `json:"email" bson:"email, omitempty"`
	Password      string             `json:"password" bson:"password, omitempty"`
	State         string             `json:"state" bson:"state, omitempty"`
	Role          string             `json:"role" bson:"role, omitempty"`
	UserType      UserType           `json:"userType" bson:"type, omitempty"`
	MailSignature string             `json:"mail_signature" bson:"mail_signature, omitempty"`
	CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at, omitempty"`
	UpdatedAt     primitive.DateTime `json:"updated_at" bson:"updated_at, omitempty"`
}
