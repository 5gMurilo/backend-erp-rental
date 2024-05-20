package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ContractData struct {
	Id             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ContractName   string             `json:"contractName" bson:"contractName" binding:"required"`
	Description    string             `json:"description" bson:"description" binding:"required"`
	File           []byte             `json:"file" bson:"file" binding:"required"`
	ContractValue  float64            `bson:"contractValue" json:"contractValue" binding:"required"`
	SalaryDiscount float64            `bson:"salaryDiscount,omitempty" json:"salaryDiscount,omitempty"`
	Course         string             `json:"course,omitempty" bson:"course,omitempty"`
}

type Contract struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Contract ContractData       `json:"contract" bson:"contract" binding:"required"`
	Employee Employee           `json:"employee" bson:"employee" binding:"required"`
	Status   ContractStatus     `json:"status" bson:"status" binding:"required"`
}

type ContractStatus string

const (
	WAITING_SIGNATURE ContractStatus = "Aguardando assinatura"
	ACTIVE            ContractStatus = "Ativo"
	INACTIVE          ContractStatus = "Inativo"
)
