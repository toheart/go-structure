package server

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
)

// ProviderServer is server providers.
var ProviderServer = wire.NewSet(NewHTTPServer, NewGRPCServer)

func NewServerList(http *HttpServer, grpc *GrpcServer) []transport.Server {
	return []transport.Server{
		http,
		grpc,
	}
}
