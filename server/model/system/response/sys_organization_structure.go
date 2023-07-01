package response

import "github.com/Luxii44/library/server/model/system"

// Get Organization List structure
type Organization struct {
	system.SysOrganization
	//OrganizationCode     string        `json:"organizationCode"`     // 组织编号
	//OrganizationName     string        `json:"organizationName"`     // 组织名
	//OrganizationFullName string        `json:"organizationFullName"` // 组织全称
	//CompanyName         string `json:"companyName"`
	//TopOrganizationName string `json:"topOrganizationName"`
}
