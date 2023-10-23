package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("username").Unique().Comment("账号"),
		field.String("phone").Optional().Comment("手机号"),
		field.String("password").Comment("密码"),
		field.String("avatar").Optional().Comment("头像"),
		field.String("introduction").Optional().Comment("简介"),
		field.String("email").Optional().Comment("邮箱"),
		field.String("address").Optional().Comment("地址"),
		field.String("remark").Optional().Comment("备注"),
		field.Int("status").Default(1).Comment("状态"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
