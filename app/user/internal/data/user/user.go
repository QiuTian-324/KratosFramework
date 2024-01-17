package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"go-Hermes/app/user/internal/biz"
	"go-Hermes/app/user/internal/conf"
	"go-Hermes/app/user/internal/dao"
	"go-Hermes/app/user/internal/data/user/ent"
)

//// User 定义数据表结构体
//type User struct {
//	ID          int64      `gorm:"primaryKey"`
//	UserName    string     `gorm:"column:user_name;index:idx_username;unique;type:varchar(25) comment '用户名，用户唯一标识';not null"`
//	NickName    string     `gorm:"column:nick_name;type:varchar(25) comment '用户昵称'"`
//	Mobile      string     `gorm:"column:mobile;index:idx_mobile;unique;type:varchar(11) comment '手机号码，用户唯一标识';not null"`
//	Password    string     `gorm:"column:pass_word;type:varchar(100) comment '用户密码';not null "` // 用户密码的保存需要注意是否加密
//	Birthday    *time.Time `gorm:"column:birthday;type:datetime comment '出生日日期'"`
//	Gender      string     `gorm:"column:gender;default:male;type:varchar(16) comment 'female:女,male:男'"`
//	Role        int        `gorm:"column:role;default:1;type:int comment '1:普通用户，2:管理员'"`
//	CreatedAt   time.Time  `gorm:"column:add_time"`
//	UpdatedAt   time.Time  `gorm:"column:update_time"`
//	DeletedAt   gorm.DeletedAt
//	IsDeletedAt bool
//}

type userRepo struct {
	data   *dao.DAO
	db     *ent.Client
	config *conf.Data
	log    *log.Helper
}

// NewUserRepo 初始化用户实例
// NewUserRepo . 这里需要注意，上面 data 文件 wire 注入的是此方法，方法名不要写错了
func NewUserRepo(data *dao.DAO, config *conf.Data, logger log.Logger) biz.UserRepo {
	db := ent.NewClient(ent.Driver(data.DbDriver))
	return &userRepo{
		data:   data,
		db:     db,
		config: config,
		log:    log.NewHelper(logger),
	}
}

// CreateUser .
//func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
//var user User
//// 验证是否已经创建
//result := r.data.db.Where(&biz.User{Mobile: u.Mobile}).First(&user)
//if result.RowsAffected == 1 {
//	return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
//}
//
//// 用户不存在就初始化user数据，用于创建
//user.Mobile = u.Mobile
//user.NickName = u.NickName
//user.Password = encrypt(u.Password) // 密码加密
//res := r.data.db.Create(&user)
//if res.Error != nil {
//	return nil, status.Errorf(codes.Internal, res.Error.Error())
//}
//
//return &biz.User{
//	ID:       user.ID,
//	Mobile:   user.Mobile,
//	Password: user.Password,
//	NickName: user.NickName,
//	Gender:   user.Gender,
//	Role:     user.Role,
//}, nil
//	return nil, nil
//}

//func (r *userRepo) LoginByUserName(ctx context.Context, u *biz.User) (*biz.User, error) {
//var user User
//result := r.data.db.Where(&biz.User{UserName: u.UserName}).First(&user)
//
//if result.RowsAffected == 0 {
//	return nil, status.Errorf(codes.Unauthenticated, result.Error.Error())
//}
//
//return &biz.User{
//	ID:       user.ID,
//	Mobile:   user.Mobile,
//	UserName: user.UserName,
//	Password: user.Password,
//	NickName: user.NickName,
//	Birthday: user.Birthday.Unix(),
//	Gender:   user.Gender,
//	Role:     user.Role,
//}, nil
//	return nil, nil
//}
//
//// Password encryption
//func encrypt(psd string) string {
//	options := &password.Options{SaltLen: 16, Iterations: 10000, KeyLen: 32, HashFunction: sha512.New}
//	salt, encodedPwd := password.Encode(psd, options)
//	return fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
//}
