package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB){
	// Create Train Table
	statement, driverError := dbDriver.Prepare(train)
	if driverError != nil{
		log.Println(driverError)
	}
	_, err := statement.Exec()
	if err != nil{
		log.Fatalf("table train already exists!")
	}

	// Create Station Table
	statement, driverError = dbDriver.Prepare(station)
	if driverError != nil{
		log.Println(driverError)
	}
	_, err = statement.Exec()
	if err != nil{
		log.Fatalf("table station already exists!")
	}

	// Create Schedule Table
	statement, driverError = dbDriver.Prepare(schedule)
	if driverError != nil{
		log.Println(driverError)
	}
	_, err = statement.Exec()
	if err != nil{
		log.Fatalf("table schedule already exists!")
	}
}
