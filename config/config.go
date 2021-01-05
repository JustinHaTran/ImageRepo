
package config

import (
	"time"

	"github.com/JustinHaTran/ImageRepo/ent"
	"github.com/facebook/ent/dialect/sql"
)

func Open() (*ent.Client, error) {
    drv, err := sql.Open("mysql", "root:022816@tcp(127.0.0.1)/ImageRepoDB")
    if err != nil {
        return nil, err
    }
    // Get the underlying sql.DB object of the driver.
    db := drv.DB()
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(100)
    db.SetConnMaxLifetime(time.Hour)
    return ent.NewClient(ent.Driver(drv)), nil
}
