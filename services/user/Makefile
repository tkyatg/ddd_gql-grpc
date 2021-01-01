gen-mock:
	mockgen -source queries/userQueryService/usecase.go -destination  queries/userQueryService/usecase_mock.go -package userqueryservice
test:
	go test -race -covermode=atomic ./... -test.v
migrationfile:
	migrate create -ext sql -dir migration/ddl -seq create_$(table)_table
migrate:
	go run ./migration/migrate.go