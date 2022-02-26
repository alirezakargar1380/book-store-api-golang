package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alirezakargar1380/book-store-api-golang/app/model"
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
