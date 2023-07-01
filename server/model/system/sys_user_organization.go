package system

// SysUserAuthority 是 sysUser 和 sysOrganization 的连接表
type SysUserOrganization struct {
	SysUserId         uint `gorm:"column:sys_user_id"`
	SysOrganizationId uint `gorm:"column:sys_organization_id"`
}

func (s *SysUserOrganization) TableName() string {
	return "sys_user_organization"
}
