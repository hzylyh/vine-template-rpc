/**
 * @Description:
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2024/1/2 11:30
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2024/1/2 11:30
 */

package schema

import "gorm.io/gorm"

type OrderHistory struct {
	gorm.Model
	OrderID    uint   `json:"order_id" gorm:"column:order_id;comment:工单ID"`
	Status     int    `json:"status" gorm:"column:status;comment:工单状态"`
	ChangeUser string `json:"change_user" gorm:"column:change_user;comment:变更用户"`
	CreateTime string `json:"create_time" gorm:"column:create_time;comment:创建时间"`
}
