package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	return user, nil
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
