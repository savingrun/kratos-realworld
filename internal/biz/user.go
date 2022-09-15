package biz

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string `json:"username"`
	Token        string `json:"token"`
	Email        string `json:"email"`
	Bio          string `json:"bio"`
	Image        string `json:"image"`
	Password     string `json:"password"`
	PasswordHash string `json:"passwordHash"`
}

func hashPassword(password string) string {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", bcryptPassword)
	return string(bcryptPassword)
}

func verifyPassword(hashPassword, inputPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(inputPassword)); err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
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

func (uuc *UserUseCase) Register(ctx context.Context, username, email, password string) (*User, error) {
	user := &User{
		Username: username,
		Email:    email,
		Password: hashPassword(password),
	}
	if userRepo, err := uuc.ur.CreateUser(ctx, user); err != nil {
		return userRepo, err
	}
	return nil, nil

}

func (uuc *UserUseCase) Login(ctx context.Context, email, password string) (*User, error) {
	//return uuc.ur.Login(ctx, email, password)
	return nil, nil
}
