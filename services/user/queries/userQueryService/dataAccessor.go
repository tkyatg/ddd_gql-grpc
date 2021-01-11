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
		UserUUID        string `db:"user_uuid"`
		Name            string `db:"name"`
		Email           string `db:"email"`
		Password        string `db:"password"`
		TelephoneNumber string `db:"telephone_number"`
		Gender          int64  `db:"gender"`
	}

	d.db.Raw(sql, req.userUUID).Scan(&rslt)

	return getByIDResponse{
		userUUID:        rslt.UserUUID,
		name:            rslt.Name,
		email:           rslt.Email,
		password:        rslt.Password,
		telephoneNumber: rslt.TelephoneNumber,
		gender:          rslt.Gender,
	}, nil
}
