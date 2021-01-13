
package handlers

import (
	"log"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JustinHaTran/ImageRepo/models"
	"github.com/JustinHaTran/ImageRepo/ent/image"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

/*
User ApiHandling
*/


func CreateImage(w http.ResponseWriter, r *http.Request) {
	
	var u models.Image

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		http.Error(w,"invalid payload",http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	image, err := client.Image.    	
		Create().               	
		SetTitle(u.Title).       		
		SetDescription(u.Description).  
		SetFileLocation(u.FileLocation).
		SetPrice(u.Price).				
		SetPublic(u.Public).			
		Save(ctx)            	    

	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	log.Println("image was created: ", image)
}

func GetImageById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	imageId, err := strconv.Atoi(vars["imageId"])
	if err != nil {
		log.Fatalf("url var is invalid: %v", err)
	}

	u, err := client.Image.
        Query().
        Where(image.IDEQ(imageId)).
        // `Only` fails if no user found,
        // or more than 1 user returned.
		Only(ctx)
	
	if err != nil {
		log.Fatalf("failed to find user: %v", err)
	}

	respondWithJSON(w, http.StatusOK, u)
}

func UpdateImage(w http.ResponseWriter, r *http.Request) {

	var u models.Image
	vars := mux.Vars(r)
	imageId, err := strconv.Atoi(vars["imageId"])
	if err != nil {
		log.Fatalf("url var is invalid: %v", err)
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		http.Error(w,"invalid payload",http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	image, err := client.Image.
		UpdateOneID(imageId).        // Pet update builder.
		SetTitle(u.Title).       		
		SetDescription(u.Description).  
		SetFileLocation(u.FileLocation).
		SetPrice(u.Price).				
		SetPublic(u.Public).
		Save(ctx)
	
	if err != nil {
		log.Fatalf("failed to find image: %v", err)
	}

	respondWithJSON(w, http.StatusOK, image)
}
