package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alirezakargar1380/book-store-api-golang/app/model"
	"github.com/alirezakargar1380/book-store-api-golang/app/utils"
	"github.com/gorilla/mux"
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

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookId string = vars["book"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		panic(err)
	}
	bookDetail, _ := model.GetBookById(ID)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
