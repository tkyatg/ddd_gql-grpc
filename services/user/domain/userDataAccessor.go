package domain

import (
	"github.com/jinzhu/gorm"
)

type (
	userDataAccessor struct {
		db *gorm.DB
	}
)

// NewUserDataAccessor fun
func NewUserDataAccessor(
	db *gorm.DB,
) UserDataAccessor {
	return &userDataAccessor{db}
}

func (d userDataAccessor) create(attr UserAttributes) (UserUUID, error) {
	sql := `
insert into users.users
     ( name
     , email
     , password
     , telephone_number
     , gender )
values
     ( ?
     , ?
     , ?
     , ?
	 , ? )
RETURNING user_uuid;
`
	var rslt struct {
		userUUID string `db:"user_uuid"`
	}
	if result := d.db.Raw(sql, attr.name, attr.email, attr.password, attr.telephoneNumber, attr.gender).Scan(&rslt); result.Error != nil {
		return "", result.Error
	}
	res, err := ParseUserUUID(rslt.userUUID)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (d userDataAccessor) update(id UserUUID, attr UserAttributes) error {
	sql := `
update users.users 
   set name = ?
     , email = ?
     , password = ?
     , telephone_number = ?
     , gender = ?
 where user_uuid = ?;
`
	// ここのresult何が返ってくるか気になる
	if result := d.db.Exec(sql, attr.name, attr.email, attr.password, attr.telephoneNumber, attr.gender, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d userDataAccessor) delete(id UserUUID) error {
	sql := `
delete users.users 
 where user_uuid= ?;
`
	// ここのresult何が返ってくるか気になる
	if result := d.db.Exec(sql, id); result.Error != nil {
		return result.Error
	}
	return nil
}
