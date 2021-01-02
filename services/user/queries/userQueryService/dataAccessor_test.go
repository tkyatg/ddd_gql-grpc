package userqueryservice

// import (
// 	"log"
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// 	"github.com/google/uuid"
// 	"github.com/takuya911/project-services/services/user/adapter/sql"
// )

// func TestDataAccessorGetUserByID(t *testing.T) {
// 	dbConn, err := sql.NewGormConnect()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer dbConn.Close()
// 	dataAccessor := NewDataAccessor(dbConn)

// 	userUUID, _ := uuid.NewUUID()
// 	sql := `
// 	insert into users.users(user_uuid, name, email, password, telephone_number, gender)
// 	values(?,'name','email','password','telephone_number',1)`

// 	dbConn.Exec(sql, userUUID.String())

// 	res, err := dataAccessor.getUserByID(getUserByIDRequest{
// 		userUUID: "id",
// 	})

// 	opts := cmp.Options{}
// 	if diff := cmp.Diff(getUserByIDResponse{
// 		userUUID:        "id",
// 		name:            "name",
// 		email:           "email",
// 		password:        "password",
// 		telephoneNumber: "telephoneNumber",
// 		gender:          1,
// 	}, res, opts); diff != "" {
// 		t.Fatal(diff)
// 	}

// }
