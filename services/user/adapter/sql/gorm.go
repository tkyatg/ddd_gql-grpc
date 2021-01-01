package sql

import (
	"os"

	"github.com/jinzhu/gorm"
	// qg
	_ "github.com/lib/pq"
)

const defaultDbUser = "postgres"
const defaultDbPass = "password"
const defaultDbName = "postgres"
const defaultDbHost = "postgres"
const defaultDbPort = "5432"

// NewGormConnect function
func NewGormConnect() (*gorm.DB, error) {
	user := os.Getenv("DB_USER")
	if user == "" {
		user = defaultDbUser
	}
	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = defaultDbPass
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = defaultDbName
	}
	dbHostName := os.Getenv("DB_HOST")
	if dbHostName == "" {
		dbHostName = defaultDbHost
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = defaultDbPort
	}

	connect := "postgres://" + user + ":" + pass + "@" + dbHostName + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	db, err := gorm.Open("postgres", connect)
	if err != nil {
		return nil, err
	}

	return db, nil
}
