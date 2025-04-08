package models

import (
	"errors"

	"github.com/google/uuid"

)

// Engine is the model of the engine
type Engine struct {
	EngineID      uuid.UUID `json:"engine_id" `
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int64     `json:"noOfCylinders"`
	CarRange      int64     `json:"carRange"`
}

type EngineRequest struct {
	Displacement  int64 `json:"displacement"`
	NoOfCylinders int64 `json:"noOfCylinders"`
	CarRange      int64 `json:"carRange"`
}

// ValideteEngineRequest validates the request body
func ValideteEngineRequest(engineReq EngineRequest) error {
	if err := validateDisplacement(engineReq.Displacement); err != nil {
		return err
	}
	if err := validateNoOfCylinders(engineReq.NoOfCylinders); err != nil {
		return err
	}
	if err := validateCarRange(engineReq.CarRange); err != nil {
		return err
	}
	return nil
}

// valideteDisplacement validates the displacement of the engine
func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("Displacement must be greater than zero")
	}
	return nil
}

// valideteNoOfCylinders validates the noOfCylinders of the engine
func validateNoOfCylinders(noOfCylinders int64) error {
	if noOfCylinders <= 0 {
		return errors.New("NoOfCylinders must be greater than zero")
	}
	return nil
}

// valideteCarRange validates the carRange of the engine
func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("CarRange must be greater than zero")
	}
	return nil
}
