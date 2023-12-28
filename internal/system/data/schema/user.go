package schema

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"column:username;comment:账号"`
	Phone        string `json:"phone" gorm:"column:phone;comment:手机号"`
	Password     string `json:"password" gorm:"column:password;comment:密码"`
	Avatar       string `json:"avatar" gorm:"column:avatar;comment:头像"`
	Introduction string `json:"introduction" gorm:"column:introduction;comment:简介"`
	Email        string `json:"email" gorm:"column:email;comment:邮箱"`
	Address      string `json:"address" gorm:"column:address;comment:地址"`
	Remark       string `json:"remark" gorm:"column:remark;comment:备注"`
	RoleID       int    `json:"role_id" gorm:"column:role_id;comment:角色ID"`
	Type         int    `json:"type" gorm:"column:type;comment:类型,0: 管理员 1: 普通用户"` // 0: 管理员 1: 普通用户
	Status       int    `json:"status" gorm:"column:status;comment:状态"`            // 0: 禁用 1: 启用
}
