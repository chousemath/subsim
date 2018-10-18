package models

import "database/sql"

// PaymentCategory xxx
type PaymentCategory uint

const (
	// UndefinedPaymentCategory xxx
	UndefinedPaymentCategory = 0
	// PaymentVehiclePurchase xxx
	PaymentVehiclePurchase = 1
	// PaymentSubscription xxx
	PaymentSubscription = 2
	// PaymentVehicleSale xxx
	PaymentVehicleSale = 3
	// PaymentRemoved xxx
	PaymentRemoved = 4
)

// Payment xxx
type Payment struct {
	Amount         float64
	Timestamp      uint
	Category       PaymentCategory
	ID             int64
	PersonID       int64
	SubscriptionID int64
	PurchaseID     int64
	SaleID         int64
	VehicleID      int64
}

// CreateTablePayment xxx
func CreateTablePayment(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE "Payment" (
		"ID"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"PersonID"	INTEGER NOT NULL,
		"SubscriptionID"	INTEGER NOT NULL,
		"PurchaseID"	INTEGER NOT NULL,
		"SaleID"	INTEGER NOT NULL,
		"VehicleID"	INTEGER NOT NULL,
		"Category"	INTEGER NOT NULL,
		"Timestamp"	INTEGER NOT NULL,
		"Amount"	REAL NOT NULL
	);`)
	return err
}
