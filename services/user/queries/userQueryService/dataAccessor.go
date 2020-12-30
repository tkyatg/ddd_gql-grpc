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
	  FROM users
     WHERE id = ?`
	var rslt struct {
		id              string `db:"id"`
		name            string `db:"name"`
		email           string `db:"email"`
		password        string `db:"password"`
		telephoneNumber string `db:"telephoneNumber"`
		gender          int64  `db:"gender"`
	}

	d.db.Raw(sql, req.id).Scan(&rslt)

	return getUserByIDResponse{
		id:              rslt.id,
		name:            rslt.name,
		email:           rslt.email,
		password:        rslt.password,
		telephoneNumber: rslt.telephoneNumber,
		gender:          rslt.gender,
	}, nil
}
