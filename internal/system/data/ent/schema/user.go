package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_user"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New).Comment("用户id"),
		field.String("username").Unique().Comment("账号"),
		field.String("phone").Optional().Comment("手机号"),
		field.String("password").Comment("密码"),
		field.String("avatar").Optional().Comment("头像"),
		field.String("introduction").Optional().Comment("简介"),
		field.String("email").Optional().Comment("邮箱"),
		field.String("address").Optional().Comment("地址"),
		field.String("remark").Optional().Comment("备注"),
		field.String("type").Default("1").Comment("类型,0: 管理员 1: 普通用户"), // 0: 管理员 1: 普通用户
		field.Int("status").Default(1).Comment("状态"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
