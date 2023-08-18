package example

import (
	v1 "github.com/Luxii44/library/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ChatRouter struct{}

func (e *ChatRouter) InitChatRouter(Router *gin.RouterGroup) {
	//recruitRouter := Router.Group("chat").Use(middleware.OperationRecord())
	chatRouterWithoutRecord := Router.Group("chat")
	exaChatApi := v1.ApiGroupApp.ExampleApiGroup.ChatApi
	{
		//recruitRouter.GET("ping", exaChatApi.CreateWS) // 创建聊天室
	}
	{
		chatRouterWithoutRecord.GET("ping", exaChatApi.CreateWS) // 创建聊天室
		//recruitRouterWithoutRecord.GET("account", exaChatApi.GetExaCustomer)          // 获取单一招聘账号信息
		//recruitRouterWithoutRecord.GET("customerList", exaChatApi.GetExaCustomerList) // 获取招聘账号列表
	}
}
