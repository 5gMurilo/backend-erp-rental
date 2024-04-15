package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnableEpi struct {
	ID                string              `json:"id" bson:"_id"`
	EpiToBeReturned   *Epi                `json:"epiToBeReturned" bson:"epiToBeReturned"`
	Employee          *Employee           `json:"employee" bson:"employee"`
	Quantity          int                 `json:"quantity" bson:"quantity"`
	DateToReturn      *primitive.DateTime `json:"dateToReturn" bson:"dateToReturn"`
	ReturnedDate      *primitive.DateTime `json:"returnedDate" bson:"returnedDate"`
	ReturnInformation string              `json:"returnInformation" bson:"returnInformation"`
	GivenDate         primitive.DateTime  `json:"givenDate" bson:"givenDate"`
	CreatedBy         *User               `json:"createdBy" bson:"createdBy"`
}
