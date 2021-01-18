module github.com/takuya911/project-services/services/auth

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang-migrate/migrate/v4 v4.14.1 // indirect
	github.com/golang/mock v1.4.4
	github.com/google/go-cmp v0.5.1
	github.com/google/uuid v1.1.2
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.8.0
	github.com/takuya911/project-auth-definition v0.0.0-20210117101525-89d87b5f36f0
	github.com/takuya911/project-services/services/user v0.0.0-20210118110920-853d8a1da2bc
	github.com/takuya911/project-user-definition v0.0.0-20210117110601-750ed00fd461
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	google.golang.org/grpc v1.34.0

)
