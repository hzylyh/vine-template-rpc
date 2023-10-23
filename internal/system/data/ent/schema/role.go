package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("角色名称"),
		field.String("code").NotEmpty().Comment("角色编码"),
		field.String("remark").Comment("备注"),
		field.Int("status").Default(1).Comment("状态"),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
