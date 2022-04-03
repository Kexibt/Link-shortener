package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var data map[int]string

const (
	pathData = "data/data.json"
)

func getDataNonDB() {
	if checkCreateData() {
		data = make(map[int]string)
		return
	}

	file, err := ioutil.ReadFile(pathData)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
}

func checkCreateData() bool {
	exdir, err := exists(pathData[:strings.Index(pathData, "/")])
	if err != nil {
		log.Fatal(err)
	}

	if !exdir {
		err = os.Mkdir(pathData[:strings.Index(pathData, "/")], 0777)
		if err != nil {
			panic(err)
		}
	}

	ex, err := exists(pathData)
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

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func findActualNonDB(ID int) string {
	getDataNonDB()
	link, exist := data[ID]
	if !exist {
		return "error"
	}

	return link
}

func getLastNonDB() int {
	getDataNonDB()
	return len(data)
}

func wasHereNonDB(actualLink string) int {
	getDataNonDB()
	for key, value := range data {
		if value == actualLink {
			return key
		}
	}

	return -1
}

func insertNonDB(ID int, actualLink string) {
	getDataNonDB()
	data[ID] = actualLink
	str, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(pathData, str, 0777)
}
