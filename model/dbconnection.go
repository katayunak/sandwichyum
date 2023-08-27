package model

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connection() *sql.DB {
	testDB, errConnection := sql.Open("mysql", "root:newpassword@/sandwichyum")
	if errConnection != nil {
		log.Println(errConnection.Error())
	}
	errPing := testDB.Ping()
	if errPing != nil {
		log.Println(errPing.Error())
	} else {
		log.Println("Ping OK :)")
	}
	DB = testDB
	return testDB
}
