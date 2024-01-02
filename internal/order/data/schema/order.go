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
}

func (Order) TableName() string {
	return "tb_order"
}
