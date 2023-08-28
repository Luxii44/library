package example

import (
	v1 "github.com/Luxii44/library/server/api/v1"
	"github.com/Luxii44/library/server/middleware"
	"github.com/gin-gonic/gin"
)

type MarkdownRouter struct{}

func (e *MarkdownRouter) InitMarkdownRouter(Router *gin.RouterGroup) {
	markdownRouter := Router.Group("markdown").Use(middleware.OperationRecord())
	//markdownRouterWithoutRecord := Router.Group("markdown")
	exaMarkdownApi := v1.ApiGroupApp.ExampleApiGroup.MarkdownApi
	{
		markdownRouter.POST("saveMarkdown", exaMarkdownApi.SaveMarkdown) // 保存markdown
	}
	{
		//markdownRouterWithoutRecord.POST("saveMarkdown", exaMarkdownApi.SaveMarkdown) // 创建聊天室
		//recruitRouterWithoutRecord.GET("account", exaChatApi.GetExaCustomer)          // 获取单一招聘账号信息
		//recruitRouterWithoutRecord.GET("customerList", exaChatApi.GetExaCustomerList) // 获取招聘账号列表
	}
}
