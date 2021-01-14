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

func (d authenticationDataAccessor) login(email Email) (UserUUID, Password, error) {
	sql := `
select user_uuid
     , password
  from users.users
 where email = ?
`
	var rslt struct {
		uuid     string `db:"user_uuid"`
		password string `db:"password"`
	}
	if result := d.db.Raw(sql, string(email)).Scan(&rslt); result.Error != nil {
		return "", "", result.Error
	}
	return UserUUID(rslt.uuid), Password(rslt.password), nil
}
