package sign

// BusinessType 表示业务类型的整数枚举。
type BusinessType int

const (
	// ResetPasswordType 用于表示密码重置业务类型。
	ResetPasswordType BusinessType = iota
	// RegisterType 用于表示用户注册业务类型。
	RegisterType
	// LoginType 用于表示用户登录业务类型。
	LoginType
	// RetrievePasswordType 用于表示用户找回密码业务类型。
	RetrievePasswordType
)

// BusinessTypeMap 是业务类型到字符串的映射。
var BusinessTypeMap = map[BusinessType]string{
	ResetPasswordType:    "resetPassword",
	RegisterType:         "register",
	LoginType:            "login",
	RetrievePasswordType: "retrievePassword",
}

// JwtSmsInfoClaims 结构表示 JWT 中包含的短信信息声明。
type JwtSmsInfoClaims struct {
	Type        BusinessType `json:"type"`         // 业务类型
	ResourceIDs []uint64     `json:"resource_ids"` // 资源 IDs
}

// GetResIDs 方法用于获取 JWT 中的资源 IDs。
// 返回：
// - []uint64: 资源 IDs。
func (juic *JwtSmsInfoClaims) GetResIDs() []uint64 {
	return juic.ResourceIDs
}
