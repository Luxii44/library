package system

import (
	"github.com/Luxii44/library/server/global"
	"github.com/satori/go.uuid"
)

type SysOrganization struct {
	global.GVA_MODEL
	UUID                 uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	OrganizationName     string    `json:"organizationName" gorm:"comment:组织名"`      // 组织名
	OrganizationFullName string    `json:"organizationFullName" gorm:"comment:组织全称"` // 组织全称
	ParentUUID           uuid.UUID `json:"parentUUID" gorm:"comment:父级组织"`           // 父级组织
	CompanyUUID          string    `json:"companyUUID" gorm:"comment:组织公司"`          // 组织公司
	TopOrganizationUUID  string    `json:"topOrganizationUUID" gorm:"comment:顶级组织"`  // 顶级组织
}

func (SysOrganization) TableName() string {
	return "sys_organization"
}
