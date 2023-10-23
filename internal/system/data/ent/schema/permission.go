package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("权限名称"),
		field.String("code").NotEmpty().Comment("权限编码"),
		field.String("remark").Comment("备注"),
		field.Int("status").Default(1).Comment("状态"),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return nil
}
