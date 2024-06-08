package iface

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

/**
@file:
@author: levi.Tang
@time: 2024/6/8 16:30
@description:
**/

type InitGrpcHttp interface {
	RegisterHttp(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
	RegisterGrpc(svr *grpc.Server)
}
