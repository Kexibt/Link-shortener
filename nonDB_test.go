package main

import (
	"testing"
)

func TestGetMap(t *testing.T) {
	getDataNonDB()
	last := len(data)
	data[last+1] = "test"

	if _, exist := data[last+1]; !exist {
		t.Log("getDataNonDB is broken.")
		t.Fail()
	} else {
		delete(data, last+1)
	}
}

func TestFindActualNonDB(t *testing.T) {
	act := findActualNonDB(convertToID("iDontThinЯ"))
	if act != "error" && act != errorPage {
		t.Log("not error when error. (non-db)")
		t.Fail()
	}
}

func TestFindInsertAndWasHereNonDB(t *testing.T) {
	link := "test"

	insertNonDB(getLastDB()+1, link)
	ind := wasHereNonDB(link)

	if ind <= 0 {
		t.Log("insertNonDB() or wasHereNonDB() is broken.")
		t.Fail()
	}
}

func TestGetLastDBNonDB(t *testing.T) {
	lastID := getLastNonDB()

	getDataNonDB()
	ind := len(data)

	if ind != lastID {
		t.Log("getLastNonDB() is broken.")
		t.Fail()
	}
}
