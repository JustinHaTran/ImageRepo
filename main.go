
package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"net/http"

	"github.com/JustinHaTran/ImageRepo/ent"
	"github.com/JustinHaTran/ImageRepo/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"contrib.go.opencensus.io/integrations/ocsql"
	"github.com/go-sql-driver/mysql"
	entsql "github.com/facebook/ent/dialect/sql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
