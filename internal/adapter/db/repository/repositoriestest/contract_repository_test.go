package repositoriestest

import (
	"america-rental-backend/internal/adapter/db/repository"
	"america-rental-backend/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContractCreation(t *testing.T) {
	repo := repository.NewContractRepository(worker)
	newContract := domain.Contract{
		ContractName:  "Contrato",
		Description:   "teste de contrato",
		File:          []byte("asdalsdjklajdlkadna"),
		ContractValue: 3000.00,
		Course:        "Design",
	}

	rst, err := repo.CreateContract(newContract)

	contract, err := repo.GetContractById(rst.InsertedID.(primitive.ObjectID))
	if err != nil {
		return
	}

	assert.Nil(t, err)
	assert.NotNil(t, rst.InsertedID)
	assert.NotNil(t, contract.ContractName)
}

func TestContractUpdate(t *testing.T) {
	repo := repository.NewContractRepository(worker)
	id, err := primitive.ObjectIDFromHex("66451b76cad5fc0d106aa11f")
	if err != nil {
		log.Fatalln(err)
		return
	}
	newContractData := domain.Contract{
		Id:             id,
		ContractName:   "Contrato de faculdade",
		Description:    "Contrato da faculdade do fulano",
		File:           []byte("teste de update de contrato"),
		ContractValue:  3500.46,
		SalaryDiscount: 0,
		Course:         "Universidade",
	}

	contract, err := repo.UpdateContract(newContractData, id)
	if err != nil {
		log.Fatalln(err)
		return
	}

	assert.Nil(t, err)
	assert.NotNil(t, contract)
	assert.Equal(t, newContractData.ContractName, contract.ContractName)

}
