package system

import (
	"github.com/Luxii44/library/server/global"
	"github.com/satori/go.uuid"
)

type SysOrganization struct {
	global.GVA_MODEL
	UUID                 uuid.UUID  `json:"uuid" gorm:"index;comment:用户UUID"`
	OrganizationCode     string     `json:"organizationCode" gorm:"comment:组织编码"`            // 组织编号
	OrganizationName     string     `json:"organizationName" gorm:"comment:组织名"`             // 组织名
	OrganizationFullName string     `json:"organizationFullName" gorm:"comment:组织全称"`        // 组织全称
	ParentUUID           uuid.UUID  `json:"parentUUID" gorm:"comment:父级组织"`                  // 父级组织
	CompanyUUID          *uuid.UUID `json:"companyUUID" gorm:"comment:组织公司"`                 // 组织公司(子公司还是分公司）
	TopOrganizationUUID  uuid.UUID  `json:"topOrganizationUUID" gorm:"comment:顶级组织"`         // 顶级组织
	Enable               int        `json:"enable" gorm:"default:1;comment:组织是否被冻结 1正常 2冻结"` //组织是否被冻结 1正常 2冻结
}

func (SysOrganization) TableName() string {
	return "sys_organization"
}
