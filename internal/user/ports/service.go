package ports

import (
	"america-rental-backend/internal/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Get(c *gin.Context, id primitive.ObjectID) (*user.User, error)
	GetAll(c *gin.Context) ([]user.User, error)
	Create(c *gin.Context, user user.User) (*user.User, error)
	Update(c *gin.Context, user user.User, id primitive.ObjectID) (*user.User, error)
	Delete(c *gin.Context, id primitive.ObjectID) error
}
