package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Resource holds the schema definition for the Resource entity.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("资源名称"),
		field.String("code").NotEmpty().Comment("资源编码"),
		field.String("path").NotEmpty().Comment("资源路径"),
		field.String("type").NotEmpty().Comment("资源类型"),
		field.String("remark").Comment("备注"),
		field.Int("status").Default(1).Comment("状态"),
	}
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return nil
}
