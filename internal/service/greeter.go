package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	v1 "github.com/toheart/go-structure/api/helloworld/v1"
	"github.com/toheart/go-structure/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer
	uc *biz.GreeterUsecase

	log *log.Helper
}

func (s *GreeterService) RegisterGrpc(svr *grpc.Server) {
	v1.RegisterGreeterServer(svr, s)
}

func (s *GreeterService) RegisterHttp(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return v1.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
