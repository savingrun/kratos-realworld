package biz

import (
	"context"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string `json:"email"`
	Username     string `json:"username"`
	Bio          string `json:"bio"`
	Image        string `json:"image"`
	PasswordHash string `json:"passwordHash"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

func hashPassword(password string) string {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
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
	ur      UserRepo
	pr      ProfileRepo
	jwtConf *conf.Jwt

	log *log.Helper
}

func NewUserUseCase(ur UserRepo, pr ProfileRepo, jwtConf *conf.Jwt, logger log.Logger) *UserUseCase {
	return &UserUseCase{ur: ur, pr: pr, jwtConf: jwtConf, log: log.NewHelper(logger)}
}

func (uuc *UserUseCase) generateToken(username string) string {
	return auth.GenerateToken(uuc.jwtConf.Secret, username)
}

func (uuc *UserUseCase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	user := &User{
		Username:     username,
		Email:        email,
		PasswordHash: hashPassword(password),
	}
	if _, err := uuc.ur.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    user.Email,
		Username: user.Username,
		Token:    uuc.generateToken(user.Username),
	}, nil

}

func (uuc *UserUseCase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	if len(email) == 0 {
		return nil, errors.New(422, "email", "cannot be empty")
	}
	//return uuc.ur.Login(ctx, email, password)
	u, err := uuc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(u.PasswordHash, password) {
		return nil, errors.Unauthorized("user", "login failed")
	}
	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    uuc.generateToken(u.Username),
	}, nil
}
