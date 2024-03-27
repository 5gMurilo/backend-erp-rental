package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id            primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	Name          string             `json:"name" bson:"name, omitempty"`
	Email         string             `json:"email" bson:"email, omitempty"`
	Password      string             `json:"password" bson:"password, omitempty"`
	State         string             `json:"state" bson:"state, omitempty"`
	Role          string             `json:"role" bson:"role, omitempty"`
	UserType      string             `json:"userType" bson:"type, omitempty"`
	MailSignature string             `json:"mail_signature" bson:"mail_signature, omitempty"`
	CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at, omitempty"`
	UpdatedAt     primitive.DateTime `json:"updated_at" bson:"updated_at, omitempty"`
}

func New(name string, email string, password string, state string, role string, userType string, mail_signature string) (User, error) {
	return User{
		Id:            primitive.NewObjectID(),
		Name:          name,
		Email:         email,
		Password:      password,
		State:         state,
		Role:          role,
		UserType:      userType,
		MailSignature: mail_signature,
		CreatedAt:     primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:     primitive.NewDateTimeFromTime(time.Now()),
	}, nil
}

func Update(id primitive.ObjectID, name string, email string, password string, state string, role string, userType string, mail_signature string, createdAt primitive.DateTime) (User, error) {
	return User{
		Id:            id,
		Name:          name,
		Email:         email,
		Password:      password,
		State:         state,
		Role:          role,
		UserType:      userType,
		MailSignature: mail_signature,
		CreatedAt:     createdAt,
		UpdatedAt:     primitive.NewDateTimeFromTime(time.Now()),
	}, nil
}

func Get(id primitive.ObjectID) (User, error) {
	return User{}, nil
}

func GetAll() (users []User, err error) {
	return []User{}, nil
}

func Delete(id primitive.ObjectID) error {
	return nil
}
