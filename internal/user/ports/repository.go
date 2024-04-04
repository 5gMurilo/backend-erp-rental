package ports

import (
	"america-rental-backend/internal/user"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Get(c context.Context, id primitive.ObjectID) (*user.User, error)
	GetAll(c context.Context) (*[]user.User, error)
	Create(c context.Context, data user.User) (*primitive.ObjectID, error)
	Update(c context.Context, data user.User, id primitive.ObjectID) (*user.User, error)
	Delete(c context.Context, id primitive.ObjectID) error
}
