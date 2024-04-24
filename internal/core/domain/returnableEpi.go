package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnableEpi struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	EpiToBeReturned   Epi                `json:"epiToBeReturned" bson:"epiToBeReturned"`
	Employee          *Employee          `json:"employee,omitempty" bson:"employee,omitempty"`
	Quantity          int                `json:"quantity" bson:"quantity"`
	DateToReturn      *time.Time         `json:"dateToReturn,omitempty" bson:"dateToReturn,omitempty"`
	ReturnedDate      *time.Time         `json:"returnedDate,omitempty" bson:"returnedDate,omitempty"`
	ReturnInformation string             `json:"returnInformation" bson:"returnInformation"`
	GivenDate         time.Time          `json:"givenDate" bson:"givenDate"`
	CreatedBy         *User              `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
}
