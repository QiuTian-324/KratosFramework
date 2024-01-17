package migrate

import (
	"akita/quantum_cat/app/sms/cmd/sms/subcommand/serve"
	"akita/quantum_cat/app/sms/internal/conf"
	"akita/quantum_cat/app/sms/internal/dao"
	"akita/quantum_cat/app/sms/internal/data"
	"context"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// RunUp 数据库schema同步及数据初始化
func RunUp(cmd *cobra.Command, args []string) {
	var c config.Config
	var err error
	var mCtx context.Context
	timeout, _ := cmd.Flags().GetUint("timeout")
	var cancel context.CancelFunc

	if timeout > 0 {
		mCtx, cancel = context.WithTimeout(context.TODO(), time.Second*time.Duration(timeout))
		defer cancel()
	} else {
		mCtx = context.Background()
	}
	if len(args) > 0 {
		c, err = serve.LoadConfig(args[0])
	} else {
		c, err = serve.LoadConfig(serve.Flagconf)
	}
	if err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	daoIn, cleanFun, err := dao.NewDao(bc.Data, log.NewStdLogger(os.Stdout))
	if err != nil {
		panic(err)
	} else {
		defer cleanFun()
	}

	debug, _ := cmd.Flags().GetBool("verbose")
	dropColumn, _ := cmd.Flags().GetBool("drop-column")
	dropIndex, _ := cmd.Flags().GetBool("drop-index")
	foreignKey, _ := cmd.Flags().GetBool("create-foreign-key")

	if err := data.Migrate(mCtx, daoIn, &data.MigrateOptions{
		Debug:            debug,
		DropColumn:       dropColumn,
		DropIndex:        dropIndex,
		CreateForeignKey: foreignKey,
	}); err != nil {
		panic(err)
	}
}
