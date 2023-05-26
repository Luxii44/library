package request

import (
	"github.com/Luxii44/library/server/model/common/request"
	"github.com/Luxii44/library/server/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
