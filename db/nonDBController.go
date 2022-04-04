package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var Data map[int]string

var (
	pathData = "data/data.json"
)

func GetPath() string {
	return pathData
}

func GetDataNonDB(pathData string) {
	if CheckCreateData(pathData) {
		Data = make(map[int]string)
		return
	}

	file, err := ioutil.ReadFile(pathData)
	if err != nil {
		log.Fatal(err)
	}

	if len(file) == 0 {
		Data = make(map[int]string)
		return
	}
	err = json.Unmarshal(file, &Data)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckCreateData(pathData string) bool {
	exdir, err := Exists(pathData[:strings.Index(pathData, "/")])
	if err != nil {
		log.Fatal(err)
	}

	if !exdir {
		err = os.Mkdir(pathData[:strings.Index(pathData, "/")], 0777)
		if err != nil {
			panic(err)
		}
	}

	ex, err := Exists(pathData)
	if err != nil {
		log.Fatal(err)
	}

	if !ex {
		f, err := os.Create(pathData)
		if err != nil {
			panic(err)
		}
		defer f.Close()
	} else {
		return false
	}

	return true
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func FindActualNonDB(ID int) string {
	GetDataNonDB(pathData)
	link, exist := Data[ID]
	if !exist {
		return "error"
	}

	return link
}

func GetLastNonDB() int {
	GetDataNonDB(pathData)
	return len(Data)
}

func WasHereNonDB(actualLink string) int {
	GetDataNonDB(pathData)
	for key, value := range Data {
		if value == actualLink {
			return key
		}
	}

	return -1
}

func InsertNonDB(ID int, actualLink string) {
	GetDataNonDB(pathData)
	Data[ID] = actualLink
	str, err := json.Marshal(Data)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(pathData, str, 0777)
}
