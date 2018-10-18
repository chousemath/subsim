package models

import "database/sql"

// Purchase xxx
type Purchase struct {
	ID       int64
	Original float64
	Ereg     float64
	Epre     float64
	Total    float64
}

// CreateTablePurchase xxx
func CreateTablePurchase(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE "Purchase" (
		"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"Original"	REAL NOT NULL,
		"Ereg"	REAL NOT NULL,
		"Epre"	REAL NOT NULL,
		"Total"	REAL NOT NULL
	);`)
	return err
}
