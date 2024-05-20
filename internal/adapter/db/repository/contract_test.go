package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestCreateContract(t *testing.T) {
	mtDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtDb.Run("should create and return a new contract", func(mt *mtest.T) {
		contract := domain.ContractData{
			Id:             primitive.NewObjectID(),
			ContractName:   "contrato teste",
			Description:    "teste de contrato",
			File:           []byte("teste"),
			ContractValue:  2300,
			SalaryDiscount: 0,
			Course:         "faculdade",
		}

		mt.AddMockResponses(
			bson.D{
				{Key: "ok", Value: 1},
				{Key: "n", Value: 1},
				{Key: "acknowledged", Value: true},
			},
			mtest.CreateCursorResponse(
				1,
				fmt.Sprintf(".%s%s", mt.DB.Name(), "contract"),
				mtest.FirstBatch,
				bson.D{
					{"_id", contract.Id},
					{"contractName", contract.ContractName},
					{"description", contract.Description},
					{"file", contract.File},
					{"contractValue", contract.ContractValue},
					{"salaryDiscount", contract.SalaryDiscount},
					{"course", contract.Course},
				},
			),
		)

		managerWorker := db.ManagerWorker{
			Client:   mt.Client,
			Database: mt.DB.Name(),
		}

		repo := NewContractRepository(managerWorker)

		createdContract, err := repo.CreateContract(contract)

		assert.Nil(t, err)
		assert.NotNil(t, createdContract)
		assert.Equal(t, contract.ContractName, createdContract.ContractName)
	})
}

func TestContractUpdate(t *testing.T) {
	mtDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtDb.Run("should update an existing contract", func(mt *mtest.T) {
		id := primitive.NewObjectID()

		newContractData := domain.ContractData{
			ContractName:   "Contrato para testes",
			Description:    "teste de contrato",
			File:           []byte("teste"),
			ContractValue:  2400,
			SalaryDiscount: 10,
			Course:         "faculdade",
		}

		mt.AddMockResponses(
			bson.D{
				{Key: "ok", Value: 1},
				{"value", bson.D{
					{"_id", id},
					{"contractName", newContractData.ContractName},
					{"description", newContractData.Description},
					{"file", newContractData.File},
					{"contractValue", newContractData.ContractValue},
					{"salaryDiscount", newContractData.SalaryDiscount},
					{"course", newContractData.Course},
				}},
			},
			mtest.CreateCursorResponse(1, fmt.Sprintf(".%s%s", mt.DB.Name(), "contract"),
				mtest.FirstBatch, bson.D{
					{Key: "ok", Value: 1},
					{"value", bson.D{
						{"_id", id},
						{"contractName", newContractData.ContractName},
						{"description", newContractData.Description},
						{"file", newContractData.File},
						{"contractValue", newContractData.ContractValue},
						{"salaryDiscount", newContractData.SalaryDiscount},
						{"course", newContractData.Course},
					}},
				}),
		)

		managerWorker := db.ManagerWorker{
			Client:   mt.Client,
			Database: mt.DB.Name(),
		}

		repo := NewContractRepository(managerWorker)

		updateContract, err := repo.UpdateContract(newContractData, id)

		assert.Nil(t, err)
		assert.Equal(t, newContractData.ContractValue, updateContract.ContractValue)
		assert.Equal(t, newContractData.Description, updateContract.Description)
		assert.Equal(t, newContractData.SalaryDiscount, updateContract.SalaryDiscount)
	})
}

func TestContractDelete(t *testing.T) {
	mtDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtDb.Run("Should delete existing contract", func(mt *mtest.T) {
		worker := db.ManagerWorker{
			Client:   mt.DB.Client(),
			Database: mt.DB.Name(),
		}

		repo := NewContractRepository(worker)

		mt.AddMockResponses(bson.D{{"ok", 1}, {"value", bson.D{}}})
		err := repo.DeleteContract(primitive.NewObjectID())

		assert.Nil(t, err)
	})
}
