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
INSERT INTO users.users(name , email , password , telephone_number , gender)
VALUES ( ?, ?, ?, ?, ? ) returning user_uuid;
`
	var rslt struct {
		UserUUID string `db:"user_uuid"`
	}
	if result := d.db.Raw(sql, attr.name, attr.email, attr.password, attr.telephoneNumber, attr.gender).Scan(&rslt); result.Error != nil {
		return "", result.Error
	}

	res, err := ParseUserUUID(rslt.UserUUID)
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
	if result := d.db.Exec(sql, attr.name, attr.email, attr.password, attr.telephoneNumber, attr.gender, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d userDataAccessor) delete(id UserUUID) error {
	sql := `
delete from users.users 
 where user_uuid= ?;
`
	if result := d.db.Exec(sql, id); result.Error != nil {
		return result.Error
	}
	return nil
}
