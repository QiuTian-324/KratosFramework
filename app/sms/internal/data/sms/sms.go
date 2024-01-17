package sms

import (
	"akita/quantum_cat/app/sms/internal/biz"
	"akita/quantum_cat/app/sms/internal/conf"
	"akita/quantum_cat/app/sms/internal/dao"
	"akita/quantum_cat/app/sms/internal/data/sms/ent"
	"github.com/go-kratos/kratos/v2/log"
)

type SmsRepo struct {
	data   *dao.DAO
	db     *ent.Client
	config *conf.Data
	log    *log.Helper
}

// NewUserRepo 初始化用户实例
// NewUserRepo . 这里需要注意，上面 data 文件 wire 注入的是此方法，方法名不要写错了
func NewUserRepo(data *dao.DAO, config *conf.Data, logger log.Logger) biz.SmsRepo {
	db := ent.NewClient(ent.Driver(data.DbDriver))
	return &SmsRepo{
		data:   data,
		db:     db,
		config: config,
		log:    log.NewHelper(logger),
	}
}
