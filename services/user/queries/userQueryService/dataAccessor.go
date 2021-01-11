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

func (d *dataAccessor) getByID(req getByIDRequest) (getByIDResponse, error) {
	sql := `
SELECT user_uuid
     , name
     , email
     , password
     , telephone_number
     , gender
FROM users.users
WHERE user_uuid = ?
`
	var rslt struct {
		userUUID        string `db:"user_uuid"`
		name            string `db:"name"`
		email           string `db:"email"`
		password        string `db:"password"`
		telephoneNumber string `db:"telephone_number"`
		gender          int64  `db:"gender"`
	}

	d.db.Raw(sql, req.userUUID).Scan(&rslt)

	return getByIDResponse{
		userUUID:        rslt.userUUID,
		name:            rslt.name,
		email:           rslt.email,
		password:        rslt.password,
		telephoneNumber: rslt.telephoneNumber,
		gender:          rslt.gender,
	}, nil
}
