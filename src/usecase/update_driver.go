package usecase

import "time"

type UpdateDriverInput struct {
	CNH        *string `json:"cnh"`
	Gender     *string `json:"gender"`
	HasVehicle *bool   `json:"has_vehicle"`
	Name       *string `json:"name"`
}

type UpdateDriverOutput struct {
	BirthDate  time.Time `json:"birth_date"`
	CNH        string    `json:"cnh"`
	CPF        string    `json:"cpf"`
	Gender     string    `json:"gender"`
	HasVehicle bool      `json:"has_vehicle"`
	Name       string    `json:"name"`
	UpdatedAt  time.Time `json:"updated_at"`
}
