package dao

import (
	"akita/quantum_cat/app/sms/internal/conf"
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
)

type DAO struct {
	DbDriver dialect.Driver
}

func NewDao(c *conf.Data, logger log.Logger) (*DAO, func(), error) {
	// 创建一个日志记录助手
	l := log.NewHelper(logger)

	// 使用给定的数据库驱动和数据源信息打开数据库连接
	drv, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database, err: %v", err)
	}

	// 定义用于关闭资源的 cleanup 函数
	cleanup := func() {
		log.Info("closing the data resources")
		if err := drv.Close(); err != nil {
			l.Warnf("database error on closing: %v", err)
		}
	}

	// 使用 dialect.DebugWithContext 包装数据库连接，以便在执行 SQL 语句时记录调试信息
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		l.WithContext(ctx).Info(i...)
	})

	// 返回创建的 DAO 对象及 cleanup 函数
	return &DAO{
		DbDriver: sqlDrv,
	}, cleanup, nil
}
