package global

import (
	uuid "github.com/satori/go.uuid"
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey;autoIncrement"` // 主键ID
	Creator   uuid.UUID      // 创建者
	Modifier  uuid.UUID      // 更新者
	Deleter   uuid.UUID      `gorm:"index" json:"-"` // 删除者
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
