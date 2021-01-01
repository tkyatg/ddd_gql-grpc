package userqueryservice

import (
	"github.com/jinzhu/gorm"
)

type dataAccessor struct {
	db *gorm.DB
}

// NewDataAccessor function
func NewDataAccessor(db *gorm.DB) DataAccessor {
	return &dataAccessor{db}
}

func (d *dataAccessor) getUserByID(req getUserByIDRequest) (getUserByIDResponse, error) {
	sql := `
    SELECT id
         , name
         , email
         , password
         , telephoneNumber
         , gender
	  FROM users.users
     WHERE id = ?`
	var rslt struct {
		userUUID        string `db:"user_uuid"`
		name            string `db:"name"`
		email           string `db:"email"`
		password        string `db:"password"`
		telephoneNumber string `db:"telephoneNumber"`
		gender          int64  `db:"gender"`
	}

	d.db.Raw(sql, req.userUUID).Scan(&rslt)

	return getUserByIDResponse{
		userUUID:        rslt.userUUID,
		name:            rslt.name,
		email:           rslt.email,
		password:        rslt.password,
		telephoneNumber: rslt.telephoneNumber,
		gender:          rslt.gender,
	}, nil
}
