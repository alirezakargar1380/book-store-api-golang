package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alirezakargar1380/book-store-api-golang/app/types"
)

type User struct {
	Name string
}

func Get_json(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := User{"alireza"}
	js, err := json.Marshal(user)
	if err != nil {
		return
	}
	w.Write(js)
	// json.NewEncoder(w).Encode(js)
	fmt.Println("im controller")
}

func TestJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	m := map[string]types.TestJson{
		"a": {
			Num: 1,
		},
		"b": {
			Num: 1,
		},
	}

	// ADD DATA IN JSON
	m["c"] = types.TestJson{
		Num: 465,
	}

	// UPDATE DATA IN JSON
	m["a"] = types.TestJson{
		Num: 2,
	}

	js, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	w.Write(js)
}
