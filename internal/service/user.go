package service

import (
	"context"

	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, request *v1.LoginRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{User: &v1.UserReply_Data{Email: "Saving"}}, nil
}

func (s *RealWorldService) Registration(ctx context.Context, request *v1.RegistrationRequest) (reply *v1.UserReply, err error) {
	u, err := s.uc.Register(ctx, request.User.Username, request.User.Email, request.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{User: &v1.UserReply_Data{
		Email:    u.Email,
		Username: u.Username,
		Token:    u.Token,
	}}, nil
}

func (s *RealWorldService) GetCurrentUser(ctx context.Context, request *v1.GetCurrentUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{User: &v1.UserReply_Data{Email: "Saving"}}, nil
}

func (s *RealWorldService) UpdateUser(ctx context.Context, request *v1.UpdateUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{User: &v1.UserReply_Data{Email: "Saving"}}, nil
}

func (s *RealWorldService) GetProfile(ctx context.Context, request *v1.UsernamePathVariableRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{Profile: &v1.ProfileReply_Data{Username: "Saving"}}, nil
}

func (s *RealWorldService) FollowUser(ctx context.Context, request *v1.UsernamePathVariableRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{Profile: &v1.ProfileReply_Data{Username: "Saving"}}, nil
}

func (s *RealWorldService) UnfollowUser(ctx context.Context, request *v1.UsernamePathVariableRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{Profile: &v1.ProfileReply_Data{Username: "Saving"}}, nil
}
