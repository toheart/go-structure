package cmd

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/toheart/go-structure/internal/conf"
	"os"
)

/**
@file:
@author: levi.Tang
@time: 2024/6/8 14:42
@description:
**/

type InitApp struct {
	c config.Config
}

func NewInitApp() *InitApp {
	i := &InitApp{}
	return i
}

func (i *InitApp) InitConfig() *conf.Bootstrap {
	i.c = config.New(
		config.WithSource(
			file.NewSource(cfgFile),
		),
	)
	if err := i.c.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := i.c.Scan(&bc); err != nil {
		panic(err)
	}
	// TODO 可通过watch监听

	return &bc
}

func (i *InitApp) InitLogger() log.Logger {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", AppName,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	return logger
}

func (i *InitApp) CloseConfig() {
	i.c.Close()
}
