package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var billDB *sql.DB

func OpenDb() {
	VerifyDb()
	os.Chdir("~/.billKeeper")
	data, err := sql.Open("sqlite3", "bills.db")
	if err != nil {
		panic(err.Error())
	}

	err = data.Ping()
	if err != nil {
		panic(err.Error())
	}

	billDB = data
}

func VerifyDb() {
	var db *sql.DB

	defer db.Close()

	if _, err := os.Stat("~/.billKeeper"); os.IsNotExist(err) {
		err := os.Mkdir("~/.billKeeper", os.ModeDir)
		if err != nil {
			panic(err.Error())
		}
	}

	if _, err := os.Stat("~/.billKeeper/bills.db"); os.IsNotExist(err) {
		f, err := os.Create("~/.billKeeper/bills.db")
		defer f.Close()
		if err != nil {
			panic(err.Error())
		}
	}
	db, _ = sql.Open("sqlite3", "~/.billKeeper/bills.db")
	db.Exec("CREATE TABLE bills " +
		"(Name string, Id int primary_key auto increment," +
		" Exp int, Payed bool)")
}
