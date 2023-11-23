package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

func (Equipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tb_equipment"},
	}
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Default(uuid.New).Comment("设备ID"),
		field.String("name").Optional().Comment("设备名称"),
		field.String("type").Optional().Comment("设备类型"),
		field.String("status").Optional().Comment("设备状态"),
		field.String("remark").Optional().Comment("备注"),
		field.String("lon").Optional().Comment("经度"),
		field.String("lat").Optional().Comment("纬度"),
		field.String("address").Optional().Comment("地址"),
		field.String("install_time").Optional().Comment("安装时间"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return nil
}
