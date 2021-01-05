package rpc

import (
	definition "github.com/takuya911/project-user-definition"
)

func (s *server) RegisterService(
	userQueryServer definition.UserQueryServiceServer,
	userCommandServer definition.UserCommandServiceServer,
) {
	definition.RegisterUserQueryServiceServer(s.rpc, userQueryServer)
	definition.RegisterUserCommandServiceServer(s.rpc, userCommandServer)
}
