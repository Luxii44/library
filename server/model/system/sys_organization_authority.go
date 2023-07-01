package system

// SysOrganizationAuthority 是 sysOrganization 和 sysAuthority 的连接表
type SysOrganizationAuthority struct {
	SysAuthorityAuthorityId uint `gorm:"column:sys_authority_authority_id"`
	SysOrganizationId       uint `gorm:"column:sys_organization_id"`
}

func (s *SysOrganizationAuthority) TableName() string {
	return "sys_organization_authority"
}
