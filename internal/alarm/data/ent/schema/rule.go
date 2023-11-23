package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Rule holds the schema definition for the Rule entity.
type Rule struct {
	ent.Schema
}

func (Rule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_alarm_rule"},
	}
}

// Fields of the Rule.
func (Rule) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New).Comment("告警规则ID"),
		field.String("name").Comment("告警规则名称"),
		field.String("type").Comment("告警规则类型"),
		field.String("level").Comment("告警规则级别"),
		field.String("content").Comment("告警规则内容"),
		field.String("remark").Optional().Comment("备注"),
		field.Int("status").Default(1).Comment("状态"), // 0: 停用 1: 在用
		field.String("create_time").Optional().Comment("创建时间"),
		field.String("creator").Optional().Comment("创建人"),
	}
}

// Edges of the Rule.
func (Rule) Edges() []ent.Edge {
	return nil
}
