package models

import "database/sql"

// YearlyDepreciation xxx
type YearlyDepreciation uint

const (
	// UndefinedYearlyDepreciation xxx
	UndefinedYearlyDepreciation YearlyDepreciation = 0
	// DiscreteYears xxx
	DiscreteYears YearlyDepreciation = 1
	// Daily xxx
	Daily YearlyDepreciation = 2
)

// Sale xxx
type Sale struct {
	ID                 int64
	DaysInSubscription uint
	YearlyDepreciation YearlyDepreciation
	Original           float64
	ResellPercentage   float64
	Epre               float64
	Total              float64
	KMDepreciation     float64
	InAccident         bool
}

// CreateTableSale xxx
func CreateTableSale(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE "Sale" (
		"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"DaysInSubscription"	INTEGER NOT NULL,
		"YearlyDepreciation"	INTEGER NOT NULL,
		"Original"	REAL NOT NULL,
		"ResellPercentage"	REAL NOT NULL,
		"Epre"	REAL NOT NULL,
		"Total"	REAL NOT NULL,
		"KMDepreciation"	REAL NOT NULL,
		"InAccident"	INTEGER NOT NULL
	);`)
	return err
}
