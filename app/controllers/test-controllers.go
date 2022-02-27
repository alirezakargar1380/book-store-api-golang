package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alirezakargar1380/book-store-api-golang/app/model"
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

var NewBool model.Book

func All_data(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// user := User{"io"}
	// js, err := json.Marshal(user)
	// if err != nil {
	// 	return
	// }
	// w.Write(js)
	newBooks := model.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func TestJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	m := map[string]types.TestJson{
		"one": {
			Num: 1,
		},
		"tow": {
			Num: 1,
		},
	}

	fmt.Println(m)
	js, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(js)
	w.Write(js)
}
