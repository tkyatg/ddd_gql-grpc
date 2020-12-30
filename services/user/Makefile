gen-mock:
	mockgen -source queries/userQueryService/usecase.go -destination  queries/userQueryService/usecase_mock.go -package userqueryservice
test:
	go test -race -covermode=atomic ./... -test.v