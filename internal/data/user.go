package data

import (
	"context"
	"errors"

	kErrors "github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"

	"kratos-realworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	Email        string `gorm:"size:500"`
	Username     string `gorm:"size:500"`
	Bio          string `gorm:"size:1000"`
	Image        string `gorm:"size:1000"`
	PasswordHash string `gorm:"size:1000"`
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	u := User{
		Email:        user.Email,
		Bio:          user.Bio,
		Image:        user.Image,
		Username:     user.Username,
		PasswordHash: user.PasswordHash}
	rv := r.data.mysql.Create(&u)
	return user, rv.Error
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	u := new(User)
	result := r.data.mysql.Where("email = ?", email).Find(&u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, kErrors.NotFound("user", "not found by email")
	}
	if result.Error != nil {
		return nil, kErrors.InternalServer("db", result.Error.Error())
	}
	return &biz.User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}, nil
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

//
//func (r *userRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
//	return g, nil
//}
//
//func (r *userRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
//	return nil, nil
//}
//
//func (r *userRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
//	return nil, nil
//}
//
//func (r *userRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
//	return nil, nil
//}
