
package main

import (
	"log"
	"fmt"
	// "context"
	// "database/sql"
	// "database/sql/driver"
	"net/http"

	//"github.com/JustinHaTran/ImageRepo/ent"
	"github.com/JustinHaTran/ImageRepo/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "contrib.go.opencensus.io/integrations/ocsql"
	// "github.com/go-sql-driver/mysql"
	// entsql "github.com/facebook/ent/dialect/sql"
	"github.com/JustinHaTran/ImageRepo/handlers"

)

func main() {
	
	handlers.InitClient();

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	fmt.Println("Now listening on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
