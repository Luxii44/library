package system

// SysMenuPermission 是 sysMenu 和 sysOrganization 的连接表
type SysMenuPermission struct {
	SysMenuId       uint `gorm:"column:sys_menu_id"`
	SysPermissionId uint `gorm:"column:sys_permission_id"`
}

func (s *SysMenuPermission) TableName() string {
	return "sys_menu_permission"
}
