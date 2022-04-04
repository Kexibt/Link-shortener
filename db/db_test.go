package db

import (
	"testing"

	"github.com/Kexibt/Link-shortener/url"
)

func TestConnect(t *testing.T) {
	db := ConnectDB()

	err := db.Ping()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestFindActualDB(t *testing.T) {
	act := FindActualDB(url.ConvertToID("iDontThinЯ"))
	if act != "error" && act != ErrorPage {
		t.Log("not error when error")
		t.Fail()
	}
}

func TestFindInsertAndWasHereDB(t *testing.T) {
	link := "test"

	InsertDB(GetLastDB()+1, link)
	ind := WasHereDB(link)

	if ind <= 0 {
		t.Log("insertDB() or wasHere() is broken.")
		t.Fail()
	}
}

func TestGetLastDB(t *testing.T) {
	lastID := GetLastDB()

	db := ConnectDB()
	if err := db.Ping(); err != nil {
		t.Log(err)
		t.Fail()
	}

	res, err := db.Query("SELECT COUNT(*) FROM urls")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	defer res.Close()

	var ind int
	res.Next()
	err = res.Scan(&ind)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if ind != lastID {
		t.Log("getLastDB() is broken.")
		t.Fail()
	}
}
