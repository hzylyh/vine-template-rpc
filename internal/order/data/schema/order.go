/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2024/1/2 10:24
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2024/1/2 10:24
 */

package schema

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Name string `json:"name,omitempty" gorm:"column:name;comment:工单名称"`
	// 关联的设备
	EquipmentID uint   `json:"equipment_id,omitempty" gorm:"column:equipment_id;comment:设备ID"`
	Describe    string `json:"describe,omitempty" gorm:"column:describe;comment:工单描述"`
	Status      int    `json:"status,omitempty" gorm:"column:status;comment:工单状态"`
}

func (Order) TableName() string {
	return "tb_order"
}
