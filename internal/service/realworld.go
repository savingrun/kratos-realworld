package service

import (
	"context"

	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"
)

// RealWorldService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.GreeterUsecase
}

// NewRealWorldService new a greeter service.
func NewRealWorldService(uc *biz.GreeterUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}

// SayHello implements realworld.GreeterServer.
func (s *RealWorldService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.UserReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.User.GetEmail()})
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{User: &v1.UserReply_Data{Email: "Hello" + g.Hello}}, nil
}
