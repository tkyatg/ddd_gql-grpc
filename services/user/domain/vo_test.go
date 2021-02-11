package domain

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/takuya911/ddd_gql-grpc/services/user/shared"
)

func TestVoParseUserUUID(t *testing.T) {
	t.Parallel()
	if userUUID, err := ParseUserUUID(""); userUUID != "" && err != errors.New(shared.RequiredUserUUID) {
		t.Fatal(userUUID, err)
	}
	if userUUID, err := ParseUserUUID("1111111"); userUUID != "" && err != errors.New(shared.InvalidUUIDFormat) {
		t.Fatal(userUUID, err)
	}
	if userUUID, err := ParseUserUUID(uuid.New().String()); userUUID == "" || err != nil {
		t.Fatal(userUUID, err)
	}
}

func TestVoParseUserName(t *testing.T) {
	if userName, err := parseUserName(""); userName != "" && err != errors.New(shared.RequiredUserName) {
		t.Fatal(userName, err)
	}
	if userName, err := parseUserName("test"); userName == "" || err != nil {
		t.Fatal(userName, err)
	}
}

func TestVoParsePassword(t *testing.T) {
	t.Parallel()
	if password, err := parsePassword(""); password != "" && err != errors.New(shared.RequiredPassword) {
		t.Fatal(password, err)
	}
	if password, err := parsePassword("111111"); password != "" && err != errors.New(shared.InvalidPasswordLength) {
		t.Fatal(password, err)
	}
	if password, err := parsePassword("gdpspvls"); password == "" || err != nil {
		t.Fatal(password, err)
	}
}

func TestVoParseEmail(t *testing.T) {
	t.Parallel()
	if email, err := parseEmail(""); email != "" && err != errors.New(shared.RequiredEmail) {
		t.Fatal(email, err)
	}
	if email, err := parseEmail("e@"); email != "" && err != errors.New(shared.InvalidEmailLength) {
		t.Fatal(email, err)
	}
	if email, err := parseEmail("testgmail.com"); email != "" && err != errors.New(shared.InvalidEmailFormat) {
		t.Fatal(email, err)
	}
	if email, err := parsePassword("test@gmail.com"); email == "" || err != nil {
		t.Fatal(email, err)
	}
}

func TestVoParseTelephoneNumber(t *testing.T) {
	t.Parallel()
	if tel, err := parseTelephoneNumber(""); tel != "" && err != errors.New(shared.RequiredTelephoneNumber) {
		t.Fatal(tel, err)
	}
	if tel, err := parseTelephoneNumber("verbesbest"); tel != "" && err != errors.New(shared.InvalidTelephoneNumberFormat) {
		t.Fatal(tel, err)
	}
	if tel, err := parseTelephoneNumber("090-4554-6456"); tel != "" && err != errors.New(shared.InvalidTelephoneNumberFormat) {
		t.Fatal(tel, err)
	}
	if tel, err := parseTelephoneNumber("09034543453"); tel == "" || err != nil {
		t.Fatal(tel, err)
	}
}
func TestVoParseGender(t *testing.T) {
	t.Parallel()
	if g, err := parseGender(0); g != 0 && err != errors.New(shared.InvalidGenderFormat) {
		t.Fatal(g, err)
	}
	if g, err := parseGender(4); g != 0 && err != errors.New(shared.InvalidGenderFormat) {
		t.Fatal(g, err)
	}
	if g, err := parseGender(1); g == 0 || err != nil {
		t.Fatal(g, err)
	}
}

func TestVoNewUserAttributes(t *testing.T) {
	t.Parallel()
	res, err := NewUserAttributes("name", "password", "test@gmail.com", "09009009090", 1)
	if err != nil {
		t.Fatal(err)
	}
	expect := UserAttributes{
		name:            UserName("name"),
		password:        Password("password"),
		email:           Email("test@gmail.com"),
		telephoneNumber: TelephoneNumber("09009009090"),
		gender:          Gender(1),
	}
	result := UserAttributes{
		name:            res.name,
		password:        res.password,
		email:           res.email,
		telephoneNumber: res.telephoneNumber,
		gender:          res.gender,
	}
	if result != expect {
		t.Fatal(res)
	}
}
