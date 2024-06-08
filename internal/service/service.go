package service

import (
	"github.com/google/wire"
	"github.com/toheart/go-structure/internal/server/iface"
)

// ProviderServices is service providers.
var ProviderServices = wire.NewSet(NewGreeterService, NewHttpServiceList)

func NewHttpServiceList(gs *GreeterService) []iface.InitGrpcHttp {
	return []iface.InitGrpcHttp{
		gs,
	}
}
