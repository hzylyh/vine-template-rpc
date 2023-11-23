package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Site holds the schema definition for the Site entity.
type Site struct {
	ent.Schema
}

func (Site) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_site"},
	}
}

// Fields of the Site.
func (Site) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New).Comment("站点ID"),
		field.String("name").Unique().Comment("站点名称"),
		field.String("type").Optional().Comment("手机号"),
		field.String("dept").Optional().Comment("责任部门"),
		field.String("owner").Optional().Comment("关系责任人"),
		field.String("remark").Optional().Comment("现场机MN编号"),
		field.String("lon").Optional().Comment("经度"),
		field.String("lat").Optional().Comment("纬度"),
		field.Int("status").Default(1).Comment("状态"), // 0: 停用 1: 在用
		field.String("address").Optional().Comment("地址"),
		field.String("last_service_time").Optional().Comment("最后一次服务时间"),
		field.String("create_time").Optional().Comment("创建时间"),
		field.String("creator").Optional().Comment("创建人"),
	}
}

// Edges of the Site.
func (Site) Edges() []ent.Edge {
	return nil
}
