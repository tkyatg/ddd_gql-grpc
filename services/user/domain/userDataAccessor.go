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

func (d userDataAccessor) create(attr *UserAttributes) error {
	sql := `
insert into users.users
     ( name
     , email
     , password
     , telephone_number
     , gender )
values ( ?
       , ?
       , ?
       , ?
       , ? );
`
	if result := d.db.Exec(sql, attr.name, attr.email, attr.password, attr.telephoneNumber, attr.gender); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d userDataAccessor) update(id UserUUID, attr *UserAttributes) error {
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
delete users.users 
 where user_uuid= ?;
`
	if result := d.db.Exec(sql, id); result.Error != nil {
		return result.Error
	}
	return nil
}
