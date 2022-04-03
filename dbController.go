package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host      = "127.0.0.1"
	port      = 5432
	user      = "postgres"
	password  = "277353"
	dbname    = "urls"
	errorPage = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
)

func connectDB() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Panic(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
		return nil
	}

	return db
}

func findActualDB(ID int) string {
	db := connectDB()
	if db == nil {
		log.Panic("database is not connected")
	}
	defer db.Close()

	if ID > getLastDB() {
		return "error"
	}

	selectStmt := "SELECT * FROM urls WHERE id = " + strconv.Itoa(ID)
	res, err := db.Query(selectStmt)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()
	res.Next()

	var id int
	var link string
	defer res.Close()

	err = res.Scan(&id, &link)
	if err != nil {
		log.Panic(err)
	}

	return link
}

func getLastDB() int {
	db := connectDB()
	if db == nil {
		log.Panic("database is not connected")
	}
	defer db.Close()

	selectStmt := "SELECT COUNT(*) FROM urls"
	res, err := db.Query(selectStmt)
	if err != nil {
		log.Fatal(err)
	}

	var lastID int
	defer res.Close()

	res.Next()
	err = res.Scan(&lastID)
	if err != nil {
		log.Panic(err)
	}

	return lastID
}

func wasHereDB(actualLink string) int {
	db := connectDB()
	if db == nil {
		log.Panic("database is not connected")
	}
	defer db.Close()

	selectStmt := "SELECT * FROM urls WHERE longurl LIKE '%" + actualLink + "%'"
	res, err := db.Query(selectStmt)
	if err != nil {
		log.Fatal(err)
	}

	var ID int
	var link string
	defer res.Close()

	res.Next()
	err = res.Scan(&ID, &link)
	if err != nil {
		return -1
	}
	return ID
}

func insertDB(ID int, actualLink string) {
	db := connectDB()
	if db == nil {
		log.Panic("database is not connected")
	}
	defer db.Close()

	selectStmt := "INSERT INTO urls(id, longURL) VALUES (" + strconv.Itoa(ID)
	selectStmt = selectStmt + ", '" + actualLink + "')"

	_, err := db.Exec(selectStmt)
	if err != nil {
		log.Fatal(err)
	}
}
