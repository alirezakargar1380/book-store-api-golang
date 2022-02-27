package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/alirezakargar1380/book-store-api-golang/app/model"
	"github.com/alirezakargar1380/book-store-api-golang/app/utils"
)

var NewBool model.Book

type Response struct {
	Status bool
	Data   []model.Book
}

func All_data(w http.ResponseWriter, r *http.Request) {
	newBooks := model.GetAllBooks()
	dd := Response{
		Status: true,
		Data:   newBooks,
	}

	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(dd)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &model.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
