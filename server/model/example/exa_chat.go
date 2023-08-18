package example

import (
	"github.com/Luxii44/library/server/global"
	uuid "github.com/satori/go.uuid"
)

type ChatMessage struct {
	global.GVA_MODEL
	ChatKey        int64     `json:"chatKey" gorm:"unique_index,comment:聊天室key"`
	Sender_Id      uuid.UUID `json:"senderId" gorm:"comment:发送人"`        //发送人
	ReceiverId     uuid.UUID `json:"receiverId" gorm:"comment:接收人"`      //接收人
	Relation       uint      `json:"relation" gorm:"comment:关系标识"`       //关系标识
	MessageType    int64     `json:"messageType" gorm:"comment:消息类型"`    //消息类型
	MessageContent string    `json:"messageContent" gorm:"comment:消息内容"` //消息内容
	Time           int64     `json:"time" gorm:"comment:发送时间"`           //发送时间
	Extend         string    `json:"extend" gorm:"comment '扩展',"`
	DelFlag        uint      `json:"delFlag" gorm:"comment:删除标识"` //删除标识
	Remark         int64     `json:"remark" gorm:"comment:备注"`    //备注
}

func (ChatMessage) TableName() string {
	return "exa_chat_message"
}

type ChatGroup struct {
	global.GVA_MODEL
	GroupCode    int64     `json:"groupCode" gorm:"unique_index,comment:群聊编号"` //群聊编号
	GroupName    int64     `json:"groupName" gorm:"comment:群聊名称"`              //群聊名称
	Avatar       string    `json:"avatar" gorm:"comment:群头像"`                  //群头像
	Master       string    `json:"master" gorm:"comment:群主"`                   //群主
	Announcement string    `json:"announcement" gorm:"comment:公告"`             //公告
	OpenInvite   uuid.UUID `json:"openInvite" gorm:"comment:是否允许邀请"`           //是否允许邀请
	InviteCheck  int64     `json:"inviteCheck" gorm:"comment:是否需要审核"`          //是否需要审核
	DelFlag      uint      `json:"delFlag" gorm:"comment:删除标识"`                //删除标识
	Remark       int64     `json:"remark" gorm:"comment:备注"`                   //备注
}

func (ChatGroup) TableName() string {
	return "exa_chat_group"
}

type ChatGroupUser struct {
	global.GVA_MODEL
	GroupId int64  `json:"groupId" gorm:"comment:群聊id"` //群聊id
	UserId  int64  `json:"userId" gorm:"comment:用户id"`  //用户id
	State   string `json:"state" gorm:"comment:用户状态"`   //用户状态
	Apply   string `json:"apply" gorm:"comment:申请消息"`   //申请消息
	DelFlag uint   `json:"delFlag" gorm:"comment:删除标识"` //删除标识
	Remark  int64  `json:"remark" gorm:"comment:备注"`    //备注
}

func (ChatGroupUser) TableName() string {
	return "exa_chat_group"
}
