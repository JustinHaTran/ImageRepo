package handlers

import (
	"encoding/json"
	"net/http"
	"context"
	"log"

	"github.com/JustinHaTran/ImageRepo/ent"
	_ "github.com/go-sql-driver/mysql"
)

var client *ent.Client
var ctx context.Context

func InitClient() {
	cl, err := ent.Open("mysql", "root:022816@tcp(127.0.0.1)/ImageRepoDB")
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }
	ct := context.Background()

	ctx = ct
	client = cl
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}