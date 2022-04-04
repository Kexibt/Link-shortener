package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	password  = "277353"
	dbname    = "urls"
	ErrorPage = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
)

func ConnectDB() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	for err, i := db.Ping(), 0; err != nil && i < 5; i++ {
		time.Sleep(time.Millisecond * 1000)
		db, err = sql.Open("postgres", psqlconn)

		fmt.Println(i, "st try")
	}
	if err != nil {
		log.Panic(err)
		return nil
	}

	db.Exec("CREATE TABLE IF NOT EXIST urls")

	return db
}

func FindActualDB(ID int) string {
	db := ConnectDB()
	if db == nil {
		log.Panic("database is not connected")
	}
	defer db.Close()

	if ID > GetLastDB() {
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

func GetLastDB() int {
	db := ConnectDB()
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

func WasHereDB(actualLink string) int {
	db := ConnectDB()
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

func InsertDB(ID int, actualLink string) {
	db := ConnectDB()
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
