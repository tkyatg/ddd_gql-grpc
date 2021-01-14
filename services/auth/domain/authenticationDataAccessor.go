package domain

import (
	"github.com/jinzhu/gorm"
)

type (
	authenticationDataAccessor struct {
		db *gorm.DB
	}
)

// NewAuthenticationDataAccessor fun
func NewAuthenticationDataAccessor(
	db *gorm.DB,
) AuthenticationDataAccessor {
	return &authenticationDataAccessor{db}
}

func (d authenticationDataAccessor) login(email Email, password Password) (UserUUID, error) {
	sql := `
select user_uuid
  from users.users
 where email = ?
   and password = ?

`
	var rslt struct {
		uuid string `db:"user_uuid"`
	}
	if result := d.db.Raw(sql, string(email), string(password)).Scan(&rslt); result.Error != nil {
		return "", result.Error
	}
	return UserUUID(rslt.uuid), nil
}
