package request

import (
	"github.com/Luxii44/library/server/model/common/request"
	"github.com/Luxii44/library/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
