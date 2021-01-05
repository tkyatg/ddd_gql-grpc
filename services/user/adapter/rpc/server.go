package rpc

import (
	"fmt"
	"net"

	"github.com/jinzhu/gorm"
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
	}
)

// NewServer はインスタンスを生成します
func NewServer(port string, dbConnection *gorm.DB) (Server, error) {
	s := &server{port, grpc.NewServer()}
	s.registerServices(dbConnection)
	return s, nil
}

func (s *server) Serve() error {
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return err
	}
	return s.rpc.Serve(listenPort)
}
