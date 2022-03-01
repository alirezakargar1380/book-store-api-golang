package model

import (
	"github.com/alirezakargar1380/book-store-api-golang/app/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name  string `gorm:""json:"name"`
	Genre string `json:"Genre"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

// func DeleteBookById(ID) Book {
// 	var book Book
// 	db.Where("ID=?", )
// }
