package system

import (
	"github.com/Luxii44/library/server/global"
	"github.com/satori/go.uuid"
)

type SysPosition struct {
	global.GVA_MODEL
	UUID             uuid.UUID `json:"uuid" gorm:"index;comment:职位UUID"`
	PositionCode     string    `json:"organizationCode" gorm:"comment:组织编码"`            // 职位编码
	PositionName     string    `json:"positionName" gorm:"comment:职位全称"`                // 职位名
	OrganizationUUID uuid.UUID `json:"organizationUUID" gorm:"comment:组织ID"`            // 组织编码
	ParentUUID       uuid.UUID `json:"parentUUID" gorm:"comment:上级id"`                  // 上级编码
	CompanyUUID      uuid.UUID `json:"companyUUID" gorm:"comment:组织公司"`                 // 公司编号
	Enable           int       `json:"enable" gorm:"default:1;comment:职位是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	Multiuser        []SysUser `json:"multiuser" gorm:"many2many:sys-position-user;"`
}

func (SysPosition) TableName() string {
	return "sys_position"
}
