package example

import (
	"github.com/Luxii44/library/server/global"
	uuid "github.com/satori/go.uuid"
)

type MarkdownContent struct {
	global.GVA_MODEL
	Title        string    `json:"title" gorm:"unique_index,comment:标题"`
	Content      string    `json:"content" gorm:"unique_index,comment:博客内容"`
	Belong       uuid.UUID `json:"belong" gorm:"comment:所属人"`       //所属人
	Status       int64     `json:"status" gorm:"comment:博客状态"`      //博客状态
	ReadCount    int64     `json:"readCount" gorm:"comment:阅读数"`    //阅读数
	CommentCount int64     `json:"commentCount" gorm:"comment:评论数"` //评论数
	CollectCount int64     `json:"collectCount" gorm:"comment:收藏数"` //评论数
	Extend       string    `json:"extend" gorm:"comment '扩展',"`
	Remark       int64     `json:"remark" gorm:"comment:备注"` //备注
}

func (MarkdownContent) TableName() string {
	return "exa_Markdown_message"
}
