package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Name          string             `json:"name" bson:"name"`
	Email         string             `json:"email" bson:"email"`
	Password      string             `json:"password" bson:"password"`
	State         string             `json:"state" bson:"state"`
	Role          string             `json:"role" bson:"role"`
	UserType      string             `json:"userType" bson:"userType"`
	MailSignature string             `json:"mail_signature" bson:"mail_signature, omitempty"`
	CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at, omitempty"`
	UpdatedAt     primitive.DateTime `json:"updated_at" bson:"updated_at, omitempty"`
	UpdatedBy     string             `json:"updatedBy" bson:"updatedBy"`
}
