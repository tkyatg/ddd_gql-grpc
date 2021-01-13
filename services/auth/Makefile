test:
	go test -race -covermode=atomic ./... -test.v
migrationfile:
	migrate create -ext sql -dir migration/ddl -seq create_$(table)_table
migrate:
	go run ./migration/migrate.go