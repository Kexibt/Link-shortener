package db

import (
	"testing"

	"github.com/Kexibt/Link-shortener/url"
)

func TestGetMap(t *testing.T) {
	pathData = "../data/data.json"
	GetDataNonDB(pathData)
	last := len(Data)
	Data[last+1] = "test"

	if _, exist := Data[last+1]; !exist {
		t.Log("getDataNonDB is broken.")
		t.Fail()
	} else {
		delete(Data, last+1)
	}
	pathData = "data/data.json"
}

func TestFindActualNonDB(t *testing.T) {
	pathData = "../data/data.json"
	act := FindActualNonDB(url.ConvertToID("iDontThinЯ"))
	if act != "error" && act != ErrorPage {
		t.Log("not error when error. (non-db)")
		t.Fail()
	}
	pathData = "data/data.json"
}

func TestFindInsertAndWasHereNonDB(t *testing.T) {
	pathData = "../data/data.json"
	link := "test"

	InsertNonDB(GetLastDB()+1, link)
	ind := WasHereNonDB(link)

	if ind <= 0 {
		t.Log("insertNonDB() or wasHereNonDB() is broken.")
		t.Fail()
	}
	pathData = "data/data.json"
}

func TestGetLastDBNonDB(t *testing.T) {
	pathData = "../data/data.json"
	lastID := GetLastNonDB()

	GetDataNonDB(pathData)
	ind := len(Data)

	if ind != lastID {
		t.Log("getLastNonDB() is broken.")
		t.Fail()
	}
	pathData = "data/data.json"
}
