package system

import (
	"github.com/Luxii44/library/server/global"
	uuid "github.com/satori/go.uuid"
	"time"
)

type SysApiOperatingRecord struct {
	global.GVA_MODEL
	APICode    string    `json:"APICode" gorm:"comment:api编码"`          // api编码
	UserId     uuid.UUID `json:"path" gorm:"comment:使用人id"`             // 使用人id
	Time       time.Time `json:"time" gorm:"comment:时间"`                // 时间
	IP         string    `json:"ip" gorm:"comment:IP地址"`                // IP地址
	Params     string    `json:"params" gorm:"comment:参数"`              // json字符串
	MacAddress string    `json:"method" gorm:"default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (SysApiOperatingRecord) TableName() string {
	return "sys_api_operating_records"
}
