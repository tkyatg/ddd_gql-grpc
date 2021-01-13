package sql

import (
	"github.com/jinzhu/gorm"
	// qg
	_ "github.com/lib/pq"
)

// NewGormConnect func
func NewGormConnect(dbUser string, dbPassword string, dbName string, dbHost string, dbport string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "postgres://"+dbUser+":"+dbPassword+"@"+dbHost+":"+dbport+"/"+dbName+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
