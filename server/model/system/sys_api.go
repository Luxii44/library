package system

import (
	"github.com/Luxii44/library/server/global"
)

type SysApi struct {
	global.GVA_MODEL
	Code             string               `json:"code" gorm:"comment:api编码"`             // api编码
	Path             string               `json:"path" gorm:"comment:api路径"`             // api路径
	Description      string               `json:"description" gorm:"comment:api中文描述"`    // api中文描述
	ApiGroup         string               `json:"apiGroup" gorm:"comment:api组"`          // api组
	Method           string               `json:"method" gorm:"default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
	PermissionPoints []SysPermissionPoint `json:"permissionPoints" gorm:"many2many:sys_permission_api;"`
}

func (SysApi) TableName() string {
	return "sys_apis"
}
