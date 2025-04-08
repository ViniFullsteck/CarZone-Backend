package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// Car is the model of the car
type Car struct {
	ID        uuid.UUID ` json:"id" `
	Name      string    ` json:"name" `
	Year      string    ` json:"year" `
	Brand     string    ` json:"brand" `
	Fueltype  string    ` json:"fueltype" `
	Engine    Engine    ` json:"engine" `
	Price     float64   ` json:"price" `
	CreatedAt time.Time ` json:"created_at" `
	UpdatedAt time.Time ` json:"updated_at" `
}

// CarRequest is the request body for creating a car
type CarRequest struct {
	Name     string  ` json:"name" `
	Year     string  ` json:"year" `
	Brand    string  ` json:"brand" `
	Fueltype string  ` json:"fueltype" `
	Engine   Engine  ` json:"engine" `
	Price    float64 ` json:"price" `
}

// ValideteRequest validates the request body
func ValideteRequest(carReq CarRequest) error {
	if err := valideteYear(carReq.Year); err != nil {
		return err
	}
	if err := valideteBrand(carReq.Brand); err != nil {
		return err
	}
	if err := valideteFueltype(carReq.Fueltype); err != nil {
		return err
	}
	if err := valideteEngine(carReq.Engine); err != nil {
		return err
	}
	if err := validetePrice(carReq.Price); err != nil {
		return err
	}
	return nil
}

// valideteName validates the name of the car
func valideteName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}

// valideteYear validates the year of the car
func valideteYear(year string) error {
	if year == "" {
		return errors.New("year is required")
	}
	_, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("year must be a valid number")
	}
	currentYear := time.Now().Year()
	yearInt, _ := strconv.Atoi(year)
	if yearInt < 1886 || yearInt > currentYear {
		return errors.New("year must be between 1886 and corren year")
	}
	return nil
}

// valideteBrand validates the brand of the car
func valideteBrand(brand string) error {
	if brand == "" {
		return errors.New("brand is required")
	}
	return nil
}

// valideteFueltype validates the fueltype of the car
func valideteFueltype(fueltype string) error {
	valideteFueltypes := []string{"persol", "Diesel", "Electric", "Hybrid"}
	for _, validType := range valideteFueltypes {
		if fueltype == validType {
			return nil
		}
	}
	return errors.New("Fueltype Must be one of the following: Persol, Diesel, Electric, Hybrid")
}

// valideteEngine validates the engine of the car
func valideteEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return errors.New("EngineID is required")
	}
	if engine.Displacement <= 0 {
		return errors.New("Displacement must be greater than zero")
	}
	if engine.NoOfCylinders <= 0 {
		return errors.New("NoOfCylinders must be greater than zero")
	}
	if engine.CarRange <= 0 {
		return errors.New("CarRange must be greater than zero")
	}
	return nil
}

// validetePrice validates the price of the car
func validetePrice(price float64) error {
	if price <= 0 {
		return errors.New("Price must be greater than is Required")
	}
	return nil
}
