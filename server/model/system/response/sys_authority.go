package response

import "github.com/Luxii44/library/server/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      system.SysAuthority `json:"authority"`
	OldAuthorityId uint                `json:"oldAuthorityId"` // 旧角色ID
}
