package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	filePath := "file://./migration/ddl"
	postgresConnect := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"

	m, err := migrate.New(filePath, postgresConnect)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
