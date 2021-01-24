package userqueryservice

import (
	"time"

	"github.com/jinzhu/gorm"
)

type dataAccessor struct {
	db *gorm.DB
}

// NewDataAccessor function
func NewDataAccessor(db *gorm.DB) DataAccessor {
	return &dataAccessor{db}
}

func (d *dataAccessor) getByID(req getByIDRequest) (getByIDResponse, error) {
	sql := `
SELECT user_uuid
     , name
     , email
     , password
     , telephone_number
	 , gender
	 , created_at
	 , updated_at
FROM users.users
WHERE user_uuid = ?
`
	var rslt struct {
		UserUUID        string    `db:"user_uuid"`
		Name            string    `db:"name"`
		Email           string    `db:"email"`
		Password        string    `db:"password"`
		TelephoneNumber string    `db:"telephone_number"`
		Gender          int64     `db:"gender"`
		CreatedAt       time.Time `db:"created_at"`
		UpdatedAt       time.Time `db:"updated_at"`
	}

	d.db.Raw(sql, req.userUUID).Scan(&rslt)

	return getByIDResponse{
		userUUID:        rslt.UserUUID,
		name:            rslt.Name,
		email:           rslt.Email,
		password:        rslt.Password,
		telephoneNumber: rslt.TelephoneNumber,
		gender:          rslt.Gender,
		createdAt:       rslt.CreatedAt,
		updatedAt:       rslt.UpdatedAt,
	}, nil
}

func (d *dataAccessor) getByEmail(req getByEmailAndPasswordRequest) (getByEmailAndPasswordResponse, error) {
	sql := `
SELECT user_uuid
     , name
     , email
     , password
     , telephone_number
	 , gender
	 , created_at
	 , updated_at
FROM users.users
WHERE email = ?
`
	var rslt struct {
		UserUUID        string    `db:"user_uuid"`
		Name            string    `db:"name"`
		Email           string    `db:"email"`
		Password        string    `db:"password"`
		TelephoneNumber string    `db:"telephone_number"`
		Gender          int64     `db:"gender"`
		CreatedAt       time.Time `db:"created_at"`
		UpdatedAt       time.Time `db:"updated_at"`
	}

	d.db.Raw(sql, req.email).Scan(&rslt)

	return getByEmailAndPasswordResponse{
		userUUID:        rslt.UserUUID,
		name:            rslt.Name,
		email:           rslt.Email,
		password:        rslt.Password,
		telephoneNumber: rslt.TelephoneNumber,
		gender:          rslt.Gender,
		createdAt:       rslt.CreatedAt,
		updatedAt:       rslt.UpdatedAt,
	}, nil
}
