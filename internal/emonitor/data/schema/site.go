package schema

import (
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model
	Name            string `json:"name" gorm:"column:name;comment:站点名称"`
	Type            string `json:"type" gorm:"column:type;comment:类型"`
	Dept            string `json:"dept" gorm:"column:dept;comment:责任部门"`
	Owner           string `json:"owner" gorm:"column:owner;comment:关系责任人"`
	mn              string `json:"mn" gorm:"column:mn;comment:现场机MN编号"`
	Remark          string `json:"remark" gorm:"column:remark;comment:备注"`
	Lon             string `json:"lon" gorm:"column:lon;comment:经度"`
	Lat             string `json:"lat" gorm:"column:lat;comment:纬度"`
	Status          int32  `json:"status" gorm:"column:status;comment:状态"` // 0: 停用 1: 在用
	Address         string `json:"address" gorm:"column:address;comment:地址"`
	LastServiceTime string `json:"lastServiceTime" gorm:"column:last_service_time;comment:最后一次服务时间"`
	Creator         string `json:"creator" gorm:"column:creator;comment:创建人"`
}

func (Site) TableName() string {
	return "tb_site"
}
