package main

import (
	"fmt"
	"net/http"
	"github.com/google/uuid"
)

type ItemsDetail struct {
	ID string 
	Name string `json:"name"`
}

var listItemDetails []ItemsDetail

func simpleHttpReqRes (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/items" {
		return http.NotFound(w, r)
	}

	switch r.Method {
	case "GET":
		itemJson, _ := json.Marshal(listItemDetails)

		w.write(itemJson)
	case "POST":
		reqBody, err := ioutil.ReadAll(r.body)
		if err != nil {
			panic(err)
		}

		var item ItemsDetail
		json.Unmarshal(reqBody, &item)

		item.ID = uuid.New().String()

		listItemDetails = append(listItemDetails, item)

		w.write([]byte("Data well received"))
	}
}

func main() {
	http.HanleFunc("/items", simpleHttpReqRes)
	http.ListenAndServe(":8080", nil)
}