package routers

import (
	"github.com/alirezakargar1380/book-store-api-golang/app/controllers"
	"github.com/gorilla/mux"
)

func TestRouters(router *mux.Router) {
	router.HandleFunc("/get_json", controllers.Get_json).Methods("GET")
	router.HandleFunc("/test_json", controllers.TestJson).Methods("GET")
}
