package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	Id                    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name                  string             `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Role                  string             `json:"role,omitempty" bson:"role,omitempty" binding:"required"`
	Sector                string             `json:"sector,omitempty" bson:"sector,omitempty" binding:"required"`
	HiringRegime          string             `json:"hiringRegime,omitempty" bson:"hiringRegime,omitempty" binding:"required"`
	CivilStatus           string             `json:"civilStatus,omitempty" bson:"civilStatus,omitempty" binding:"required"`
	Salary                float64            `json:"salary,omitempty" bson:"salary,omitempty" binding:"required"`
	WorkDays              []string           `json:"workDays,omitempty" bson:"workDays,omitempty" binding:"required"`
	EntryTime             string             `json:"entryTime,omitempty" bson:"entryTime,omitempty" binding:"required"`
	ExitTime              string             `json:"exitTime,omitempty" bson:"exitTime,omitempty" binding:"required"`
	Address               string             `json:"address,omitempty" bson:"address,omitempty" binding:"required"`
	Neighborhood          string             `json:"neighborhood,omitempty" bson:"neighborhood,omitempty" binding:"required"`
	City                  string             `json:"city,omitempty" bson:"city,omitempty" binding:"required"`
	State                 string             `json:"state,omitempty" bson:"state,omitempty" binding:"required"`
	Cep                   string             `json:"cep,omitempty" bson:"cep,omitempty" binding:"required"`
	Phone                 string             `json:"phone,omitempty" bson:"phone,omitempty" binding:"required"`
	MothersName           string             `json:"mothersName,omitempty" bson:"mothersName,omitempty" binding:"required"`
	FathersName           *string            `json:"fathersName,omitempty" bson:"fathersName,omitempty"`
	EducationLevel        string             `json:"educationLevel,omitempty" bson:"educationLevel,omitempty" binding:"required"`
	BornDate              time.Time          `json:"bornDate,omitempty" bson:"bornDate,omitempty" binding:"required"`
	Rg                    string             `json:"rg,omitempty" bson:"rg,omitempty" binding:"required"`
	RgExpiration          time.Time          `json:"rgExpiration,omitempty" bson:"rgExpiration,omitempty" binding:"required"`
	RgExpedition          time.Time          `json:"rgExpedition,omitempty" bson:"rgExpedition,omitempty" binding:"required"`
	Forwarder             string             `json:"forwarder,omitempty" bson:"forwarder,omitempty" binding:"required"`
	Naturalness           string             `json:"naturalness,omitempty" bson:"naturalness,omitempty" binding:"required"`
	Cpf                   string             `json:"cpf" bson:"cpf" binding:"required"`
	VoterRegistration     string             `json:"voterRegistration,omitempty" bson:"voterRegistration,omitempty" binding:"required"`
	Zone                  *string            `json:"zone,omitempty" bson:"zone,omitempty"`
	Section               *string            `json:"section,omitempty" bson:"section,omitempty"`
	Reservist             *bool              `json:"reservist,omitempty" bson:"reservist,omitempty"`
	ReservistNumber       *string            `json:"reservistNumber,omitempty" bson:"reservistNumber,omitempty"`
	ReservistSerie        *string            `json:"reservistSerie,omitempty" bson:"reservistSerie,omitempty"`
	Cnh                   *string            `json:"cnh,omitempty" bson:"cnh,omitempty"`
	CnhExpiration         *time.Time         `json:"cnhExpiration,omitempty" bson:"cnhExpiration,omitempty"`
	Pis                   string             `json:"pis" bson:"pis" binding:"required"`
	Ctps                  *string            `json:"ctps,omitempty" bson:"ctps,omitempty"`
	CtpsSerie             *string            `json:"ctpsSerie" bson:"ctpsSerie"`
	AsoExpiration         time.Time          `json:"asoExpiration" bson:"asoExpiration" binding:"required"`
	VacationExpiration    *time.Time         `json:"vacationExpiration" bson:"vacationExpiration" binding:"required"`
	AdmissionDate         *time.Time         `json:"admissionDate,omitempty" bson:"admissionDate,omitempty"`
	Spouse                *Spouse            `json:"spouse,omitempty" bson:"spouse,omitempty"`
	Dependents            *[]Dependents      `json:"dependents,omitempty" bson:"dependents,omitempty"`
	Allergies             *string            `json:"allergies,omitempty" bson:"allergies,omitempty"`
	Disease               *string            `json:"disease,omitempty" bson:"disease,omitempty"`
	EmergencyInstructions *string            `json:"emergencyInstructions,omitempty" bson:"emergencyInstructions,omitempty"`
	Unhealthiness         *bool              `json:"unhealthiness,omitempty" bson:"unhealthiness,omitempty"`
	TransportCard         *bool              `json:"transportCard,omitempty" bson:"transportCard,omitempty"`
	FuelAid               *bool              `json:"fuelAid,omitempty" bson:"fuelAid,omitempty"`
	EmployeeState         *string            `json:"employeeState,omitempty" bson:"employeeState,omitempty"`
	Experience            *bool              `json:"experience,omitempty" bson:"experience,omitempty"`
	ShirtSize             string             `json:"shirtSize" bson:"shirtSize" binding:"required"`
	PantsSize             string             `json:"pantsSize" bson:"pantsSize" binding:"required"`
	BootSize              string             `json:"bootSize" bson:"bootSize" binding:"required"`
	EpiFile               string             `json:"epiFile" bson:"epiFile" binding:"required"`
	CreatedAt             primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt             primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	ModifiedBy            string             `json:"modifiedBy" bson:"modifiedBy"`
}

type Dependents struct {
	DependentName string `json:"dependentName,omitempty" bson:"dependentName,omitempty" binding:"required"`
	DependentRG   string `json:"dependentRg,omitempty" bson:"dependentRg,omitempty" binding:"required"`
	DependentCPF  string `json:"dependentCpf,omitempty" bson:"dependentCpf,omitempty" binding:"required"`
}

type Spouse struct {
	SpouseName string `json:"spouseName,omitempty" bson:"spouseName,omitempty" binding:"required"`
	SpouseRG   string `json:"spouseRG,omitempty" bson:"spouseRG,omitempty" binding:"required"`
	SpouseCpf  string `json:"spouseCpf,omitempty" bson:"spouseCpf,omitempty" binding:"required"`
}
