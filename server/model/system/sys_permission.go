package system

import (
	"github.com/Luxii44/library/server/global"
	"github.com/satori/go.uuid"
)

type SysPermissionPoint struct {
	global.GVA_MODEL
	UUID            uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`                // 权限UUID
	PermissionCode  string    `json:"permissionCode" gorm:"index;comment:权限编码"`        // 权限编码
	PermissionName  string    `json:"permissionName"  gorm:"comment:权限名称"`             //权限名称
	PermissionGrade int       `json:"permissionGrade"  gorm:"comment:权限等级"`            //权限等级
	Enable          int       `json:"enable" gorm:"default:1;comment:权限是否被冻结 1正常 2冻结"` //权限是否被冻结 1正常 2冻结
	APIS            []SysApi  `json:"apis" gorm:"many2many:sys_permission_api;"`
	MenuId          int64     `json:"menuId" gorm:"default:1;comment:关联菜单"` //关联菜单
	//ermissionType  string    `json:"permissionType"  gorm:"comment:权限类型"`             //权限类型
}

func (SysPermissionPoint) TableName() string {
	return "sys_permission_point"
}

type SysBaseMenuPermissionPoint struct {
	global.GVA_MODEL
	BaseMenuId        int64 `json:"baseMenuId" gorm:"comment:模块ID;column:sys_base_menu_id"`                // 权限UUID
	PermissionPointId int64 `json:"permissionPointId" gorm:"comment:权限点ID;column:sys_permission_point_id"` // 权限编码
}

func (SysBaseMenuPermissionPoint) TableName() string {
	return "sys_base_menu_permission_point"
}
