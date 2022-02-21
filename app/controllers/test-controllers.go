package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
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
