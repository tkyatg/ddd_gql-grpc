// Package xogen contains the types for schema 'users'.
package xogen

// Code generated by xo. DO NOT EDIT.

import (
	"time"

	acallsql "github.com/acall-inc/acall-sql/v2"
	"github.com/google/uuid"
)

// UserInsertArgs represents a row from 'users'.
type UserInsertArgs struct {
	UserUUID        uuid.UUID `db:"user_uuid"`        // user_uuid
	Name            string    `db:"name"`             // name
	Email           string    `db:"email"`            // email
	Password        string    `db:"password"`         // password
	TelephoneNumber string    `db:"telephone_number"` // telephone_number
	Gender          int       `db:"gender"`           // gender
	CreatedAt       time.Time `db:"created_at"`       // created_at
	UpdatedAt       time.Time `db:"updated_at"`       // updated_at
}

// User represents a row from 'users'.
type User struct {
	UserUUID        uuid.UUID `db:"user_uuid"`        // user_uuid
	Name            string    `db:"name"`             // name
	Email           string    `db:"email"`            // email
	Password        string    `db:"password"`         // password
	TelephoneNumber string    `db:"telephone_number"` // telephone_number
	Gender          int       `db:"gender"`           // gender
	CreatedAt       time.Time `db:"created_at"`       // created_at
	UpdatedAt       time.Time `db:"updated_at"`       // updated_at
}

// InsertUser is insert model to users
func (t *UserInsertArgs) Insert(dbAccessor acallsql.DBAccessor) (*User, error) {
	const sql = `
		INSERT INTO users.users (
			"user_uuid"
			,"name"
			,"email"
			,"password"
			,"telephone_number"
			,"gender"
			,"created_at"
			,"updated_at"
		) VALUES (
			:user_uuid
			,:name
			,:email
			,:password
			,:telephone_number
			,:gender
			,:created_at
			,:updated_at
		)
		RETURNING
			"user_uuid"
			,"name"
			,"email"
			,"password"
			,"telephone_number"
			,"gender"
			,"created_at"
			,"updated_at"
	`
	var obj User
	if err := dbAccessor.ExecuteCallback(sql, t, func(rows acallsql.DBRows) error {
		for rows.Next() {
			rows.StructScan(&obj)
			return nil
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &obj, nil
}