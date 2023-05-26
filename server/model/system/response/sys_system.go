package response

import "github.com/Luxii44/library/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
