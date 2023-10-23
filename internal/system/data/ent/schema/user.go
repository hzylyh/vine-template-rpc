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
		field.String("name").NotEmpty().Comment("姓名"),
		field.String("phone").NotEmpty().Comment("手机号"),
		field.String("password").NotEmpty().Comment("密码"),
		field.String("avatar").Comment("头像"),
		field.String("introduction").Comment("简介"),
		field.String("email").Comment("邮箱"),
		field.String("address").Comment("地址"),
		field.String("remark").Comment("备注"),
		field.Int("status").Default(1).Comment("状态"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
