package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderBiz = wire.NewSet(NewGreeterUsecase)
