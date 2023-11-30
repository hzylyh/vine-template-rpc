package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_role"},
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable().Comment("角色id"),
		field.String("name").NotEmpty().Comment("角色名称"),
		field.String("description").Optional().Comment("备注"),
		field.Int("status").Default(1).Comment("状态"),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
