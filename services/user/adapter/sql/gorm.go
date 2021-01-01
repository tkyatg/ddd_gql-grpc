package sql

import (
	"os"

	"github.com/jinzhu/gorm"
	// qg
	_ "github.com/lib/pq"
)

// NewGormConnect function
func NewGormConnect() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHostName := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connect := "postgres://" + user + ":" + pass + "@" + dbHostName + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	db, err := gorm.Open("postgres", connect)
	if err != nil {
		return nil, err
	}
	return db, nil
}
