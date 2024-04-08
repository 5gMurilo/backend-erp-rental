package domain

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenPayload struct {
	Id       uuid.UUID
	UserId   primitive.ObjectID
	Name     string
	UserType string
}
