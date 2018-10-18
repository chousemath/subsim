package models

import (
	"database/sql"
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
