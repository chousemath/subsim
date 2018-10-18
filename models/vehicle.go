package models

import (
	"database/sql"
)

// VehicleStatus xxx
type VehicleStatus uint

// Vehicle xxx
type Vehicle struct {
	Make               string
	Model              string
	SubModel           string
	Trim               string
	Name               string
	Price              uint
	KM                 uint
	Year               uint
	DaysInSubscription uint
	Category           uint
	SalesDate          uint
	FinalSalesDate     uint
	BoughtAt           uint
	Status             VehicleStatus
	ID                 int64
	PersonID           int64
	SubscriptionID     int64
}

const (
	// UndefinedVehicleStatus xxx
	UndefinedVehicleStatus VehicleStatus = 0
	// InSubscription xxx
	InSubscription VehicleStatus = 1
	// CanBeSwapped xxx
	CanBeSwapped VehicleStatus = 2
	// Inactive xxx
	Inactive VehicleStatus = 3
	// ReadyForSale xxx
	ReadyForSale VehicleStatus = 4
	// Sold xxx
	Sold VehicleStatus = 5
)

// CreateTableVehicle xxx
func CreateTableVehicle(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE "Vehicle" (
		"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"Make"	TEXT NOT NULL,
		"Model"	TEXT NOT NULL,
		"SubModel"	TEXT NOT NULL,
		"Trim"	TEXT NOT NULL,
		"Name"	TEXT NOT NULL,
		"Price"	INTEGER NOT NULL,
		"KM"	INTEGER NOT NULL,
		"Year"	INTEGER NOT NULL,
		"DaysInSubscription"	INTEGER NOT NULL,
		"Category"	INTEGER NOT NULL,
		"Status"	INTEGER NOT NULL,
		"SalesDate"	INTEGER NOT NULL,
		"FinalSalesDate"	INTEGER NOT NULL,
		"PersonID"	INTEGER NOT NULL,
		"SubscriptionID"	INTEGER NOT NULL,
		"BoughtAt"	INTEGER NOT NULL
	);`)
	return err
}
