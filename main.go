package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	urlHost  = "http://127.0.0.1"
	portHost = ":3000"
)

type Request struct {
	Key  string `json:"key"`
	Link string `json:"link"`
}

func index(w http.ResponseWriter, r *http.Request) {
	var request Request

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()
	json.Unmarshal(body, &request)

	switch r.Method {
	case "GET":
		if request.Link == "" {
			request.Link = r.URL.Path
		}

		ind := strings.LastIndex(request.Link, "/")
		shortLink := request.Link[ind+1:]

		actualLink := ""
		if *dbUsage {
			actualLink = findActualDB(convertToID(shortLink))
		} else {
			actualLink = findActualNonDB(convertToID(shortLink))
		}

		response, err := json.Marshal(Request{Link: actualLink})
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "json")
		w.Write(response)

	case "POST":
		actualLink := request.Link
		ind := 0

		if *dbUsage {
			ind = wasHereDB(actualLink)
			if ind <= 0 {
				ind = getLastDB() + 1
				insertDB(ind, actualLink)
			}
		} else {
			ind = wasHereNonDB(actualLink)
			if ind <= 0 {
				ind = getLastNonDB() + 1
				insertNonDB(ind, actualLink)
			}
		}
		shortLink := urlHost + portHost + "/" + convertToShort(ind)

		response, err := json.Marshal(Request{Link: shortLink})
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "json")
		w.Write(response)

	}
}

var dbUsage *bool

func main() {
	dbUsage = flag.Bool("db", false, "database usage")
	flag.Parse()

	// в задании не указано, нужно ли дублировать записи в обе базы данных,
	// поэтому я не дублировал)
	// P.S. это несложно сделать, нужно в каждом if-else блоке, где проверяется
	// dbUsage, при обработке POST-запроса делать запись в каждую бд
	if !*dbUsage {
		getDataNonDB()
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	http.ListenAndServe(portHost, mux)
}
