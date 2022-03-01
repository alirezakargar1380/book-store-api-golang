package model

import (
	"fmt"

	"github.com/alirezakargar1380/book-store-api-golang/app/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name  string `gorm:""json:"name"`
	Genre string `json:"genre"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	fmt.Println("hello world")
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// when we use *Book in function parameters
// we want the address
// in other side when we use *Book like fmt.PrintLn(*Book)
// we want the storted data in the address

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
