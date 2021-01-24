DB_USER=postgres
DB_HOST=127.0.0.1:5432
DB_NAME=postgres
DB_PASSWORD=password

gen-mock:
	mockgen -source queries/userQueryService/usecase.go -destination  queries/userQueryService/usecase_mock.go -package userqueryservice
	mockgen -source commands/userCommandService/usecase.go -destination  commands/userCommandService/usecase_mock.go -package usercommandservice
	mockgen -source domain/userRepository.go -destination  domain/userRepository_mock.go -package domain
test:
	go test -race -covermode=atomic ./... -test.v
migrationfile:
	migrate create -ext sql -dir migration/ddl -seq create_$(table)_table
migrate:
	go run ./migration/migrate.go
