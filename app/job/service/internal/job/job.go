package job

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewJobServer, NewOrderJobServer)

func NewJobServer() *Server {
	return &Server{}
}