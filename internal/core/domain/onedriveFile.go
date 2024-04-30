package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OnedriveFile struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Filename    string             `json:"filename" bson:"filename"`
	FileUrl     string             `json:"fileUrl" bson:"fileUrl"`
	DriveItemId string             `json:"driveItemId,omitempty" bson:"driveItemId,omitempty"`
	Employee    string             `json:"employee" bson:"employee"`
	Type        string             `json:"type" bson:"type"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
	UpdatedBy   string             `json:"updatedBy" bson:"updatedBy"`
}
