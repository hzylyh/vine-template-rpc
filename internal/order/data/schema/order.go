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
	Name string `json:"name,omitempty" gorm:"column:name;comment:工单名称 0:故障工单 1:计划工单"`
	Type string `json:"type,omitempty" gorm:"column:type;comment:工单类型"`
	// 关联的用户
	UserID uint `json:"userId,omitempty" gorm:"column:user_id;comment:用户ID"`
	// 关联的设备
	EquipmentID uint   `json:"equipmentId,omitempty" gorm:"column:equipment_id;comment:设备ID"`
	Priority    string `json:"priority" gorm:"default:0;column:priority;comment:工单优先级"`
	Describe    string `json:"describe,omitempty" gorm:"column:describe;comment:工单描述"`
	Status      string `json:"status,omitempty" gorm:"default:0;column:status;comment:工单状态 0:待审核 1:待处理 2:已完成 3:延期 4:已取消"`
}

func (Order) TableName() string {
	return "tb_order"
}
