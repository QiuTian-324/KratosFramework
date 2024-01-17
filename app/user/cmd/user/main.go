package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-Hermes/app/user/cmd/user/subcommand/migrate"
	"go-Hermes/app/user/cmd/user/subcommand/serve"
	"go-Hermes/pkg/tools"
	"os"
)

var Verbose bool

func main() {

	// rootCmd 表示根命令，当没有任何子命令被调用时使用
	var rootCmd = &cobra.Command{
		Use:   "user",
		Short: "用户模块.",
		Long: `用户模块.
默认运行 "serve". 使用 "-h" 参数查看所有可用命令.`,
	}

	// CmdServe 表示 "serve" 子命令
	var CmdServe = &cobra.Command{
		Use:   "serve",
		Short: "启动服务",
		Long:  `启动用户服务。提供服务及相关接口`,
		Run:   serve.StartServe,
		Args:  cobra.MaximumNArgs(1),
	}
	CmdServe.Flags().String("pprof-debug-addr", "", "设置值后，将开启pprof debug服务. 如:'127.0.0.1:8686'")

	// CmdMigrate 表示 "migrate" 子命令
	var CmdMigrate = &cobra.Command{
		Use:   "migrate",
		Short: "数据迁移",
		Long:  `数据迁移。将user下的schema同步到数据库，并进行相关数据初始化工作`,
		Run:   migrate.RunUp,
		Args:  cobra.MaximumNArgs(1),
	}
	{
		CmdMigrate.PersistentFlags().UintP("timeout", "t", 0, "迁移执行超时时间，单位：秒。大于等于0的整数，等于0时，永不超时。")
		CmdMigrate.PersistentFlags().BoolP("verbose", "v", false, "显示详细日志，如：打印sql日志等。")
		CmdMigrate.PersistentFlags().Bool("drop-column", false, "是否删除原有字段")
		CmdMigrate.PersistentFlags().Bool("drop-index", false, "是否删除原有索引")
		CmdMigrate.PersistentFlags().Bool("foreign-key", true, "是否创建外键")
	}

	// 将子命令添加到根命令
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "conf", "c", false, "配置文件所在目录或路径, 例如: -c ../../configs")
	rootCmd.AddCommand(
		CmdServe,
		CmdMigrate,
	)

	// 注册默认命令
	tools.RegisterDefaultCmd(rootCmd, CmdServe)

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		tools.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
