package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// User 定义返回数据结构体
type User struct {
	ID       int64
	Mobile   string
	UserName string
	Password string
	NickName string
	Birthday int64
	Gender   string
	Token    string
	Role     int
}

// UserRepo 接口
type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
	LoginByUserName(context.Context, *User) (*User, error)
}

// UserUsecase 用户实例
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase 初始化用户实例
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// Create 实现创建用户接口，调用data的Create方法，进行创建用户
func (uc *UserUsecase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

// LoginByUserName 实现用户登录接口
func (uc *UserUsecase) LoginByUserName(ctx context.Context, u *User) (*User, error) {
	return uc.repo.LoginByUserName(ctx, u)
}
