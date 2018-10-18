package models

import "database/sql"

// PersonStatus xxx
type PersonStatus uint

const (
	// ActiveUser xxx
	ActiveUser PersonStatus = 1
	// InactiveUser xxx
	InactiveUser PersonStatus = 2
)

// Person xxx
type Person struct {
	StartDate         uint
	EndDate           uint
	Insurance         uint
	SubscriptionCount uint
	Status            PersonStatus
	ID                int64
	FirstVehicleID    int64
	CurrentVehicleID  int64
}

// CreateTablePerson xxx
func CreateTablePerson(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE "Person" (
		"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"FirstVehicleID"	INTEGER NOT NULL,
		"CurrentVehicleID"	INTEGER NOT NULL,
		"StartDate"	INTEGER NOT NULL,
		"EndDate"	INTEGER NOT NULL,
		"Insurance"	INTEGER NOT NULL,
		"SubscriptionCount"	INTEGER NOT NULL,
		"Status"	INTEGER NOT NULL
	);`)
	return err
}
