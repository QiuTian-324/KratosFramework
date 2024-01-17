package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema // User 结构体嵌套 ent.Schema，表示它是 ent 框架的数据模型一部分。
}

// Fields 定义了 User 实体的字段。
func (User) Fields() []ent.Field {
	return []ent.Field{
		// age: 整数类型的字段。
		field.Int("age"),
		// name: 字符串类型的字段，带有注释 "用户名称"，并启用数据库表字段的注释。
		field.String("name").Comment("用户名称").Annotations(entsql.WithComments(true)),
		// email: 字符串类型的可选字段，带有注释 "邮箱"。
		field.String("email").Comment("邮箱").Annotations(entsql.WithComments(true)).Optional(),
		// phone: 字符串类型的可选字段，带有注释 "手机号-作为账号登录"。
		field.String("phone").Comment("手机号-作为账号登录").Annotations(entsql.WithComments(true)).Optional(),
		// password: 字符串类型的字段，带有注释 "账号密码"。
		field.String("password").Comment("账号密码").Annotations(entsql.WithComments(true)),
	}
}

// Edges 定义了 User 实体的边缘（Edges），在这个例子中没有定义任何边缘。
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes 定义了 User 实体的索引，其中包含一个唯一索引，该索引涉及到 name 字段。
// 唯一索引确保了表中的 name 字段的唯一性，即不允许两行具有相同的 name 值。
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
	}
}
