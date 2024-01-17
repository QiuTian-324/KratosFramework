package data

import (
	"akita/quantum_cat/app/sms/internal/dao"
	ci "akita/quantum_cat/app/sms/internal/data/sms/ent"
	"akita/quantum_cat/pkg/model"
	"context"
	entschema "entgo.io/ent/dialect/sql/schema"
)

// MigrateOptions 定义了数据库迁移的选项。
type MigrateOptions struct {
	Debug            bool // 是否启用调试模式。
	DropColumn       bool // 是否删除列。
	DropIndex        bool // 是否删除索引。
	CreateForeignKey bool // 是否创建外键。
}

// Migrate 执行数据库迁移。
func Migrate(ctx context.Context, daoIn *dao.DAO, opt *MigrateOptions) error {
	var schemas []model.EntSchema

	// 根据调试选项决定是否启用调试模式。
	if opt.Debug {
		schemas = append(schemas,
			ci.NewClient(ci.Driver(daoIn.DbDriver)).Debug().Schema,
		)
	} else {
		schemas = append(schemas,
			ci.NewClient(ci.Driver(daoIn.DbDriver)).Schema,
		)
	}

	// 使用 model.EntMigrateSchemas 执行数据库迁移，提供所需的选项。
	return model.EntMigrateSchemas(ctx, schemas,
		entschema.WithAtlas(true),
		entschema.WithDropColumn(opt.DropColumn),
		entschema.WithDropIndex(opt.DropIndex),
		entschema.WithForeignKeys(opt.CreateForeignKey),
	)
}
