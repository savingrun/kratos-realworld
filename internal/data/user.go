package data

import (
	"context"
	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/internal/biz"
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
	rv := r.data.db.Create(&u)
	return user, rv.Error
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
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
