package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Username string `json:"username"`
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
}

type ProfileRepo interface {
}

type UserUseCase struct {
	ur UserRepo
	pr ProfileRepo

	log *log.Helper
}

func NewUserUseCase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{ur: ur, pr: pr, log: log.NewHelper(logger)}
}

func (uuc *UserUseCase) Register(ctx context.Context, user *User) error {
	if err, _ := uuc.ur.CreateUser(ctx, user); err != nil {

	}
	return nil
}
