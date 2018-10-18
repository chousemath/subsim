package models

import "database/sql"

// SubscriptionStatus xxx
type SubscriptionStatus uint

const (
	// UndefinedSubscriptionStatus RoundStatus xxx
	UndefinedSubscriptionStatus SubscriptionStatus = 0
	// First RoundStatus xxx
	First SubscriptionStatus = 1
	// Continuation RoundStatus xxx
	Continuation SubscriptionStatus = 2
	// Final RoundStatus xxx
	Final SubscriptionStatus = 3
	// FirstAndFinal RoundStatus xxx
	FirstAndFinal SubscriptionStatus = 4
	// Removed RoundStatus xxx
	Removed SubscriptionStatus = 5
)

// Subscription xxx
type Subscription struct {
	ID           int64
	PersonID     int64
	VehicleID    int64
	StartDate    uint
	EndDate      uint
	DecisionDate uint
	Status       SubscriptionStatus
	Assigned     bool
	Finished     bool
	Swap         bool
}

// CreateTableSubscription xxx
func CreateTableSubscription(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE "Subscription" (
		"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"PersonID"	INTEGER NOT NULL,
		"VehicleID"	INTEGER NOT NULL,
		"StartDate"	INTEGER NOT NULL,
		"EndDate"	INTEGER NOT NULL,
		"DecisionDate"	INTEGER NOT NULL,
		"Status"	INTEGER NOT NULL,
		"Assigned"	INTEGER NOT NULL,
		"Finished"	INTEGER NOT NULL,
		"Swap"	INTEGER NOT NULL
	);`)
	return err
}
