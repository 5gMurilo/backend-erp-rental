package repository

import (
	"america-rental-backend/internal/adapter/db"
	"america-rental-backend/internal/core/domain"
	"america-rental-backend/internal/core/ports"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepositoryImpl struct {
	db db.ManagerWorker
}

func NewEmployeeRepositoryImpl(db db.ManagerWorker) ports.EmployeeRepository {
	return &EmployeeRepositoryImpl{db}
}

// Delete implements ports.EmployeeRepository.
func (e *EmployeeRepositoryImpl) Delete(ctx context.Context, id primitive.ObjectID) error {
	panic("unimplemented")
}

// GetAll implements ports.EmployeeRepository.
func (e *EmployeeRepositoryImpl) GetAll(ctx context.Context) ([]*domain.Employee, error) {
	panic("unimplemented")
}

// GetByCPF implements ports.EmployeeRepository.
func (e *EmployeeRepositoryImpl) GetByCPF(ctx context.Context, cpf string) (*domain.Employee, error) {
	var employee domain.Employee
	err := e.db.GetCollection("employee").FindOne(ctx, bson.M{"cpf": cpf}).Decode(&employee)
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

// GetById implements ports.EmployeeRepository.
func (e *EmployeeRepositoryImpl) GetById(ctx context.Context, id primitive.ObjectID) (*domain.Employee, error) {
	var employee domain.Employee
	err := e.db.GetCollection("employee").FindOne(ctx, bson.M{"_id": id}).Decode(&employee)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

// New implements ports.EmployeeRepository.
func (e *EmployeeRepositoryImpl) New(ctx context.Context, employee domain.Employee) (*primitive.ObjectID, error) {
	session, err := e.db.StartSession()
	if err != nil {
		return nil, err
	}

	employee.Id = primitive.NewObjectID()
	employee.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	employee.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	rst, err := e.db.GetCollection("employee").Find(ctx, bson.M{"cpf": employee.Cpf})
	if err != nil {
		return nil, err
	}
	defer rst.Close(ctx)

	if rst.Next(ctx) {
		return nil, errors.New("colaborador já existente")
	}

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		rst, err := e.db.GetCollection("employee").InsertOne(ctx, employee)
		if err != nil {
			session.AbortTransaction(sc)
			return err
		}
		if rst.InsertedID == nil {
			session.AbortTransaction(sc)
			return errors.New("erro ao concluir operação de insert")
		}

		if err := session.CommitTransaction(sc); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &employee.Id, nil
}

// Update implements ports.EmployeeRepository.
func (e *EmployeeRepositoryImpl) Update(ctx context.Context, id primitive.ObjectID, newData domain.Employee) (*domain.Employee, error) {
	var data domain.Employee
	session, err := e.db.StartSession()
	if err != nil {
		return nil, err
	}

	data = domain.Employee{
		Id:                    id,
		Name:                  newData.Name,
		Role:                  newData.Role,
		Sector:                newData.Sector,
		HiringRegime:          newData.HiringRegime,
		CivilStatus:           newData.CivilStatus,
		Salary:                newData.Salary,
		WorkDays:              newData.WorkDays,
		EntryTime:             newData.EntryTime,
		ExitTime:              newData.ExitTime,
		Address:               newData.Address,
		Neighborhood:          newData.Neighborhood,
		City:                  newData.City,
		State:                 newData.State,
		Cep:                   newData.Cep,
		Phone:                 newData.Phone,
		MothersName:           newData.MothersName,
		FathersName:           newData.FathersName,
		EducationLevel:        newData.EducationLevel,
		BornDate:              newData.BornDate,
		Rg:                    newData.Rg,
		RgExpiration:          newData.RgExpiration,
		RgExpedition:          newData.RgExpedition,
		Forwarder:             newData.Forwarder,
		Naturalness:           newData.Naturalness,
		Cpf:                   newData.Cpf,
		VoterRegistration:     newData.VoterRegistration,
		Zone:                  newData.Zone,
		Section:               newData.Section,
		Reservist:             newData.Reservist,
		ReservistNumber:       newData.ReservistNumber,
		ReservistSerie:        newData.ReservistSerie,
		Cnh:                   newData.Cnh,
		CnhExpiration:         newData.CnhExpiration,
		Pis:                   newData.Pis,
		Ctps:                  newData.Ctps,
		CtpsSerie:             newData.CtpsSerie,
		AsoExpiration:         newData.AsoExpiration,
		VacationExpiration:    newData.VacationExpiration,
		AdmissionDate:         newData.AdmissionDate,
		Spouse:                newData.Spouse,
		Dependents:            newData.Dependents,
		Allergies:             newData.Allergies,
		Disease:               newData.Disease,
		EmergencyInstructions: newData.EmergencyInstructions,
		Unhealthiness:         newData.Unhealthiness,
		TransportCard:         newData.TransportCard,
		FuelAid:               newData.FuelAid,
		EmployeeState:         newData.EmployeeState,
		Experience:            newData.Experience,
		ShirtSize:             newData.ShirtSize,
		PantsSize:             newData.PantsSize,
		BootSize:              newData.BootSize,
		EpiFile:               newData.EpiFile,
		CreatedAt:             data.CreatedAt,
		UpdatedAt:             primitive.NewDateTimeFromTime(time.Now()),
		ModifiedBy:            newData.ModifiedBy,
	}

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		err = e.db.GetCollection("employee").FindOneAndReplace(ctx, bson.M{"_id": id}, bson.M{"_set": data}).Decode(&data)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &data, nil
}
