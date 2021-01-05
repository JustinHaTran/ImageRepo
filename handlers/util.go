package handlers

import (
	"fmt"
	"net/http"

	//"github.com/facebook/ent"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
)

/*
User ApiHandling
*/

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO 1");
	// CreateBook := &models.Book{}
	// utils.ParseBody(r, CreateBook)
	// b:= CreateBook.CreateBook()
	// res,_ := json.Marshal(b)
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO 2");
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO 3");
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HELLO 4");
}