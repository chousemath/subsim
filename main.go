package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/chousemath/subsim/models"
	"github.com/chousemath/subsim/utils"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("v0.0.1")
	defer cleanup()

	f, err := os.Create("./subscription.db")
	utils.CheckErr(err)
	f.Close()

	// open up database connection
	db, err := sql.Open("sqlite3", "./subscription.db")
	utils.CheckErr(err)

	// create all database tables
	err = models.CreateTableVehicle(db)
	utils.CheckErr(err)
	err = models.CreateTableSubscription(db)
	utils.CheckErr(err)
	err = models.CreateTablePerson(db)
	utils.CheckErr(err)
	err = models.CreateTablePayment(db)
	utils.CheckErr(err)
	err = models.CreateTablePurchase(db)
	utils.CheckErr(err)
	err = models.CreateTableSale(db)
	utils.CheckErr(err)
}

func cleanup() {
	os.Remove("./subscription.db")
	os.Remove("./subscription.db-journal")
}
