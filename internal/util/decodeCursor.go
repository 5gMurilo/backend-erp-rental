package util

import (
	"america-rental-backend/internal/user"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type decodeTypes interface {
	user.User
}

func DecodeCursor[T decodeTypes](ctx context.Context, cursor *mongo.Cursor) (*[]T, error) {
	var result []T
	for cursor.Next(ctx) {
		var bsonData bson.M
		var generic T

		if err := cursor.Decode(&bsonData); err != nil {
			return nil, err
		}

		bytes, err := bson.Marshal(bsonData)
		if err != nil {
			return nil, err
		}

		err = bson.Unmarshal(bytes, &generic)
		if err != nil {
			return nil, err
		}

		result = append(result, generic)
	}
	return &result, nil
}
