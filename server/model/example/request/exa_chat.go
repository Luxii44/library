package request

import (
	uuid "github.com/satori/go.uuid"
)

// Get Recruit Plan List structure
type Message struct {
	Sender_Id      uuid.UUID `json:"senderId" gorm:"comment:发送人"`        //发送人
	ReceiverId     uuid.UUID `json:"receiverId" gorm:"comment:接收人"`      //接收人
	MessageType    int64     `json:"messageType" gorm:"comment:消息类型"`    //消息类型
	MessageContent string    `json:"messageContent" gorm:"comment:消息内容"` //消息内容
	Time           int64     `json:"time" gorm:"comment:发送时间"`           //发送时间
}
