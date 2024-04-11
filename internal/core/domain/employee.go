package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	Id                    primitive.ObjectID `json:"id" bson:"id, omitempty"`
	Name                  string             `json:"name" bson:"name, omitempty" binding:"required"`
	Role                  string             `json:"role" bson:"role, omitempty" binding:"required"`
	Sector                string             `json:"sector" bson:"sector, omitempty" binding:"required"`
	HiringRegime          string             `json:"hiringRegime" bson:"hiringRegime, omitempty" binding:"required"`
	CivilStatus           string             `json:"civilStatus" bson:"civilStatus, omitempty" binding:"required"`
	Salary                float64            `json:"salary" bson:"salary, omitempty" binding:"required"`
	WorkDays              []string           `json:"workDays" bson:"workDays, omitempty" binding:"required"`
	EntryTime             string             `json:"entryTime" bson:"entryTime, omitempty" binding:"required"`
	ExitTime              string             `json:"exitTime" bson:"exitTime, omitempty" binding:"required"`
	Address               string             `json:"address" bson:"address, omitempty" binding:"required"`
	Neighborhood          string             `json:"neighborhood" bson:"neighborhood, omitempty" binding:"required"`
	City                  string             `json:"city" bson:"city, omitempty" binding:"required"`
	State                 string             `json:"state" bson:"state, omitempty" binding:"required"`
	Cep                   string             `json:"cep" bson:"cep, omitempty" binding:"required"`
	Phone                 string             `json:"phone" bson:"phone, omitempty" binding:"required"`
	MothersName           string             `json:"mothersName" bson:"mothersName, omitempty" binding:"required"`
	FathersName           string             `json:"fathersName" bson:"fathersName, omitempty"`
	EducationLevel        string             `json:"educationLevel" bson:"educationLevel, omitempty" binding:"required"`
	BornDate              primitive.DateTime `json:"bornDate" bson:"bornDate, omitempty" binding:"required"`
	Rg                    string             `json:"rg" bson:"rg, omitempty" binding:"required"`
	RgExpiration          primitive.DateTime `json:"rgExpiration" bson:"rgExpiration, omitempty" binding:"required"`
	RgExpedition          primitive.DateTime `json:"rgExpedition" bson:"rgExpedition, omitempty" binding:"required"`
	Forwarder             string             `json:"forwarder" bson:"forwarder, omitempty" binding:"required"`
	Naturalness           string             `json:"naturalness" bson:"naturalness, omitempty" binding:"required"`
	Cpf                   string             `json:"cpf" bson:"cpf" binding:"required"`
	VoterRegistration     string             `json:"voterRegistration" bson:"voterRegistration, omitempty" binding:"required"`
	Zone                  string             `json:"zone" bson:"zone, omitempty"`
	Section               string             `json:"section" bson:"section, omitempty"`
	Reservist             bool               `json:"reservist" bson:"reservist, omitempty"`
	ReservistNumber       string             `json:"reservistNumber" bson:"reservistNumber, omitempty"`
	ReservistSerie        string             `json:"reservistSerie" bson:"reservistSerie, omitempty"`
	Cnh                   string             `json:"cnh" bson:"cnh, omitempty"`
	CnhExpiration         primitive.DateTime `json:"cnhExpiration" bson:"cnhExpiration, omitempty"`
	Pis                   string             `json:"pis" bson:"pis" binding:"required"`
	Ctps                  string             `json:"ctps" bson:"ctps, omitempty"`
	CtpsSerie             string             `json:"ctpsSerie" bson:"ctpsSerie"`
	AsoExpiration         primitive.DateTime `json:"asoExpiration" bson:"asoExpiration" binding:"required"`
	VacationExpiration    primitive.DateTime `json:"vacationExpiration" bson:"vacationExpiration" binding:"required"`
	AdmissionDate         primitive.DateTime `json:"admissionDate" bson:"admissionDate, omitempty"`
	Spouse                Spouse             `json:"spouse" bson:"spouse, omitempty"`
	Dependents            []Dependents       `json:"dependents" bson:"dependents, omitempty"`
	Allergies             string             `json:"allergies" bson:"allergies, omitempty"`
	Disease               string             `json:"disease" bson:"disease, omitempty"`
	EmergencyInstructions string             `json:"emergencyInstructions" bson:"emergencyInstructions, omitempty"`
	Unhealthiness         *bool              `json:"unhealthiness" bson:"unhealthiness, omitempty"`
	TransportCard         *bool              `json:"transportCard" bson:"transportCard, omitempty"`
	FuelAid               *bool              `json:"fuelAid" bson:"fuelAid, omitempty"`
	EmployeeState         string             `json:"employeeState" bson:"employeeState, omitempty"`
	Experience            *bool              `json:"experience" bson:"experience, omitempty"`
	ShirtSize             string             `json:"shirtSize" bson:"shirtSize" binding:"required"`
	PantsSize             string             `json:"pantsSize" bson:"pantsSize" binding:"required"`
	BootSize              string             `json:"bootSize" bson:"bootSize" binding:"required"`
	EpiFile               string             `json:"epiFile" bson:"epiFile" binding:"required"`
	CreatedAt             primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt             primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	ModifiedBy            string             `json:"modifiedBy" bson:"modifiedBy"`
}

type Dependents struct {
	DependentName string `json:"dependentName" bson:"dependentName, omitempty" binding:"required"`
	DependentRG   string `json:"dependentRg" bson:"dependentName, omitempty" binding:"required"`
	DependentCPF  string `json:"dependentCpf" bson:"dependentName, omitempty" binding:"required"`
}

type Spouse struct {
	SpouseName string `json:"spouseName" bson:"spouseName, omitempty" binding:"required"`
	SpouseRG   string `json:"spouseRG" bson:"spouseRG, omitempty" binding:"required"`
	SpouseCpf  string `json:"spouseCpf" bson:"spouseCpf, omitempty" binding:"required"`
}
