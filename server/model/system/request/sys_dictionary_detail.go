package request

import (
	"github.com/Luxii44/library/server/model/common/request"
	"github.com/Luxii44/library/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
