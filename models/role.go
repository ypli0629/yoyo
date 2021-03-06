package models

import (
	"github.com/yoyo-inc/yoyo/common/db"
	"github.com/yoyo-inc/yoyo/core"
)

// Role represents user role
type Role struct {
	core.Model
	Name        string       `json:"name" gorm:"size:256;unique;not null;comment:角色名" binding:"required,max=250"`
	Remark      string       `json:"remark" gorm:"size:256;comment:备注"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"`
}

func init() {
	db.AddAutoMigrateModel(&Role{})
}

type QueryRole struct {
	// 角色名
	Name string `json:"name"`
}
