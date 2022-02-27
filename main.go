package main

import (
	"fmt"
	"net/http"

	"github.com/alirezakargar1380/book-store-api-golang/app/routers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("this is API")
	r := mux.NewRouter()
	routers.GetAllBooks(r)
	routers.TestRouters(r)

	http.ListenAndServe(":8000", r)
}
