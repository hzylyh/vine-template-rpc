package schema

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name;comment:设备名称"`
	Type        string `json:"type" gorm:"column:type;comment:设备类型"`
	Status      string `json:"status" gorm:"column:status;comment:设备状态"`
	Remark      string `json:"remark" gorm:"column:remark;comment:备注"`
	Lon         string `json:"lon" gorm:"column:lon;comment:经度"`
	Lat         string `json:"lat" gorm:"column:lat;comment:纬度"`
	Address     string `json:"address" gorm:"column:address;comment:地址"`
	InstallTime string `json:"installTime" gorm:"column:install_time;comment:安装时间"`
}

func (Equipment) TableName() string {
	return "tb_equipment"
}
