package rpc

import (
	"fmt"
	"net"

	definition "github.com/takuya911/project-user-definition"
	"google.golang.org/grpc"
)

type (
	server struct {
		port string
		rpc  *grpc.Server
	}
	// Server interface
	Server interface {
		Serve() error
		RegisterService(
			userQueryServer definition.UserQueryServiceServer,
			userCommandServer definition.UserCommandServiceServer,
		)
	}
)

// NewServer はインスタンスを生成します
func NewServer(port string) (Server, error) {
	s := &server{port, grpc.NewServer()}

	return s, nil
}

func (t *server) Serve() error {
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%s", t.port))
	if err != nil {
		return err
	}
	return t.rpc.Serve(listenPort)
}
