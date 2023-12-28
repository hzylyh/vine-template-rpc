package schema

type Role struct {
	ID          int64  `json:"id" gorm:"column:id;primary_key;auto_increment;comment:角色id"`
	Name        string `json:"name" gorm:"column:name;comment:角色名称"`
	Description string `json:"description" gorm:"column:description;comment:备注"`
	Status      int    `json:"status" gorm:"column:status;comment:状态"`
}
