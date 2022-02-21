package routers

import (
	"github.com/alirezakargar1380/book-store-api-golang/app/controllers"
	"github.com/gorilla/mux"
)

type User struct {
	Num int
}

// type User struct {
// 	Name string
// }

type Profile struct {
	Name    string
	Hobbies []string
}

type Book struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var books []Book

func GetAllBooks(router *mux.Router) {
	router.HandleFunc("/get_json", controllers.Get_json).Methods("GET")
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")

	// 	// user := User{"hello"}
	// 	// js, err := json.Marshal(user)

	// 	// profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	// 	// js, err := json.Marshal(profile)

	// 	// m := map[string]User{
	// 	// 	// "o": 1,
	// 	// 	"one": {
	// 	// 		Num: 1,
	// 	// 	},
	// 	// 	"tow": {
	// 	// 		Num: 1,
	// 	// 	},
	// 	// }

	// 	// js, err := json.Marshal(m)

	// 	// if err != nil {
	// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	// 	return
	// 	// }
	// 	// json.NewEncoder(w).Encode(user)
	// 	// w.Write(js)

	// 	books = append(books, Book{
	// 		Id:   "1",
	// 		Name: "how we should sex with our partneradsasdf",
	// 	})

	// 	json.NewEncoder(w).Encode(books)

	// }).Methods("GET")
}
