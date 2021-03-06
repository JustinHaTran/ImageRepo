
package handlers

import (
	"log"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JustinHaTran/ImageRepo/models"
	"github.com/JustinHaTran/ImageRepo/ent/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

/*
User ApiHandling
*/


func CreateUser(w http.ResponseWriter, r *http.Request) {
	
	var u models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		http.Error(w,"invalid payload",http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, err := client.User.    	// UserClient.
		Create().               	// User create builder.
		SetName(u.Name).       		// Set field values.
		SetAge(u.Age).    			// Set field values.
		Save(ctx)            	    // Create and return.

	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	log.Println("user was created: ", user)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Fatalf("url var is invalid: %v", err)
	}

	u, err := client.User.
        Query().
        Where(user.IDEQ(userId)).
        // `Only` fails if no user found,
        // or more than 1 user returned.
		Only(ctx)
	
	if err != nil {
		log.Fatalf("failed to find user: %v", err)
	}

	respondWithJSON(w, http.StatusOK, u)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var u models.User
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		log.Fatalf("url var is invalid: %v", err)
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		http.Error(w,"invalid payload",http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, err := client.User.
		UpdateOneID(userId).        // Pet update builder.
		SetName(u.Name).       // Set field name.
		SetAge(u.Age).      // Set unique edge, using id.
		Save(ctx)
	
	if err != nil {
		log.Fatalf("failed to find user: %v", err)
	}

	respondWithJSON(w, http.StatusOK, user)
}
