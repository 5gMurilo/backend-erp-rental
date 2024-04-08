package ports

import (
	"america-rental-backend/internal/core/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Get(c context.Context, id primitive.ObjectID) (*domain.User, error)
	GetByEmail(c context.Context, email string) (*domain.User, error)
	GetAll(c context.Context) (*[]domain.User, error)
	Create(c context.Context, data domain.User) (*primitive.ObjectID, error)
	Update(c context.Context, data domain.User, id primitive.ObjectID) (*domain.User, error)
	Delete(c context.Context, id primitive.ObjectID) error
}

type UserService interface {
	Get(c context.Context, id primitive.ObjectID) (*domain.User, error)
	GetAll(c context.Context) (*[]domain.User, error)
	Create(c context.Context, user domain.User) (*domain.User, error)
	Update(c context.Context, user domain.User, id primitive.ObjectID) (*domain.User, error)
	Delete(c context.Context, id primitive.ObjectID) error
}
