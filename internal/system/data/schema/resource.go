package schema

import "gorm.io/gorm"

//field.String("name").NotEmpty().Comment("资源名称"),
//field.String("code").NotEmpty().Comment("资源编码"),
//field.String("path").NotEmpty().Comment("资源路径"),
//field.String("type").NotEmpty().Comment("资源类型"),
//field.String("remark").Comment("备注"),
//field.Int("status").Default(1).Comment("状态"),

type Resource struct {
	gorm.Model
	Name   string `json:"name" gorm:"column:name;comment:资源名称"`
	Code   string `json:"code" gorm:"column:code;comment:资源编码"`
	Path   string `json:"path" gorm:"column:path;comment:资源路径"`
	Type   string `json:"type" gorm:"column:type;comment:资源类型"`
	Remark string `json:"remark" gorm:"column:remark;comment:备注"`
	Status int    `json:"status" gorm:"column:status;comment:状态"`
}
