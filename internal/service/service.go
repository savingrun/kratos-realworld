package service

import (
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.UserUseCase
}

func NewRealWorldService(uc *biz.UserUseCase) *RealWorldService {
	return &RealWorldService{uc: uc}
}
