package request

import (
	"github.com/Luxii44/library/server/model/common/request"
	"github.com/Luxii44/library/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
