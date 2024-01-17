package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"user/app/sms/pkg/sign"
)

// Captcha 结构表示验证码信息。
type Captcha struct {
	CaptchaID    string // 验证码ID
	CaptchaImage string // 验证码图片数据
}

// CodeMessage 结构表示短信操作的消息和状态码。
type CodeMessage struct {
	Message string // 状态描述
	Code    int32  // 状态码
}

// SmsBaseInfo 结构包含了短信基本信息。
type SmsBaseInfo struct {
	Type sign.BusinessType // 业务类型
}

// SmsTokenInfo 结构表示包含短信基本信息的令牌信息。
type SmsTokenInfo struct {
	SmsBaseInfo
	Jwt string // JWT 令牌
}

// SmsRepo 接口定义了短信相关操作的存储库接口。
type SmsRepo interface {
}

// SmsUseCase 结构定义了短信服务的用例。
type SmsUseCase struct {
	repo SmsRepo // 短信存储库
	//grpcHS sms.UserClient      // 用户 gRPC 客户端
	log *log.Helper // 日志记录器助手
	//cfg    *auth.JwtIssueConfig // JWT 配置
}

// NewSmsUseCase 函数用于创建 SmsUseCase 的实例。
// 参数：
// - repo: 短信存储库实例。
// - grpcUser: 用户 gRPC 客户端实例。
// - jwtConfig: JWT 配置实例。
// - logger: 日志记录器。
// 返回：
// - *SmsUseCase: 创建的 SmsUseCase 实例。
func NewSmsUseCase(repo SmsRepo, logger log.Logger) *SmsUseCase {
	return &SmsUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/sms")), // 创建日志记录器助手。
	}
}
