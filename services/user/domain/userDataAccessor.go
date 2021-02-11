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
INSERT INTO users.users(name , email , password , telephone_number , gender, created_at)
VALUES ( ?, ?, ?, ?, ? ,now() ) returning user_uuid
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
	 , updated_at = now()
 where user_uuid = ?
`
	if result := d.db.Exec(sql, attr.name, attr.email, attr.password, attr.telephoneNumber, attr.gender, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d userDataAccessor) delete(id UserUUID) error {
	sql := `
delete from users.users 
 where user_uuid= ?
`
	if result := d.db.Exec(sql, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d userDataAccessor) emailAlreadyUsedCreate(email Email) (bool, error) {
	sql := `
select exists (
       select user_uuid from users.users
        where email= ?
)
`
	var rslt struct {
		Exist bool `db:"exist"`
	}
	if result := d.db.Raw(sql, email).Scan(&rslt); result.Error != nil {
		return true, result.Error
	}
	return rslt.Exist, nil
}

func (d userDataAccessor) emailAlreadyUsedUpdate(id UserUUID, email Email) (bool, error) {
	sql := `
select exists (
       select user_uuid from users.users
        where email= ?
          and user_uuid != ?
)
`
	var rslt struct {
		Exist bool `db:"exist"`
	}
	if result := d.db.Raw(sql, email, id).Scan(&rslt); result.Error != nil {
		return true, result.Error
	}
	return rslt.Exist, nil
}
