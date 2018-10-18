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
	err := os.Remove("./subscription.db")
	utils.LogErr(err)
	f, err := os.Create("./subscription.db")
	utils.CheckErr(err)
	f.Close()
	db, err := sql.Open("sqlite3", "./subscription.db")
	utils.CheckErr(err)
	fmt.Println(&db)
	err = models.CreateTableVehicle(db)
	utils.CheckErr(err)
}
