package schema

import (
	"akita/quantum_cat/pkg/model"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Message 短信信息.
type Message struct {
	ent.Schema
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		model.EntityStatMixin{},
	}
}

func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "sms_message",
		},
	}
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("biz_id").Comment("回执id").Annotations(entsql.WithComments(true)).Optional(),
		field.String("content").Comment("短信内容").Annotations(entsql.WithComments(true)),
		field.String("phone_number").Comment("短信接收方").Annotations(entsql.WithComments(true)),
		field.String("service_provider").Comment("短信提供商").Annotations(entsql.WithComments(true)),
		field.String("send_status").Comment("发送状态").Annotations(entsql.WithComments(true)),
		field.String("node_log").Comment("节点日志").Annotations(entsql.WithComments(true)),
		field.Time("submit_time").Comment("提交时间").Annotations(entsql.WithComments(true)),
		field.Time("arrive_time").Comment("到达时间").Annotations(entsql.WithComments(true)).Optional(),
		field.Time("receipt_time").Comment("回执时间").Annotations(entsql.WithComments(true)).Optional(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{}
}
