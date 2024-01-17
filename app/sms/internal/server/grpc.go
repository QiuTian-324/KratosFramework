package server

import (
	"akita/quantum_cat/app/sms/internal/conf"
	"akita/quantum_cat/app/sms/internal/service"
	v1 "akita/quantum_cat/proto/sms/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(c *conf.Server, smsService *service.SmsService, logger log.Logger) *grpc.Server {
	// 定义 gRPC 服务器选项的初始值
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),    // 添加恢复中间件，用于在处理请求时恢复从 panic 中恢复
			logging.Server(logger), // 添加日志中间件，用于记录服务器端的日志
		),
	}

	// 根据配置参数设置 gRPC 服务器选项
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

	// 创建 gRPC 服务器实例
	srv := grpc.NewServer(opts...)

	// 注册用户服务到 gRPC 服务器
	v1.RegisterSmsServer(srv, smsService)

	return srv
}
