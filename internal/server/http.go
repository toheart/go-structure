package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/toheart/go-structure/internal/conf"
	"github.com/toheart/go-structure/internal/server/iface"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"

	"github.com/go-kratos/kratos/v2/log"
)

var _ transport.Server = (*HttpServer)(nil)

type HttpServer struct {
	server *http.Server
	log    *log.Helper
}

func (h *HttpServer) Start(ctx context.Context) error {
	h.log.Infof("start http gateway server: %s", h.server.Addr)
	return h.server.ListenAndServe()
}

func (h *HttpServer) Stop(ctx context.Context) error {
	h.log.Infof("Shutting down the http gateway server")
	if err := h.server.Shutdown(ctx); err != nil {
		h.log.Errorf("Failed to shutdown http gateway server: %v", err)
	}
	return nil
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, services ...iface.InitGrpcHttp) *HttpServer {
	h := &HttpServer{
		log: log.NewHelper(log.With(logger, "module", "http")),
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	mux := runtime.NewServeMux()
	h.server = &http.Server{
		Addr:    c.Http.Addr,
		Handler: mux,
	}
	for _, item := range services {
		if err := item.RegisterHttp(mux, c.Grpc.Addr, opts); err != nil {
			panic(err)
		}
	}
	// 添加Prometheus 接口
	err := mux.HandlePath("GET", "/metrics", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		promhttp.Handler().ServeHTTP(w, r)
	})
	if err != nil {
		panic(err)
	}
	return h
}
