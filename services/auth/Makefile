gen-mock:
	mockgen -source queries/authQueryService/usecase.go -destination  queries/authQueryService/usecase_mock.go -package authqueryservice
test:
	go test -race -covermode=atomic ./... -test.v
migrationfile:
	migrate create -ext sql -dir migration/ddl -seq create_$(table)_table
migrate:
	go run ./migration/migrate.go
